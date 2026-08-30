# 📋 GoShip-Pro 开发执行清单 (TODO.md)

> **开发心法：** 
> 1. 每天打开电脑，只看当天那一行的 `[ ]`，完成打勾即算胜利！
> 2. 严禁提前过度重构，**“跑通逻辑 > 追求代码优雅”**。
> 3. 下班睡觉前故意留一行编译报错，方便第二天秒级启动。

---

## 阶段零：文档归档与基建配置（第 1 - 2 天）

- [ ] **Day 1：数据库 DDL 与账号凭证准备**
  - [ ] 在 `doc/` 目录下新建 `schema.sql`，把 4 张核心表 SQL（users, products, orders, system_configs）保存进去。
  - [ ] 在本地 MySQL 执行 `doc/schema.sql` 创建数据库。
  - [ ] 注册并获取 **Lemon Squeezy 沙箱 Key** 与 **Cloudflare R2 Bucket 密钥**。

- [ ] **Day 2：Go-zero 工程初始化与代码生成**
  - [ ] 进入 `backend/` 目录，执行 `go mod init goship/backend`。
  - [ ] 在 `backend/` 下编写基础 `app.api` 文件。
  - [ ] 运行 `goctl model mysql ddl -src="../doc/schema.sql" -dir="./internal/model" -c` 一键生成全部数据库 CRUD 代码。
  - [ ] 配置 `backend/etc/config.yaml.example` 和本地使用的 `config.yaml`。

---

## 阶段一：后端核心与支付闭环（第 3 - 7 天，目录：`backend/`）
*目标：不写前端页面，用 Postman/Apifox 彻底调通所有核心 API。*

- [ ] **Day 3：用户鉴权模块 (Auth)**
  - [ ] 实现 `POST /api/v1/auth/register`（密码 Bcrypt 加密落库）。
  - [ ] 实现 `POST /api/v1/auth/login`（验证密码，生成 JWT Token）。
  - [ ] 配置 Go-zero 自带的 JWT 鉴权中间件。

- [ ] **Day 4：公开商品与发起支付 API**
  - [ ] 实现 `GET /api/v1/store/products`（只查询 `status = 'active'` 的商品列表）。
  - [ ] 实现 `POST /api/v1/payment/checkout`（接收 `productId`，调用 Lemon Squeezy SDK 生成带有 `userId` 的收银台 URL）。

- [ ] **Day 5-6：Lemon Squeezy Webhook 闭环（重中之重）**
  - [ ] 实现 `POST /api/v1/webhook/lemonsqueezy`。
  - [ ] 校验 Webhook 签名（保证请求来自 Lemon Squeezy 官方）。
  - [ ] 解析支付成功事件：
    - 若为虚拟商品：自动将数据库中对应用户的 `credits += credits_reward`。
    - 若为实物商品：提取买家英文收货地址 JSON，存入 `orders.shipping_info`。
    - 将订单状态更新为 `paid`。

- [ ] **Day 7：后端自测与里程碑**
  - [ ] 用 Postman 模拟一次完整闭环：注册 ➔ 登录 ➔ 创建订单 ➔ 触发 Webhook ➔ 验证数据库。
  - [ ] 提交代码：`git commit -m "feat(backend): complete auth and payment webhook flow"`。

---

## 阶段二：Admin 管理后台开发（第 8 - 14 天，目录：`frontend/admin/`）
*目标：做出一个能给商品上下架、能看收货地址的清爽后台。*

- [ ] **Day 8-9：Vue 3 Admin 框架与鉴权**
  - [ ] 在 `frontend/admin/` 初始化 Vue 3 SPA + Tailwind CSS + `shadcn-vue`。
  - [ ] 封装 Axios 请求拦截器（自动在 Header 携带 JWT Token，401 自动跳登录）。
  - [ ] 搭建管理后台基础布局（侧边栏、顶部导航、退出登录）。

- [ ] **Day 10-11：Cloudflare R2 图片直传与商品管理**
  - [ ] 在 `backend/` 编写 `/api/v1/admin/upload` 接口（接收图片流，调用 S3 SDK 直传 R2 并返回 CDN URL）。
  - [ ] 前端制作**商品列表页**：展示商品图片、价格、类型（实物/虚拟），带有一键【上下架 Switch 开关】。
  - [ ] 前端制作**新增/编辑商品弹窗**：支持填写标题、价格、绑定 Lemon Variant ID、上传封面图。

