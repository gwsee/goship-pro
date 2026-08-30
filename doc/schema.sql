-- =============================================================================
-- Project: GoShip Pro (Go-zero + Nuxt 3 Full-Stack Commercial Starter Kit)
-- Database: MySQL 8.0+
-- File: doc/schema.sql
-- Description: Production-ready DDL with initial seed data.
-- =============================================================================

CREATE DATABASE IF NOT EXISTS `goship_db`
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE `goship_db`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- -----------------------------------------------------------------------------
-- 1. 用户表 (users)
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户主键ID',
                         `email` varchar(128) NOT NULL COMMENT '登录邮箱(唯一)',
                         `password_hash` varchar(255) NOT NULL DEFAULT '' COMMENT 'Bcrypt加密密码(OAuth用户可为空)',
                         `role` varchar(32) NOT NULL DEFAULT 'user' COMMENT '角色权限: user(普通用户) / admin(超级管理员)',
                         `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像CDN链接',
                         `credits` int NOT NULL DEFAULT 10 COMMENT 'AI调用点数/可用额度',
                         `google_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Google OAuth唯一用户标识',
                         `metadata` json DEFAULT NULL COMMENT '扩展字段: 用户画像、偏好设置等',
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uk_email` (`email`),
                         KEY `idx_google_id` (`google_id`),
                         KEY `idx_role` (`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户基础信息表';

-- -----------------------------------------------------------------------------
-- 2. 商品表 (products)
-- 支持: SaaS订阅、AI点数包(虚拟) 与 极客硬件、按需定制、独立站商品(实物)
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品主键ID',
                            `title` varchar(128) NOT NULL COMMENT '商品展示名称',
                            `slug` varchar(128) NOT NULL COMMENT 'SEO友好URL标识(唯一,如 pro-plan 或 cyberpunk-clock)',
                            `description` text COMMENT '商品详情介绍(支持Markdown或HTML)',
                            `cover_image` varchar(255) NOT NULL DEFAULT '' COMMENT '商品主图(Cloudflare R2 CDN地址)',
                            `price_cents` int unsigned NOT NULL COMMENT '售价(以美分为单位, 如 2900 = $29.00)',
                            `product_type` varchar(32) NOT NULL DEFAULT 'digital' COMMENT '商品属性: digital(虚拟SaaS/点数) / physical(实物/周边)',
                            `stock_count` int NOT NULL DEFAULT -1 COMMENT '库存数量: -1表示无限库存(虚拟品), >=0表示实物库存',
                            `lemon_variant_id` varchar(64) NOT NULL DEFAULT '' COMMENT '绑定的Lemon Squeezy收银台商品Variant ID',
                            `credits_reward` int NOT NULL DEFAULT 0 COMMENT '购买后自动赠送的AI点数(仅对虚拟品生效)',
                            `status` varchar(32) NOT NULL DEFAULT 'draft' COMMENT '上下架状态: draft(草稿) / active(上架) / archived(下架)',
                            `sort_order` int NOT NULL DEFAULT 0 COMMENT '展示排序权重(越大越靠前)',
                            `metadata` json DEFAULT NULL COMMENT '扩展字段: 实体规格(颜色/尺寸)、SEO关键词、多语言标签等',
                            `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `uk_slug` (`slug`),
                            KEY `idx_status_sort` (`status`, `sort_order`),
                            KEY `idx_lemon_variant` (`lemon_variant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品与服务表';

