-- 为 hg_th_coupon 表添加预约相关字段
-- 创建时间: 2025-10-18
-- 说明: 添加是否需要预约字段和需要预约的餐厅IDs字段

-- 添加是否需要预约字段
ALTER TABLE `hg_th_coupon` 
ADD COLUMN `need_reservation` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否需要预约：0-不需要，1-需要' AFTER `sort`;

-- 添加需要预约的餐厅IDs字段
ALTER TABLE `hg_th_coupon` 
ADD COLUMN `reservation_restaurant_ids` TEXT NULL COMMENT '需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3' AFTER `need_reservation`;

-- 添加索引以提高查询性能
ALTER TABLE `hg_th_coupon` 
ADD INDEX `idx_need_reservation` (`need_reservation`);

-- 查看表结构确认字段已添加
-- DESCRIBE `hg_th_coupon`;

-- 示例数据更新（可选，根据实际需求调整）
-- UPDATE `hg_th_coupon` SET `need_reservation` = 1, `reservation_restaurant_ids` = '1,2,3' WHERE `id` = 1;


-- 添加价格是否是礼品券兑换专属字段
ALTER TABLE `hg_food_goods` 
ADD COLUMN `is_th_coupon_exclusive` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否是礼品券兑换专属：0-否，1-是' AFTER `pay_order_amount`;

-- 添加绑定的礼品券ID字段
ALTER TABLE `hg_food_goods` 
ADD COLUMN `th_coupon_id` INT(11) NULL DEFAULT NULL COMMENT '绑定的礼品券ID，用于礼品券兑换专属商品' AFTER `is_th_coupon_exclusive`;

-- 添加索引以提高查询性能
ALTER TABLE `hg_food_goods` 
ADD INDEX `idx_is_th_coupon_exclusive` (`is_th_coupon_exclusive`);

-- 添加礼品券ID索引
ALTER TABLE `hg_food_goods` 
ADD INDEX `idx_th_coupon_id` (`th_coupon_id`);

-- 添加复合索引，便于查询特定礼品券的专属商品
ALTER TABLE `hg_food_goods` 
ADD INDEX `idx_coupon_exclusive_coupon_id` (`is_th_coupon_exclusive`, `th_coupon_id`);

-- 示例数据更新（可选，根据实际需求调整）
-- 将某个商品设置为特定礼品券的专属商品
-- UPDATE `hg_food_goods` SET `is_th_coupon_exclusive` = 1, `th_coupon_id` = 1 WHERE `id` = 10;

-- 查询示例SQL
-- 查询特定礼品券的专属商品
-- SELECT * FROM `hg_food_goods` WHERE `th_coupon_id` = 1 AND `is_th_coupon_exclusive` = 1;

-- 查询所有礼品券专属商品
-- SELECT * FROM `hg_food_goods` WHERE `is_th_coupon_exclusive` = 1;

-- food_order添加first_name和last_name字段
ALTER TABLE `hg_food_order` 
ADD COLUMN `first_name` VARCHAR(255) NULL DEFAULT NULL COMMENT '订单预定人姓' AFTER `booking_name`;

ALTER TABLE `hg_food_order` 
ADD COLUMN `last_name` VARCHAR(255) NULL DEFAULT NULL COMMENT '订单预定人名' AFTER `first_name`;


-- food_restaurant添加order_mode预定模式字段
ALTER TABLE `hg_food_restaurant` 
ADD COLUMN `order_mode` enum('ALLAMOUNT','DEPOSIT') NOT NULL DEFAULT 'DEPOSIT' COMMENT '预定模式' AFTER `id`;

-- 更新hg_food_order的order_type字段 加一项CRSALL
ALTER TABLE `hg_food_order` 
MODIFY COLUMN `order_type` enum('CRS','TORETA','CRSALL') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'CRS' COMMENT '订单类型' AFTER `id`;