- [ ] **Day 12-13：订单大盘与实物发货管理**
  - [ ] 制作**订单列表页**：查看订单号、支付状态（Paid/Pending）、实付金额。
  - [ ] 制作**实物订单详情抽屉 (Drawer)**：
    - 点击订单展开买家英文收件地址（Name, Country, City, Address, Zip, Phone）。
    - 提供“输入国际物流单号”并点击【标记已发货】的按钮。

- [ ] **Day 14：用户管理与阶段小结**
  - [ ] 制作简易用户列表页，支持后台手动给指定用户“充值点数”。
  - [ ] 提交代码：`git commit -m "feat(admin): complete product, order, and user dashboard"`。

---

## 阶段三：前台商店与 SEO 落地页（第 15 - 21 天，目录：`frontend/web/`）
*目标：打造高转化率的前台页面，让老外愿意掏卡付款。*

- [ ] **Day 15-17：Nuxt 3 落地页 (Landing Page)**
  - [ ] 在 `frontend/web/` 初始化 Nuxt 3 工程（开启 SSR）。
  - [ ] 安装 Tailwind CSS 与 Lucide 图标库。
  - [ ] 组装落地页组件：Hero 屏、产品特性卡片网格、FAQ 手风琴折叠栏、页脚。

- [ ] **Day 18-19：动态商品橱窗与购买跳转**
  - [ ] 使用 Nuxt 3 的 `useFetch('/api/v1/store/products')` 服务端异步拉取已上架商品。
  - [ ] 循环渲染商品卡片（封面图、标题、美元价格、特性标签）。
  - [ ] 绑定【Buy Now】按钮：未登录提示一键登录，已登录直接请求后端收银台链接并跳转。
  - [ ] 配置 `useSeoMeta`，自动生成 Twitter Card 和 OpenGraph 社交预览卡片。

- [ ] **Day 20-21：杀手级 AI 流式演示（模板溢价卖点）**
  - [ ] 在 `backend/` 编写 SSE (Server-Sent Events) 打字机流式 API，调用 OpenAI/Claude 并扣除用户 1 个点数。
  - [ ] 在 `frontend/web/` 制作一个“AI 文案/翻译生成器 Demo”，前台打字机逐字渲染。
  - [ ] 提交代码：`git commit -m "feat(web): complete nuxt3 landing page and dynamic store"`。

---

## 阶段四：Docker 容器化、文档与发布（第 22 - 28 天，目录：`ops/` & `doc/`）
*目标：打包成极简交付物，正式上线变现。*

- [ ] **Day 22-23：DevOps 容器化编写**
  - [ ] 编写 `backend/Dockerfile`（多阶段构建，输出 15MB 极小二进制镜像）。
  - [ ] 编写 `frontend/web/Dockerfile` 与 `frontend/admin/Dockerfile`。
  - [ ] 编写 `ops/docker-compose.yml`（一键拉起：MySQL + 后端 + 前台 + 后台）。
  - [ ] 本地测试 `docker compose -f ops/docker-compose.yml up`，验证所有端口连通。

- [ ] **Day 24：编写生产级交付文档**
  - [ ] 在 `doc/` 下编写 `deployment.md`（VPS 一键部署指南、Nginx 反代与 SSL 证书申请）。
  - [ ] 完善根目录的 `README.md` 与 `.env.example`。

- [ ] **Day 25-28：官方发售与全网宣发**
  - [ ] 购买一个 60 元域名（如 `goshipbase.com`），用这套系统部署你自己的官方销售站。
  - [ ] 将代码中的测试 Key 切换为 Lemon Squeezy 正式收款 Key。
  - [ ] 在 GitHub 开源 Lite 精简版，README 挂上商业版 Pro 购买链接（$149）。
  - [ ] 在 Twitter/X（带标签 `#buildinpublic` `#golang` `#vuejs`）、Reddit（`r/golang`）、V2EX（程序员节点）发帖宣发。

---
