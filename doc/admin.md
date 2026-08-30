这是一个非常棒的选择！**`Soybean-Admin` 是目前国内代码质量最高、工程化最优雅的 Vue 3 中后台框架之一**。

它最大的亮点是采用了 **`@elegant-router/vue`（文件即路由）** 和 **`UnoCSS + Naive UI`**。在开发时，**你只要在 `src/views/` 下建一个文件夹，路由和菜单就会自动生成**，完全不需要配繁琐的数据库权限菜单！

以下为你量身定制的 **`frontend/admin`（Soybean-Admin）完整落地方案与 6 天开发计划**：

---

# 🛠️ GoShip-Admin (Soybean-Admin) 落地实施方案

---

## 一、 系统架构与极速瘦身（清理 Demo）

Soybean-Admin 默认带了很多演示页面（如富文本、图表、多级路由等），我们需要在 5 分钟内完成“瘦身”，只保留出海商业系统的核心功能。

### 1. 我们的目标目录结构（`src/views/`）
你只需要在 `src/views/` 目录下保留或新建以下 **4 个核心模块**：

```text
frontend/admin/src/views/
├── _builtin/               # 内置页面 (login, 404, 500 - 保留)
├── dashboard/              # 首页数据大盘 (订单额/注册数概览)
│   └── index.vue
├── products/               # 【核心】商品管理 (上下架/图片/价格)
│   └── index.vue
├── orders/                 # 【核心】订单与海外发货 (查看英文地址/填运单号)
│   └── index.vue
├── users/                  # 【核心】用户管理 (查看/手动充点数)
│   └── index.vue
└── system/                 # 【辅助】全局系统配置
    └── config/
        └── index.vue
```

> 💡 **Soybean-Admin 路由生成秘诀：**
> 删掉不需要的 demo 文件夹，新建好 `products`、`orders`、`users` 目录后，在终端执行一条命令：
> ```bash
> pnpm gen-route
> ```
> 系统会自动更新 `src/router/elegant/` 下的所有类型与路由，侧边栏菜单瞬间自动更新！

---

## 二、 对接 Go-zero 后端（网络与鉴权配置）

### 1. 配置环境变量指向 Go 后端
修改 `frontend/admin/.env.test` 或 `.env.prod`：

```env
# Go-zero 后端服务地址
VITE_SERVICE_BASE_URL=http://localhost:8888
```

### 2. API 接口封装规范（`src/service/api/`）
在 Soybean-Admin 中，使用内置的 `@sa/axios`（或 `request`）来对接 Go-zero：

#### ① 商品管理接口 (`src/service/api/product.ts`)
```typescript
import { request } from '../request';

// 获取商品列表
export function fetchProductList(params?: any) {
  return request<any>({
    url: '/api/v1/admin/products',
    method: 'GET',
    params
  });
}

// 保存/更新商品 (带 R2 图片与 Lemon Variant ID)
export function saveProduct(data: any) {
  const isEdit = Boolean(data.id);
  return request<any>({
    url: isEdit ? `/api/v1/admin/products/${data.id}` : '/api/v1/admin/products',
    method: isEdit ? 'PUT' : 'POST',
    data
  });
}

// 一键上下架切换
export function toggleProductStatus(id: number, status: 'active' | 'draft') {
  return request<any>({
    url: `/api/v1/admin/products/${id}/status`,
    method: 'PATCH',
    data: { status }
  });
}

// 直传 Cloudflare R2 图片接口
export function uploadToR2(file: File) {
  const formData = new FormData();
  formData.append('file', file);
  return request<{ url: string }>({
    url: '/api/v1/admin/upload',
    method: 'POST',
    data: formData,
    headers: { 'Content-Type': 'multipart/form-data' }
  });
}
```

#### ② 订单与海外发货接口 (`src/service/api/order.ts`)
```typescript
import { request } from '../request';

// 获取订单列表
export function fetchOrderList(params?: any) {
  return request<any>({
    url: '/api/v1/admin/orders',
    method: 'GET',
    params
  });
}

// 实物订单履约发货 (回填国际物流单号)
export function fulfillOrder(orderId: number, trackingNumber: string) {
  return request<any>({
    url: `/api/v1/admin/orders/${orderId}/fulfill`,
    method: 'POST',
    data: { trackingNumber }
  });
}
```

---

## 三、 核心页面 UI 设计（基于 Naive UI 组件）

