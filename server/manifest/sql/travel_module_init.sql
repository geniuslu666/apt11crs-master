-- 一日游模块初始化 SQL
-- 创建时间: 2026-03-09

-- ----------------------------
-- 产品表
-- ----------------------------
CREATE TABLE `hg_travel_product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题（默认语言；多语言存 hg_pms_language）',
  `sub_title` varchar(255) NOT NULL DEFAULT '' COMMENT '副标题（默认语言；多语言存 hg_pms_language）',
  `content_zh` text COMMENT '内容（中文 编辑器）',
  `content_en` text COMMENT '内容（英文 编辑器）',
  `content_ko` text COMMENT '内容（韩文 编辑器）',
  `content_ja` text COMMENT '内容（日文 编辑器）',
  `content_tw` text COMMENT '内容（繁文 编辑器）',
  `list_image` varchar(500) NOT NULL DEFAULT '' COMMENT '列表图（单图 URL）',
  `carousel_images` text COMMENT '轮播图（多图，JSON）',
  `daily_capacity` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '每日最大接待人数',
  `stock` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总库存',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '售价（元）',
  `contact_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '联系电话',
  `meeting_place` varchar(500) NOT NULL DEFAULT '' COMMENT '集合地点',
  `meeting_time` varchar(5) NOT NULL DEFAULT '' COMMENT '集合时间（格式：HH:MM）',
  `gg_lat` varchar(255) DEFAULT NULL COMMENT '谷歌纬度',
  `gg_lng` varchar(255) DEFAULT NULL COMMENT '谷歌经度',
  `max_book_days` int(10) unsigned NOT NULL DEFAULT '30' COMMENT '最大可预约天数',
  `advance_book_days` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '至少提前预订天数',
  `trip_planning_zh` text COMMENT '行程规划（中文 编辑器）',
  `trip_planning_en` text COMMENT '行程规划（英文 编辑器）',
  `trip_planning_ja` text COMMENT '行程规划（日文 编辑器）',
  `trip_planning_ko` text COMMENT '行程规划（韩文 编辑器）',
  `trip_planning_tw` text COMMENT '行程规划（繁文 编辑器）',
  `booking_notes_zh` text COMMENT '预约须知（中文 编辑器）',
  `booking_notes_en` text COMMENT '预约须知（英文 编辑器）',
  `booking_notes_ja` text COMMENT '预约须知（日文 编辑器）',
  `booking_notes_ko` text COMMENT '预约须知（韩文 编辑器）',
  `booking_notes_tw` text COMMENT '预约须知（繁文 编辑器）',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态（1启用 2禁用）',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序（越大越靠前）',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间（NULL=正常）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_sort` (`sort`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='一日游产品表';

CREATE TABLE `hg_travel_product_sku` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '产品ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '车型名称（默认语言；多语言存 hg_pms_language）',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '售价（元）',
  `daily_capacity` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '每日最大接待人数',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态（1启用 2禁用）',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序（越大越靠前）',
  `sales_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '已售',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间（NULL=正常）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_sort` (`sort`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='一日游产品sku表';

-- ----------------------------
-- 订单表
-- ----------------------------
CREATE TABLE `hg_travel_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_sn` varchar(64) NOT NULL COMMENT '预约单号',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '产品ID',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'SKUID',
  `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '会员ID',
  `booking_name` varchar(100) NOT NULL DEFAULT '' COMMENT '预订人姓名',
  `first_name` varchar(255) DEFAULT NULL COMMENT '订单预定人名',
  `last_name` varchar(255) DEFAULT NULL COMMENT '订单预定人姓',
  `phone_area` varchar(10) NOT NULL DEFAULT '' COMMENT '手机区号',
  `booking_mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '预订人电话',
  `booking_email` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '预定人邮箱',
  `booking_num` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '预约人数',
  `book_date` date NOT NULL COMMENT '预约日期',
  `order_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单金额（元）',
  `coupon_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠券抵扣金额',
  `bal_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '积分抵扣金额',
  `order_status` enum('WAIT_PAY','WAIT_VERIFY','DONE','CANCEL','REFUND','OVERDUE') NOT NULL DEFAULT 'WAIT_PAY' COMMENT '订单状态',
  `pay_model` tinyint(4) NOT NULL DEFAULT '3' COMMENT '1、余额支付 2、组合支付 3、纯外部支付',
  `pay_status` enum('WAIT_PAY','HAVE_PAID','CANCEL','REFUND') NOT NULL DEFAULT 'WAIT_PAY' COMMENT '订单付款状态',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `verify_staff_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '核销人员ID',
  `verify_time` datetime DEFAULT NULL COMMENT '核销时间',
  `expiration_time` int(11) DEFAULT NULL COMMENT '订单过期时间',
  `cancel_time` datetime DEFAULT NULL COMMENT '取消时间',
  `cancel_fee` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '取消手续费（元）',
  `refund_status` enum('WAIT','PART','DONE') NOT NULL DEFAULT 'WAIT' COMMENT '退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款',
  `refund_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '退款金额（元）',
  `refund_time` datetime DEFAULT NULL COMMENT '退款时间',
  `refund_bal_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '已退款积分',
  `refund_coupon_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '已退款优惠券',
  `refund_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '退款原因',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `referrer` int(11) DEFAULT NULL COMMENT '推荐人',
  `rebate_rate` decimal(10,2) DEFAULT '0.00' COMMENT '分佣比例',
  `rebate_status` enum('WAIT','SUCCESS','FAIL') DEFAULT NULL COMMENT 'WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败',
  `rebate_amount` decimal(10,2) DEFAULT NULL COMMENT '分佣结算金额',
  `rebate_time` datetime DEFAULT NULL COMMENT '分佣结算时间',
  `is_get_open` enum('Y','N') DEFAULT 'Y' COMMENT '是否开启积分获取',
  `is_pay_open` enum('Y','N') DEFAULT 'Y' COMMENT '是否开启积分抵扣',
  `cabinet_get_rate_vip` decimal(10,2) DEFAULT NULL COMMENT '结算积分比例',
  `cabinet_get_rate_scene` decimal(10,2) DEFAULT NULL COMMENT '场景结算积分比例',
  `cabinet_get_score_status` enum('WAIT','SUCCESS','FAIL') DEFAULT 'WAIT' COMMENT '''WAIT'',''SUCCESS'',''FAIL''',
  `cabinet_get_amount` decimal(10,2) DEFAULT NULL COMMENT '结算积分金额',
  `exp_value` decimal(10,2) DEFAULT NULL COMMENT '结算的经验值',
  `exp_time` datetime DEFAULT NULL COMMENT '经验结算时间',
  `admin_refund_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '后台已退款总金额',
  `admin_refund_bal_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '后台已退款积分',
  `admin_refund_coupon_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '后台已退款优惠券',
  `admin_cancel_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '后台取消原因',
  `admin_cancel_num` int(11) NOT NULL DEFAULT '0' COMMENT '后台取消次数',
  `is_fx` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '是否是分销订单',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_sn` (`order_sn`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_order_status` (`order_status`),
  KEY `idx_book_date` (`book_date`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='一日游订单表';

-- ----------------------------
-- 核销人员表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `hg_travel_verify_staff` (
  `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
  `name`       varchar(100)    NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile`     varchar(20)     NOT NULL DEFAULT '' COMMENT '电话',
  `username`   varchar(100)    NOT NULL DEFAULT '' COMMENT '登录账号',
  `password_hash` char(32)     NOT NULL DEFAULT '' COMMENT '密码',
  `salt`       char(16)        NOT NULL COMMENT '密码盐',
  `password_reset_token` varchar(150) DEFAULT '' COMMENT '密码重置令牌',
  `status`     tinyint(1)      NOT NULL DEFAULT 1 COMMENT '状态（1启用 2禁用）',
  `created_at` datetime        DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime        DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='一日游核销人员表';

-- ----------------------------
-- 核销记录表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `hg_travel_verify_record` (
  `id`              bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id`        bigint unsigned NOT NULL DEFAULT 0 COMMENT '订单ID',
  `order_sn`        varchar(64)     NOT NULL DEFAULT '' COMMENT '预约单号',
  `product_id`      bigint unsigned NOT NULL DEFAULT 0 COMMENT '产品ID',
  `member_id`       bigint unsigned NOT NULL DEFAULT 0 COMMENT '会员ID',
  `book_date`       date            NOT NULL COMMENT '预约日期',
  `verify_staff_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '核销人员ID',
  `verify_time`     datetime        NOT NULL COMMENT '核销时间',
  `created_at`      datetime        DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_order_sn` (`order_sn`),
  KEY `idx_verify_staff_id` (`verify_staff_id`),
  KEY `idx_book_date` (`book_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='一日游核销记录表';

-- ----------------------------
-- sys_config 初始配置（取消政策）
-- ----------------------------
INSERT IGNORE INTO `hg_sys_config` (`group`, `name`, `key`, `value`, `type`) VALUES
('travelCancelPolicy', '是否允许取消', 'allowCancel',        '1',  'int'),
('travelCancelPolicy', 'xx天前', 'freeCancelDays',    '3',  'int'),
('travelCancelPolicy', '取消费率', 'cancelFeePercent',  '20', 'int');


ALTER TABLE `hg_pms_transaction`
MODIFY COLUMN `order_type` enum('BOOKING','BOOKING_CHANGE','FOOD','SPA','CAR','CABINET','TRAVEL') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'BOOKING' AFTER `change_order_sn`,
MODIFY COLUMN `scene` enum('HOTEL','FOOD','SPA','CAR','CABINET','TRAVEL') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'HOTEL' COMMENT '场景值' AFTER `capture_id`;

ALTER TABLE `hg_pms_transaction_refund`
MODIFY COLUMN `scene` enum('HOTEL','FOOD','SPA','CAR','CABINET','TRAVEL') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'HOTEL' COMMENT '场景值' AFTER `transaction_sn`;

CREATE TABLE `hg_travel_order_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '变动ID',
  `order_id` int(11) DEFAULT '0' COMMENT '订单ID',
  `order_status` enum('WAIT_PAY','WAIT_VERIFY','DONE','CANCEL','REFUND') NOT NULL DEFAULT 'WAIT_PAY' COMMENT '订单状态',
  `action_way` varchar(255) DEFAULT NULL COMMENT '操作名',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `operate_type` enum('SYSTEM','ADMIN','USER','STAFF') NOT NULL DEFAULT 'SYSTEM' COMMENT '操作员类型',
  `operate_id` int(11) DEFAULT '0' COMMENT '操作员ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `member_id` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='一日游订单日志表';

ALTER TABLE `hg_pms_member_level`
    ADD COLUMN `travel_get_rate` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 1.00 COMMENT '一日游场景获取积分倍率' AFTER `cabinet_get_rate`;

INSERT INTO `hg_pms_member_scene` (`id`, `scene_name`, `is_get_open`, `is_pay_open`, `limit_money`, `get_rate`, `pay_rate`, `is_open_reward`, `reward_type`, `reward_coupon_type_ids`, `reward_th_coupon_ids`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '一日游', 'Y', 'Y', 0.00, 1.00, 100.00, 1, '', '', '', '2024-10-31 15:39:10', '2025-08-11 16:01:19', NULL);

ALTER TABLE `hg_order_refund_log`
    MODIFY COLUMN `scene` enum('FOOD','SPA','CAR','CABINET','TRAVEL') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '场景值' AFTER `order_sn`;

-- 添加预定须知
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '预定须知-zh', 'string', 'bookingNotice_zh', '预定须知\n若迟到10分钟以上，预约可能将会被取消，恕不另行通知，迟到者不予受理退款；', '', 0, '预定须知', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '预定须知-en', 'string', 'bookingNotice_en', 'Hello\n若迟到10分钟以上，预约可能将会被取消，恕不另行通知，迟到者不予受理退款；', '', 0, '预定须知', 0, 1, NULL, '2025-10-23 11:23:30', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '预定须知-ja', 'string', 'bookingNotice_ja', '预定须知日语\n若迟到10分钟以上，预约可能将会被取消，恕不另行通知，迟到者不予受理退款；', '', 0, '预定须知', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '预定须知-ko', 'string', 'bookingNotice_ko', '안녕하세요 10분 이상 늦는 경우 사전 통보 없이 예약이 취소될 수 있습니다. 늦는 경우 환불은 제공되지 않습니다.', '', 0, '预定须知', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '预定须知-zh_CN', 'string', 'bookingNotice_zh_CN', '預定須知若遲到10分鐘以上，預約可能會被取消，恕不另行通知，遲到者不予受理退款；', '', 0, '预定须知', 0, 1, NULL, '2025-10-23 11:23:29', NULL);

-- 添加活动规则
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES('travelothersetting', '活动规则-zh', 'string', 'orderRule_zh', '活动规则-zh', '', 0, '活动规则', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES('travelothersetting', '活动规则-en', 'string', 'orderRule_en', '活动规则-en', '', 0, '活动规则', 0, 1, NULL, '2025-10-23 11:23:30', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '活动规则-ja', 'string', 'orderRule_ja', '活动规则-ja', '', 0, '活动规则', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '活动规则-ko', 'string', 'orderRule_ko', '活动规则-ko', '', 0, '活动规则', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
INSERT INTO `hg_sys_config`( `group`, `name`, `type`, `key`, `value`, `default_value`, `sort`, `tip`, `is_default`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES ('travelothersetting', '活动规则-zh_CN', 'string', 'orderRule_zh_CN', '活动规则-zh_CN', '', 0, '活动规则', 0, 1, NULL, '2025-10-23 11:23:29', NULL);
