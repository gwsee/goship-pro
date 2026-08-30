package database

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/color"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

// GormZeroLogger 实现分流输出到 logs/sql/ 目录下的 access.log, error.log, slow.log
type GormZeroLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration // 慢 SQL 判定阈值 (如 200ms)
	accessWriter  logx.Writer   // 对应 logs/sql/access.log
	errorWriter   logx.Writer   // 对应 logs/sql/error.log
	slowWriter    logx.Writer   // 对应 logs/sql/slow.log
}

// NewGormZeroLogger 创建支持 access/error/slow 三流分写的 SQL Logger
func NewGormZeroLogger(sqlLogDir string, logLevel logger.LogLevel) logger.Interface {
	var accessWriter, errorWriter, slowWriter logx.Writer
	if sqlLogDir != "" {
		// 1. 分别指定 3 个文件的路径
		accessFile := path.Join(sqlLogDir, "sql/access.log")
		errorFile := path.Join(sqlLogDir, "sql/error.log")
		slowFile := path.Join(sqlLogDir, "sql/slow.log")

		// 2. 分别创建按天切割、保留 7 天的轮转规则
		accessRule := logx.NewSizeLimitRotateRule(accessFile, ".", 7, 500, 7, false)
		errorRule := logx.NewSizeLimitRotateRule(errorFile, ".", 7, 500, 7, false)
		slowRule := logx.NewSizeLimitRotateRule(slowFile, ".", 7, 500, 7, false)

		// 3. 调用 rotatelogger.go 中的 NewLogger 初始化 3 个独立的 RotateLogger
		accLog, err1 := logx.NewLogger(accessFile, accessRule, false)
		errLog, err2 := logx.NewLogger(errorFile, errorRule, false)
		slwLog, err3 := logx.NewLogger(slowFile, slowRule, false)

		// 4. 包装成 logx.Writer
		if err1 == nil && err2 == nil && err3 == nil {
			accessWriter = logx.NewWriter(accLog)
			errorWriter = logx.NewWriter(errLog)
			slowWriter = logx.NewWriter(slwLog)
		} else {
			logx.Errorf("创建 SQL 专用分流 Writer 失败, 将降级使用全局 logx")
		}
	}

	return &GormZeroLogger{
		LogLevel:      logLevel,
		SlowThreshold: 200 * time.Millisecond, // 默认超过 200ms 认定为慢 SQL
		accessWriter:  accessWriter,
		errorWriter:   errorWriter,
		slowWriter:    slowWriter,
	}
}

func (l GormZeroLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l GormZeroLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		caller := l.getCaller()
		str := fmt.Sprintf("[%s] %s", caller, fmt.Sprintf(msg, data...))
		if l.accessWriter != nil {
			l.accessWriter.Info(str)
		} else {
			logx.WithContext(ctx).Info(str)
		}
	}
}

func (l GormZeroLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		caller := l.getCaller()
		str := fmt.Sprintf("[%s] %s", caller, fmt.Sprintf(msg, data...))
		if l.errorWriter != nil {
			l.errorWriter.Error(str)
		} else {
			logx.WithContext(ctx).Error(str)
		}
	}
}

func (l GormZeroLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		caller := l.getCaller()
		str := fmt.Sprintf("[%s] %s", caller, fmt.Sprintf(msg, data...))
		if l.errorWriter != nil {
			l.errorWriter.Error(logx.WithColor(str, color.FgRed))
		} else {
			logx.WithContext(ctx).Error(logx.WithColor(str, color.FgRed))
		}
	}
}

// response.Success(w, struct{}{})核心 Trace：根据执行结果与耗时自动分流写入不同的日志文件
func (l GormZeroLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	sql, rows := fc()
	elapsed := time.Since(begin)
	caller := l.getCaller() // 获取实际调用业务行号

	// 1. SQL 执行报错 ➔ error.log
	if err != nil {
		str := fmt.Sprintf("[%s] [%v] SQL执行错误: %s, rows: %d, err: %v", caller, elapsed, sql, rows, err.Error())
		if l.errorWriter != nil {
			l.errorWriter.Error(logx.WithColor(str, color.FgRed))
		} else {
			logx.WithContext(ctx).WithDuration(elapsed).Error(logx.WithColor(str, color.FgRed))
		}
		return
	}

	// 2. 慢 SQL ➔ slow.log
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		str := fmt.Sprintf("[%s] [%v] 慢SQL警告(超过%v): %s, rows: %d", caller, elapsed, l.SlowThreshold, sql, rows)
		if l.slowWriter != nil {
			l.slowWriter.Slow(str)
		} else {
			logx.WithContext(ctx).WithDuration(elapsed).Slow(str)
		}
		return
	}

	// 3. 正常 SQL ➔ access.log
	str := fmt.Sprintf("[%s] [%v] SQL: %s, rows: %d", caller, elapsed, sql, rows)
	if l.accessWriter != nil {
		l.accessWriter.Info(str)
	} else {
		logx.WithContext(ctx).WithDuration(elapsed).Info(str)
	}
}
func (l GormZeroLogger) getCaller() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// 1. 过滤掉框架、Gen生成代码、数据库底层代码
		if strings.Contains(file, "gorm.io") ||
			strings.Contains(file, "github.com/zeromicro") ||
			strings.Contains(file, "internal/cfgx/database") || // 当前 logger
			strings.Contains(file, "database") ||
			strings.Contains(file, "dao/query") { // 过滤 GORM Gen 生成的 query 代码
			continue
		}

		// 2. 裁剪路径：只保留从项目根目录开始的路径 (例如从 crm-back 或 internal 开始)
		cleanFile := trimProjectRoot(file)

		// 3.response.Success(w, struct{}{})直接使用遍历出来的 file 和 line，不要再调 utils.FileWithLineNum()
		return fmt.Sprintf("%s:%d", cleanFile, line)
	}

	return "unknown:0"
}

// trimProjectRoot 智能裁剪掉本地 E:/xxx/ 或 /home/xxx/ 绝对前缀
func trimProjectRoot(fullPath string) string {
	// 将 Windows 路径反斜杠转为统一斜杠
	formattedPath := strings.ReplaceAll(fullPath, "\\", "/")

	// 优先匹配你项目的根目录名称（比如 crm-back）
	const projectRoot = "crm-back/"
	if idx := strings.Index(formattedPath, projectRoot); idx != -1 {
		return formattedPath[idx+len(projectRoot):] // 输出: internal/logic/xxx/xxx.go
	}

	// 兜底匹配 internal/
	if idx := strings.Index(formattedPath, "internal/"); idx != -1 {
		return formattedPath[idx:]
	}

	return path.Base(formattedPath) // 如果都未命中，至少只展示文件名
}