### 1. 商品管理页 (`src/views/products/index.vue`)
*   **列表表格 (`NDataTable`)：**
    *   **主图：** `<NAvatar :src="row.cover_image" size="large" />`
    *   **标题与 Slug：** 显示商品名称及 SEO 链接标识。
    *   **售价：** `${(row.price_cents / 100).toFixed(2)}`（美分转美元格式化）。
    *   **类型标签：** `<NTag :type="row.product_type === 'digital' ? 'info' : 'warning'">`（区分虚拟/实物）。
    *   **上下架开关：** `<NSwitch v-model:value="row.status" checked-value="active" unchecked-value="draft" @update:value="handleStatusChange(row)" />`。
*   **新建/编辑弹窗 (`NModal` + `NForm`)：**
    *   上传封面：使用 `<NUpload :custom-request="handleCustomUpload">` 直传 Cloudflare R2，图片 URL 自动填入表单。
    *   绑定收银台：填入该商品在 Lemon Squeezy 上的 `lemon_variant_id`。

### 2. 订单管理与发货抽屉 (`src/views/orders/index.vue`)
*   **列表表格：** 查看订单号、支付状态（`<NTag type="success">PAID</NTag>`）、实付金额。
*   **发货详情抽屉 (`NDrawer`)（专门针对实物/茶叶/极客硬件）：**
    *   点击实物订单的【发货详情】，右侧滑出抽屉层。
    *   展示解析后的买家英文收货信息：
        ```text
        Name: John Doe
        Country: United States (US)
        Street: 742 Evergreen Terrace
        City: Springfield, OR
        Zip: 97477
        Phone: +1 555-0199
        ```
    *   底部提供 `<NInput v-model:value="trackingNo" placeholder="输入顺丰国际/云途物流单号" />`，点击 `<NButton type="primary">确认发货</NButton>`。

---

## 四、 6 天敏捷开发执行清单（每天只需 45 - 60 分钟）

```
[Day 1: 瘦身与接口打通] ➔ [Day 2: R2图片上传] ➔ [Day 3: 商品增删改查] ➔ [Day 4: 订单与发货] ➔ [Day 5: 用户与点数] ➔ [Day 6: 打包与联调]
```

- [ ] **Day 1：框架瘦身与后端代理打通**
    - [ ] 删掉 `src/views/` 下无关的 demo 文件夹，只保留 `dashboard`。
    - [ ] 新建 `src/views/products/index.vue` 和 `src/views/orders/index.vue`。
    - [ ] 运行 `pnpm gen-route`，确认侧边栏菜单已自动生成并正常展示。
    - [ ] 配置 `.env.test` 中的 `VITE_SERVICE_BASE_URL` 指向 `http://localhost:8888`。

- [ ] **Day 2：Cloudflare R2 图片直传对接**
    - [ ] 在 Go-zero 后端实现 `/api/v1/admin/upload` 接口（接收文件，调用 S3 SDK 推送至 Cloudflare R2 并返回 CDN URL）。
    - [ ] 在前端编写 `uploadToR2` 方法，并在商品页面测试图片上传与回显。

- [ ] **Day 3：商品管理 CRUD（核心）**
    - [ ] 开发 `src/views/products/index.vue` 表格展示。
    - [ ] 实现一键上下架 `<NSwitch>` 切换并即时同步数据库。
    - [ ] 实现新增/编辑商品弹窗：录入标题、价格（美分转换）、商品类型、Lemon Variant ID。

- [ ] **Day 4：订单管理与海外发货抽屉**
    - [ ] 开发 `src/views/orders/index.vue` 订单大盘表格。
    - [ ] 制作 `<NDrawer>` 发货抽屉，美化渲染买家的欧美详细英文地址。
    - [ ] 实现物流单号录入与标记发货逻辑。

- [ ] **Day 5：用户管理与点数充值**
    - [ ] 新建 `src/views/users/index.vue` 并运行 `pnpm gen-route`。
    - [ ] 表格展示注册用户列表、角色、当前 AI 剩余点数。
    - [ ] 提供一个“手动增减点数”的弹窗工具。

- [ ] **Day 6：构建校验与联调测试**
    - [ ] 在终端运行 `pnpm build`，确保 TypeScript 类型检查 100% 通过无报错。
    - [ ] 在浏览器全流程走通一遍：后台上架商品 ➔ 上传图片 ➔ 检查数据库状态。
    - [ ] 提交代码：`git commit -m "feat(admin): complete soybean-admin integration"`。

---

### 🎯 今天的唯一动作：

1. 打开 `frontend/admin` 目录，删掉不需要的 demo 页面；
2. 新建 `src/views/products/index.vue`，在终端运行 **`pnpm gen-route`**；
3. 运行 **`pnpm dev`** 打开后台，看侧边栏是不是已经有了漂亮的“商品管理”菜单！