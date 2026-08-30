package database

import (
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

//Mysql: gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local
//Postgres: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai

type Config struct {
	Driver          string   `json:",default=mysql"`
	SlowThreshold   int      `json:",default=2"`   //慢查询阈值，单位是秒。当查询执行时间超过这个阈值时，会记录慢查询日志。默认5秒
	MaxIdleConn     int      `json:",default=10"`  //连接池中最大空闲连接数。默认10。
	MaxOpenConn     int      `json:",default=100"` //连接池中最大打开连接数。默认100。
	ConnMaxIdleTime int      `json:",default=30"`  //连接的最大空闲时间，单位是分钟。如果连接空闲时间超过这个值，会被关闭。默认60分钟。
	ConnMaxLifetime int      `json:",default=60"`  //连接的最大存活时间，单位是小时。超过这个时间的连接会被关闭并重新建立。默认1小时。
	LogLevel        string   `json:",default=info,optional"`
	Reade           []string `json:",optional"`
	Write           string
	Log             string `json:"Log,optional"`
}

// NewDatabase 创建数据库连接
func NewDatabase(config *Config) (*gorm.DB, error) {
	// 初始化GORM配置
	gormConfig := &gorm.Config{
		PrepareStmt: true,
		//启用预编译语句（Prepared Statements）缓存。 相同的 SQL 语句会被预编译并缓存，提高重复查询的性能 减少 SQL 解析和编译的开销 提供一定的 SQL 注入防护
		SkipDefaultTransaction: true,
		//跳过默认的事务包装。GORM 默认会在每个写操作（Create、Update、Delete）外包装一个事务  设置为 true 会跳过这个默认的事务包装，提高性能
		DisableNestedTransaction: true,
		//当在一个事务中开启另一个事务时，GORM 默认使用保存点（SavePoint）实现嵌套事务 设置为 true 会禁用这种嵌套事务行为  推荐开启，避免复杂的嵌套事务逻辑
		DisableForeignKeyConstraintWhenMigrating: true,
		//GORM 在自动迁移（AutoMigrate）时会尝试创建外键约束  设置为 true 会禁用此外键约束创建
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		//默认情况下，GORM 使用结构体名的复数形式作为表名（如 User → users）
	}
	// 设置日志级别
	switch config.LogLevel {
	case "silent":
		gormConfig.Logger = NewGormZeroLogger(config.Log, logger.Silent)
	case "error":
		gormConfig.Logger = NewGormZeroLogger(config.Log, logger.Error)
	case "warn":
		gormConfig.Logger = NewGormZeroLogger(config.Log, logger.Warn)
	case "info":
		gormConfig.Logger = NewGormZeroLogger(config.Log, logger.Info)
	default:
		gormConfig.Logger = NewGormZeroLogger(config.Log, logger.Info)
	}
	// 连接主数据库
	var dialector gorm.Dialector
	switch config.Driver {
	case "mysql":
		dialector = mysql.Open(config.Write)
	case "postgres":
		dialector = postgres.Open(config.Write)
	default:
		return nil, fmt.Errorf("不支持的数据库驱动: %s", config.Driver)
	}

	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("连接主数据库失败: %v", err)
	}
	// 如果有从库配置，设置读写分离
	if len(config.Reade) > 0 {
		replicaDialectors := make([]gorm.Dialector, 0, len(config.Reade))
		for _, replica := range config.Reade {
			switch config.Driver {
			case "mysql":
				replicaDialectors = append(replicaDialectors, mysql.Open(replica))
			case "postgres":
				replicaDialectors = append(replicaDialectors, postgres.Open(replica))
			}
		}

		// 配置DBResolver
		resolverConfig := dbresolver.Config{
			Sources:  []gorm.Dialector{dialector}, // 主库
			Replicas: replicaDialectors,           // 从库列表
			Policy:   dbresolver.RandomPolicy{},   // 随机选择从库
		}

		err = db.Use(dbresolver.Register(resolverConfig).
			SetMaxIdleConns(config.MaxIdleConn).
			SetMaxOpenConns(config.MaxOpenConn).
			SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTime) * time.Minute).
			SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute))
		if err != nil {
			return nil, fmt.Errorf("配置读写分离失败: %v", err)
		}
	} else {
		// 没有从库，直接配置连接池
		sqlDB, errx := db.DB()
		if errx != nil {
			return nil, errx
		}
		sqlDB.SetMaxIdleConns(config.MaxIdleConn)
		sqlDB.SetMaxOpenConns(config.MaxOpenConn)
		sqlDB.SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTime) * time.Minute)
		sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)
	}
	err = db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError) // 将record not fund 去掉
	if err != nil {
		return nil, err
	}
	logx.Infof("数据库连接成功，驱动: %s", config.Driver)
	return db, nil
}

// Close 关闭数据库连接
func Close(db *gorm.DB) error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
