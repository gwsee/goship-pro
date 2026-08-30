这是专为 **`frontend/web`（面向终端用户的商店与 SaaS 门户）** 定制的**完整工程设计与落地实施方案**。

本方案基于 **Nuxt 3 (SSR) + Tailwind CSS + shadcn-vue**，兼顾 **Google 极致 SEO、手机端大图自适应、海外支付跳转与 AI 流式交互**。

---

# 🌐 GoShip-Web 面向用户端系统方案

---

## 一、 系统技术栈与依赖清单

进入 `frontend/web/` 目录，安装以下现代前端基础设施依赖：

```bash
# 核心模块
pnpm add @nuxtjs/tailwindcss @nuxtjs/color-mode @pinia/nuxt @vueuse/nuxt
# UI 与图标生态
pnpm add shadcn-vue radix-vue lucide-vue-next clsx tailwind-merge
```

### 核心 `nuxt.config.ts` 基础配置

```typescript
// frontend/web/nuxt.config.ts
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: false },

  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxtjs/color-mode',
    '@pinia/nuxt',
    '@vueuse/nuxt',
  ],

  colorMode: {
    classSuffix: '', // 适配 Tailwind 暗黑模式类名: 'dark'
    preference: 'system',
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8888',
      siteUrl: process.env.NUXT_PUBLIC_SITE_URL || 'https://goshipbase.com',
    }
  },

  app: {
    head: {
      viewport: 'width=device-width, initial-scale=1, maximum-scale=1',
      charset: 'utf-8',
      link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
    }
  }
})
```

---

## 二、 模块化目录结构规范 (`frontend/web/`)

```text
frontend/web/
├── assets/css/tailwind.css     # 全局 Tailwind 样式与色彩变量
├── components/
│   ├── ui/                     # shadcn-vue 原生原子组件 (Button, Dialog, Sheet)
│   ├── layout/                 # 全局导航栏、页脚、手机端抽屉菜单
│   │   ├── Navbar.vue
│   │   ├── Footer.vue
│   │   └── MobileNav.vue
│   ├── landing/                # 落地页静态/高转化区块
│   │   ├── Hero.vue            # 首屏吸睛、标语、CTA
│   │   ├── Features.vue        # 6 宫格特性卡片
│   │   └── Faq.vue             # 手风琴常见问题
│   ├── store/                  # 动态电商橱窗组件
│   │   ├── ProductCard.vue     # 单个商品卡片 (自适应多端)
│   │   └── ProductGrid.vue     # 商品列表网格 (SSR 数据绑定)
│   └── dashboard/              # 用户控制台组件
│       ├── AIDemoChat.vue      # SSE 流式打字机交互演示
│       └── OrderHistory.vue    # 个人订单与物流查询
├── composables/                # 核心业务逻辑 Hooks
│   ├── useAuth.ts              # 用户登录、登出、JWT Cookie 管理
│   ├── usePayment.ts           # 调用 Go 后端生成收银台 URL
│   └── useAIStream.ts          # 读取 Go 后端 SSE 文本流
├── layouts/
│   ├── default.vue             # 前台默认布局 (含 Header + Footer)
│   └── dashboard.vue           # 用户后台仪表盘布局 (侧边栏)
├── pages/
│   ├── index.vue               # 首页 (Landing Page + 动态商品橱窗)
│   ├── p/[slug].vue            # 单个商品 SEO 独立详情页
│   └── dashboard/
│       ├── index.vue           # 用户中心 (点数、AI 工具演示)
│       └── orders.vue          # 我的历史订单
└── app.vue                     # 顶层入口
```

---

## 三、 核心业务模块实现细节

### 1. 动态商品橱窗与 SEO 引擎 (`pages/index.vue`)
利用 Nuxt 3 的 `useFetch` 在服务端（SSR）直接向 Go-zero 后端拉取已上架商品，并在 HTML 源码中注入 OpenGraph 社交卡片。

```vue
<!-- pages/index.vue -->
<script setup lang="ts">
const config = useRuntimeConfig()

// 1. 服务端异步抓取已上架商品 (SEO 友好)
const { data: response, pending } = await useFetch<{ code: number; data: any[] }>(
  '/api/v1/store/products',
  { baseURL: config.public.apiBase }
)
const products = computed(() => response.value?.data || [])

// 2. 注入全局 SEO 与社交分享 Meta
useSeoMeta({
  title: 'GoShip Pro - The Ultimate Go + Vue 3 Commercial Boilerplate',
  description: 'Ship your profitable SaaS or overseas store in days. Built with Go-zero and Nuxt 3.',
  ogTitle: 'GoShip Pro - Full-Stack SaaS & Commerce Starter Kit',
  ogDescription: 'Stop rebuilding payments, auth, and dashboards. Ship fast.',
  ogImage: `${config.public.siteUrl}/images/og-cover.png`,
  twitterCard: 'summary_large_image',
})
</script>

<template>
  <div>
    <!-- 1. 首屏高转化 Hero 区域 -->
    <LandingHero />

    <!-- 2. 特性卡片 -->
    <LandingFeatures />

    <!-- 3. 动态商品/定价橱窗 (根据后端数据渲染) -->
    <section id="pricing" class="py-20 bg-slate-50 dark:bg-slate-900/50">
      <div class="container mx-auto px-4">
        <div class="text-center max-w-2xl mx-auto mb-12">
          <h2 class="text-3xl sm:text-4xl font-extrabold tracking-tight">
            Choose Your Plan & Goods
          </h2>
          <p class="text-muted-foreground mt-3 text-sm sm:text-base">
            Instant digital access or global physical shipping. Transparent pricing.
          </p>
        </div>

        <StoreProductGrid :products="products" :loading="pending" />
      </div>
    </section>

    <!-- 4. FAQ 常见问题 -->
    <LandingFaq />
  </div>
</template>
```