-- -----------------------------------------------------------------------------
-- 3. 交易与发货订单表 (orders)
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单主键ID',
                          `order_no` varchar(64) NOT NULL COMMENT '系统内部业务订单号(全局唯一)',
                          `user_id` bigint unsigned NOT NULL COMMENT '购买者用户ID',
                          `product_id` bigint unsigned NOT NULL COMMENT '关联商品ID',
                          `lemon_order_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Lemon Squeezy官方交易订单ID',
                          `amount_cents` int unsigned NOT NULL COMMENT '实付总金额(美分)',
                          `currency` varchar(8) NOT NULL DEFAULT 'USD' COMMENT '结算币种(默认USD)',
                          `payment_status` varchar(32) NOT NULL DEFAULT 'pending' COMMENT '支付状态: pending(待支付) / paid(已支付) / refunded(已退款)',
                          `fulfillment_status` varchar(32) NOT NULL DEFAULT 'unfulfilled' COMMENT '履约发货状态: unfulfilled(待发货/待开通) / fulfilled(已完成)',
                          `shipping_info` json DEFAULT NULL COMMENT '买家收件信息JSON(姓名、国家代码、详细街道、城市、州/省、邮编、电话)',
                          `tracking_number` varchar(128) NOT NULL DEFAULT '' COMMENT '国际物流快递单号(实物订单发货填写)',
                          `customer_note` text COMMENT '买家下单备注',
                          `metadata` json DEFAULT NULL COMMENT '扩展字段: Lemon Squeezy原始Webhook Payload镜像备份',
                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
                          `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `uk_order_no` (`order_no`),
                          KEY `idx_user_id` (`user_id`),
                          KEY `idx_product_id` (`product_id`),
                          KEY `idx_lemon_order_id` (`lemon_order_id`),
                          KEY `idx_payment_status` (`payment_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交易与履约订单表';

-- -----------------------------------------------------------------------------
-- 4. 全局系统配置表 (system_configs)
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `system_configs`;
CREATE TABLE `system_configs` (
                                  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '配置主键ID',
                                  `config_key` varchar(64) NOT NULL COMMENT '配置键名(唯一标识)',
                                  `config_value` text NOT NULL COMMENT '配置内容值',
                                  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '配置项说明描述',
                                  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uk_config_key` (`config_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='全局轻量运行时配置表';


-- =============================================================================
-- 初始化测试数据 (Seed Data)
-- 方便本地启动后立刻有数据调试前台与后台
-- =============================================================================

-- 1. 插入默认全局配置
INSERT INTO `system_configs` (`config_key`, `config_value`, `description`) VALUES
                                                                               ('site_name', 'GoShip Pro', '全站品牌展示名称'),
                                                                               ('default_free_credits', '10', '新用户注册默认赠送的免费AI调用点数'),
                                                                               ('enable_registration', 'true', '是否开放前台新用户注册(true/false)'),
                                                                               ('ai_default_model', 'deepseek-chat', '默认调用的大语言模型名称'),
                                                                               ('contact_email', 'support@yourdomain.com', '官方客服联络邮箱');

-- 2. 插入初始示例商品
INSERT INTO `products` (`id`, `title`, `slug`, `description`, `cover_image`, `price_cents`, `product_type`, `stock_count`, `lemon_variant_id`, `credits_reward`, `status`, `sort_order`, `metadata`) VALUES
                                                                                                                                                                                                         (1, 'GoShip Pro Boilerplate License', 'goship-pro-license', 'Get full access to Go-zero + Nuxt 3 source code, VIP Discord support, and lifetime updates.', 'https://images.unsplash.com/photo-1618401471353-b98afee0b2eb?w=800', 14900, 'digital', -1, 'variant_test_001', 500, 'active', 100, '{"badge": "BEST VALUE", "features": ["Full Source Code", "Docker Setup", "Lifetime Updates"]}'),
                                                                                                                                                                                                         (2, 'Starter Credit Pack (1000 AI Credits)', 'starter-credits-pack', 'Top up 1000 credits to generate viral marketing copy and perform automated tasks.', 'https://images.unsplash.com/photo-1677442136019-21780ecad995?w=800', 1900, 'digital', -1, 'variant_test_002', 1000, 'active', 90, '{"badge": "POPULAR", "features": ["1000 Generation Credits", "Instant Activation", "No Expiration"]}'),
                                                                                                                                                                                                         (3, 'Cyberpunk RGB Nixie Tube Clock', 'cyberpunk-nixie-clock', 'A futuristic desktop glowing digital clock with custom ESP32 firmware and RGB backlights.', 'https://images.unsplash.com/photo-1550745165-9bc0b252726f?w=800', 6900, 'physical', 50, 'variant_test_003', 0, 'active', 80, '{"weight_g": 450, "origin": "CN", "features": ["ESP32 Inside", "Full RGB Glow", "Type-C Powered"]}');

-- 3. 插入初始管理员用户
-- 默认密码为: admin123 (已通过 Bcrypt 加密)
INSERT INTO `users` (`id`, `email`, `password_hash`, `role`, `avatar_url`, `credits`) VALUES
    (1, 'admin@goship.dev', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin', 'https://api.dicebear.com/7.x/bottts/svg?seed=goship', 99999);

SET FOREIGN_KEY_CHECKS = 1;