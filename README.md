<div align="center">

# 🚢 GoShip Pro - Full-Stack SaaS & Commerce Starter Kit

<p align="center">
  <strong>The ultimate Go (go-zero) + Nuxt 3 boilerplate. Ship your profitable global SaaS or E-commerce product in days.</strong>
</p>

<p align="center">
  <a href="#en_doc">English</a> •
  <a href="#zh_doc">中文文档</a> •
  <a href="https://your-demo-url.com" target="_blank">Live Demo</a> •
  <a href="https://your-sales-url.com" target="_blank">Get Pro License ($149)</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go" alt="Go Version" />
  <img src="https://img.shields.io/badge/Nuxt-3.x-00DC82?style=flat&logo=nuxt.js" alt="Nuxt Version" />
  <img src="https://img.shields.io/badge/Vue-3.x-4FC08D?style=flat&logo=vue.js" alt="Vue Version" />
  <img src="https://img.shields.io/badge/Tailwind-3.x-38B2AC?style=flat&logo=tailwind-css" alt="Tailwind" />
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License" />
</p>

</div>

---

<a id="en_doc"></a>
### 💡 Why GoShip Pro?
Most overseas SaaS boilerplates on the market are built on Next.js/React. If you prefer **Go's raw performance** and **Vue 3's simplicity**, you often have to build authentication, payments, and admin dashboards from scratch.

**GoShip Pro** is built to bridge this gap. It provides an enterprise-ready, ultra-lightweight monolithic architecture that runs smoothly on a **$5/mo VPS with less than 150MB of RAM**.

### ✨ Key Features
- ⚡ **Lightweight Go-zero Monolith**: Clean REST API architecture, zero RPC/microservice bloat. Auto-generated CRUD powered by `goctl`.
- 🎨 **Nuxt 3 Storefront (SSR + SEO)**: Dynamic product showcase, OpenGraph social cards, and 100/100 Google Lighthouse ready.
- 🛠️ **Vue 3 Admin Dashboard (SPA)**: Turnkey management for product publish/unpublish, Cloudflare R2 image uploads, customer orders, and shipping tracking.
- 💳 **Global Payments (Lemon Squeezy)**: Built-in webhook listeners for automated credit top-up and physical shipping address collection.
- 🗄️ **Zero-Cost Object Storage**: Direct integration with **Cloudflare R2** (S3 compatible, 0 egress fees).
- 🤖 **Turnkey AI Demo**: Server-Sent Events (SSE) streaming output with credit deduction logic.
- 🐳 **Production-Ready DevOps**: Complete Docker Compose setups and Nginx reverse-proxy configs.

### 🏗️ Project Structure
```text
goship-pro/
├── backend/          # Go-zero Monolith REST API Backend
│   ├── etc/          # Service configurations (config.yaml.example)
│   ├── internal/     # Core business logic (handler, logic, svc, model)
│   └── Dockerfile
├── doc/              # Architecture designs, MySQL schemas, and API documentation
├── frontend/
│   ├── admin/        # Vue 3 + shadcn Admin Dashboard (SPA)
│   └── web/          # Nuxt 3 Storefront & Landing Page (SSR + SEO)
└── ops/              # DevOps: Docker Compose, Nginx SSL configs, deploy scripts
```

### 🚀 Quick Start (Local Setup)

1. **Clone the repository:**
   ```bash
   git clone https://github.com/gwsee/goship-pro.git
   cd goship-pro
   ```

2. **Setup environment variables:**
   ```bash
   cp .env.example .env
   cp backend/etc/config.yaml.example backend/etc/config.yaml
   # Fill in your MySQL, Lemon Squeezy, and Cloudflare R2 credentials
   ```

3. **Initialize Database:**
   Import `doc/schema.sql` into your local MySQL 8.0 database.

4. **Start Services with Docker Compose:**
   ```bash
   docker compose -f ops/docker-compose.yml up -d
   ```
   - Storefront (Web): `http://localhost:3000`
   - Admin Panel: `http://localhost:3001`
   - Backend API: `http://localhost:8888`

---

<a id="zh_doc"></a>
### 💡 为什么选择 GoShip Pro？
市面上 90% 的出海脚手架都基于 Next.js/React，不仅臃肿且服务器成本高昂。对于熟悉 **Go 语言的高性能** 与 **Vue 3 优雅体验** 的开发者来说，出海往往缺少一套趁手的“商业武器”。

**GoShip Pro 专为解决此痛点而生**。它为你提供了一套开箱即用的出海商业底座（同时支持 **SaaS/AI 虚拟订阅** 与 **实体周边独立站电商**），单机运行内存占用不到 150MB，月租 30 元的最低配 VPS 即可平稳运行。

### ✨ 核心特性
- ⚡ **Go-zero 极简单体后端**：摒弃分布式微服务的累赘，纯单体 REST API 架构，利用 `goctl` 秒级生成业务代码。
- 🎨 **Nuxt 3 动态前台与极致 SEO**：服务端渲染（SSR），自动抓取后台上架商品，秒级首屏响应。
- 🛠️ **全功能 Vue 3 管理后台**：开箱即用的商品上下架、Cloudflare R2 图片直传、查看买家英文地址与运单号回填。
- 💳 **Lemon Squeezy 全球收单闭环**：国内个人身份合规收美金，自动化 Webhook 回调、虚拟点数充值与实物收件地址解析。
- 🗄️ **Cloudflare R2 零成本存储**：集成 S3 兼容协议，商品图片直传，**下行流量费完全为 0**。
- 🤖 **内置 AI 流式对话演示**：封装打字机效果（SSE 流式传输），开箱即用当 AI 工具变现。
- 🐳 **一键 Docker 交付体系**：预置生产级 Dockerfile 与 `docker-compose`，运维门槛降为 0。

### 🏗️ 目录结构说明
```text
goship-pro/
├── backend/          # Go-zero 单体后端 (提供核心 REST API)
│   ├── etc/          # 配置文件目录 (包含 config.yaml.example 模板)
│   ├── internal/     # 业务核心层 (handler, logic, svc, model)
│   └── Dockerfile
├── doc/              # 项目架构文档、MySQL 建表 DDL、接口说明
├── frontend/
│   ├── admin/        # 管理后台 (Vue 3 + shadcn, SPA 单页架构)
│   └── web/          # 前台商店与 Landing Page (Nuxt 3, SSR 架构 + SEO)
└── ops/              # 运维目录: Docker Compose、Nginx 配置与部署脚本
```

### 🚀 本地极速启动

1. **克隆项目到本地：**
   ```bash
   git clone https://github.com/gwsee/goship-pro.git
   cd goship-pro
   ```

2. **配置环境变量与密钥：**
   ```bash
   cp .env.example .env
   cp backend/etc/config.yaml.example backend/etc/config.yaml
   # 填入你的 MySQL、Lemon Squeezy 测试密钥和 Cloudflare R2 密钥
   ```

3. **初始化数据库：**
   将 `doc/schema.sql` 导入你的本地 MySQL 8.0 数据库中。

4. **一键启动所有服务：**
   ```bash
   docker compose -f ops/docker-compose.yml up -d
   ```
   - 前台商店 (Nuxt 3)：`http://localhost:3000`
   - 管理后台 (Vue 3)：`http://localhost:3001`
   - 后端接口 (Go-zero)：`http://localhost:8888`

---

## 📄 License & Commercial
- 本项目的开源精简版基于 [MIT License](LICENSE) 授权。
- 商业授权版（包含完整 AI 演示、VIP 专属更新群与终身升级）：请访问 [goshipbase.com](https://your-sales-url.com)。