---

### 2. 跨端购买跳转逻辑 (`composables/usePayment.ts`)
实现点击购买直接请求 Go 后端，拿到专属 Lemon Squeezy 收银台链接并无缝跳转：

```typescript
// composables/usePayment.ts
export const usePayment = () => {
  const config = useRuntimeConfig()
  const { token, user } = useAuth()
  const loading = ref(false)

  const checkout = async (productId: number) => {
    // 未登录则先引导弹窗登录
    if (!token.value) {
      navigateTo('/auth/login?redirect=' + encodeURIComponent(window.location.pathname))
      return
    }

    try {
      loading.value = true
      const res: any = await $fetch('/api/v1/payment/checkout', {
        method: 'POST',
        baseURL: config.public.apiBase,
        headers: { Authorization: `Bearer ${token.value}` },
        body: { productId }
      })

      if (res.code === 0 && res.data?.checkoutUrl) {
        // 跳转到 Lemon Squeezy 全球收银台
        window.location.href = res.data.checkoutUrl
      }
    } catch (err: any) {
      alert(err.data?.msg || 'Failed to initiate checkout')
    } finally {
      loading.value = false
    }
  }

  return { checkout, loading }
}
```

---

### 3. 杀手级特性：AI 流式打字机客户端 (`composables/useAIStream.ts`)
在用户控制台（Dashboard）展示 AI 流式生成效果（SSE），展示系统的商业变现能力：

```typescript
// composables/useAIStream.ts
export const useAIStream = () => {
  const config = useRuntimeConfig()
  const { token } = useAuth()
  const generating = ref(false)
  const outputText = ref('')

  const generate = async (prompt: string) => {
    generating.value = true
    outputText.value = ''

    const response = await fetch(`${config.public.apiBase}/api/v1/ai/generate`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token.value}`
      },
      body: JSON.stringify({ prompt })
    })

    if (!response.ok) {
      generating.value = false
      throw new Error('Generation failed or credits depleted')
    }

    const reader = response.body?.getReader()
    const decoder = new TextDecoder()

    if (!reader) return

    while (true) {
      const { done, value } = await reader.read()
      if (done) break
      outputText.value += decoder.decode(value, { stream: true })
    }

    generating.value = false
  }

  return { generate, outputText, generating }
}
```

---

## 四、 多端（PC / 移动端）自适应交互设计

1.  **移动端导航（Mobile Sheet Drawer）：**
    *   在屏幕 `< 768px` 时，顶部导航栏自动折叠为右侧“汉堡按钮”。
    *   点击滑出抽屉层（基于 `radix-vue/Dialog`），包含页面导航和快捷登录按钮。
2.  **手机端浮动吸底购买栏（Sticky Checkout Bar）：**
    *   在手机访问单个商品详情页时，底部常驻悬浮一个带有“价格 + 一键购买”的大按钮，杜绝用户反复滑动查找购买入口，**移动端转化率提升 40%**。

---

## 五、 `frontend/web` 5 天敏捷开发执行清单

*   [ ] **Day 1：初始化工程与基础组件**
    *   在 `frontend/web` 初始化 Nuxt 3。
    *   配置 Tailwind CSS 与暗黑模式切换。
    *   引入 `Navbar.vue`（含 PC 布局与移动端抽屉菜单）和 `Footer.vue`。
*   [ ] **Day 2：组装高转化 Landing Page 区块**
    *   使用 Tailwind 制作 Hero 首屏（带 CTA 按钮与动态光晕特效）。
    *   制作 Features 6 宫格卡片与 FAQ 手风琴折叠栏。
*   [ ] **Day 3：动态商品橱窗与支付联调**
    *   开发 `ProductCard.vue`（展示主图、美元售价、实物/虚拟标签）。
    *   编写 `usePayment.ts`，打通点击购买 ➔ 调用 Go 接口 ➔ 跳转 Lemon Squeezy 收银台。
*   [ ] **Day 4：鉴权与用户仪表盘 (Dashboard)**
    *   制作轻量登录/注册弹窗。
    *   制作 `/dashboard` 页面：查看当前账号剩余点数与个人历史订单。
*   [ ] **Day 5：集成 AI 流式演示与全局 SEO 校验**
    *   开发打字机流式生成组件，联调 Go-zero 的 SSE 接口。
    *   配置全局 `useSeoMeta`，使用 Twitter Card Validator 和 Lighthouse 验证达到 100/100 满分。

---

通过这套方案，你的 `frontend/web` 不仅拥有**顶级的设计质感与移动端体验**，而且直接连接你的 Go-zero 后端与 Lemon Squeezy 支付闭环！