-- ===============================================
-- 员工福利系统数据库设计
-- Employee Benefits System Database Schema
-- ===============================================

-- 1. 员工管理表 (Employee Management)
CREATE TABLE `hg_employee` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '员工ID',
  `name` varchar(100) NOT NULL COMMENT '员工姓名',
  `phone_area` varchar(10) NOT NULL DEFAULT '+86' COMMENT '电话区号',
  `phone` varchar(20) NOT NULL COMMENT '电话号码',
  `member_id` bigint(20) unsigned DEFAULT NULL COMMENT '绑定用户ID',
  `department_id` bigint(20) unsigned NOT NULL COMMENT '员工部门ID',
  `employee_no` varchar(50) NOT NULL COMMENT '员工编号',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-正常 2-禁用',
  `remark` text COMMENT '备注',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_employee_no` (`employee_no`),
  UNIQUE KEY `uk_phone` (`phone_area`, `phone`),
  KEY `idx_department_id` (`department_id`),
  KEY `idx_member_id` (`member_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='员工管理表';

-- 2. 员工部门表 (Employee Department) - 支持多级部门结构
CREATE TABLE `hg_employee_department` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` varchar(100) NOT NULL COMMENT '部门名称',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '上级部门ID（NULL表示顶级部门）',
  `level` int(11) NOT NULL DEFAULT '1' COMMENT '部门层级（1表示顶级）',
  `path` varchar(500) NOT NULL DEFAULT '' COMMENT '部门路径（如：1,2,3）',
  `manager_id` bigint(20) unsigned DEFAULT NULL COMMENT '部门负责人ID',
  `description` text COMMENT '部门描述',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-正常 2-禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_level` (`level`),
  KEY `idx_status` (`status`),
  KEY `idx_sort` (`sort`),
  KEY `idx_manager_id` (`manager_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='员工部门表';

-- 3. 活动管理表 (Activity Management)
CREATE TABLE `hg_employee_activity` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '活动ID',
  `name` varchar(255) NOT NULL COMMENT '活动名称（多语言）',
  `cover` varchar(500) DEFAULT NULL COMMENT '活动封面',
  `description` text COMMENT '活动描述（多语言）',
  `validity_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '有效期类型：1-指定时间段 2-长期有效',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-未开始 2-进行中 3-已结束',
  `manual_closed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否手动关闭：0-否 1-是',
  `restriction_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '限制类型：1-不做任何限制 2-限制指定部门 3-限制指定员工',
  `rule` varchar(500) DEFAULT NULL COMMENT '活动规则',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` text COMMENT '备注',
  `is_enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用：1-正常 2-禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_manual_closed` (`manual_closed`),
  KEY `idx_restriction_type` (`restriction_type`),
  KEY `idx_validity_type` (`validity_type`),
  KEY `idx_start_end_time` (`start_time`, `end_time`),
  KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='员工活动表';

-- 4. 活动部门关联表 (Activity Department Relation)
CREATE TABLE `hg_employee_activity_department` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关联ID',
  `activity_id` bigint(20) unsigned NOT NULL COMMENT '活动ID',
  `department_id` bigint(20) unsigned NOT NULL COMMENT '部门ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_activity_department` (`activity_id`, `department_id`),
  KEY `idx_activity_id` (`activity_id`),
  KEY `idx_department_id` (`department_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动部门关联表';

-- 5. 活动员工关联表 (Activity Employee Relation)
CREATE TABLE `hg_employee_activity_employee` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关联ID',
  `activity_id` bigint(20) unsigned NOT NULL COMMENT '活动ID',
  `employee_id` bigint(20) unsigned NOT NULL COMMENT '员工ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_activity_employee` (`activity_id`, `employee_id`),
  KEY `idx_activity_id` (`activity_id`),
  KEY `idx_employee_id` (`employee_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动员工关联表';

-- 6. 活动礼品券关联表 (Activity Coupon Relation)
CREATE TABLE `hg_employee_activity_coupon` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关联ID',
  `activity_id` bigint(20) unsigned NOT NULL COMMENT '活动ID',
  `coupon_id` bigint(20) unsigned NOT NULL COMMENT '礼品券ID',
  `available_quantity` int(11) NOT NULL DEFAULT '0' COMMENT '可领取数量（0表示无限制）',
  `total_received` int(11) NOT NULL DEFAULT '0' COMMENT '总领取数量',
  `total_used` int(11) NOT NULL DEFAULT '0' COMMENT '已核销数量',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-正常 2-禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_activity_coupon` (`activity_id`, `coupon_id`),
  KEY `idx_activity_id` (`activity_id`),
  KEY `idx_coupon_id` (`coupon_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动礼品券关联表';

-- 7. 修改现有会员礼品券表，添加员工福利相关字段 (Modify existing member coupon table)
ALTER TABLE `hg_th_member_coupon` 
ADD COLUMN `activity_id` bigint(20) unsigned DEFAULT NULL COMMENT '活动ID（员工福利活动）' AFTER `source_order_id`,
ADD COLUMN `employee_id` bigint(20) unsigned DEFAULT NULL COMMENT '员工ID（员工福利发放）' AFTER `activity_id`,
ADD INDEX `idx_activity_id` (`activity_id`),
ADD INDEX `idx_employee_id` (`employee_id`);

-- ===============================================
-- 初始化数据
-- ===============================================

-- 插入默认部门数据
INSERT INTO `hg_employee_department` (`id`, `name`, `parent_id`, `level`, `path`, `description`, `sort`, `status`) VALUES
(1, '总公司', NULL, 1, '1', '公司总部', 1, 1),
(2, '技术部', 1, 2, '1,2', '技术研发部门', 1, 1),
(3, '市场部', 1, 2, '1,3', '市场营销部门', 2, 1),
(4, '人事部', 1, 2, '1,4', '人力资源部门', 3, 1),
(5, '前端组', 2, 3, '1,2,5', '前端开发组', 1, 1),
(6, '后端组', 2, 3, '1,2,6', '后端开发组', 2, 1),
(7, '测试组', 2, 3, '1,2,7', '软件测试组', 3, 1);


-- ===============================================
-- 索引优化建议
-- ===============================================

-- 为了提高查询性能，建议根据实际业务场景添加以下复合索引：

-- 员工表复合索引
-- ALTER TABLE `hg_employee` ADD INDEX `idx_department_status` (`department_id`, `status`);
-- ALTER TABLE `hg_employee` ADD INDEX `idx_member_status` (`member_id`, `status`);

-- 活动表复合索引  
-- ALTER TABLE `hg_employee_activity` ADD INDEX `idx_validity_status` (`validity_type`, `status`);
-- ALTER TABLE `hg_employee_activity` ADD INDEX `idx_time_status` (`start_time`, `end_time`, `status`);
-- ALTER TABLE `hg_employee_activity` ADD INDEX `idx_restriction_status` (`restriction_type`, `status`);

-- 会员礼品券表复合索引（员工福利相关）
-- ALTER TABLE `hg_th_member_coupon` ADD INDEX `idx_employee_activity` (`employee_id`, `activity_id`);
-- ALTER TABLE `hg_th_member_coupon` ADD INDEX `idx_activity_source` (`activity_id`, `source`);

-- ===============================================
-- 外键约束（可选）
-- ===============================================

-- 如果需要严格的数据完整性，可以添加外键约束：
-- ALTER TABLE `hg_employee` ADD CONSTRAINT `fk_employee_department` FOREIGN KEY (`department_id`) REFERENCES `hg_employee_department` (`id`);
-- ALTER TABLE `hg_employee_activity_department` ADD CONSTRAINT `fk_activity_dept_activity` FOREIGN KEY (`activity_id`) REFERENCES `hg_employee_activity` (`id`);
-- ALTER TABLE `hg_employee_activity_department` ADD CONSTRAINT `fk_activity_dept_department` FOREIGN KEY (`department_id`) REFERENCES `hg_employee_department` (`id`);
-- ALTER TABLE `hg_employee_activity_employee` ADD CONSTRAINT `fk_activity_emp_activity` FOREIGN KEY (`activity_id`) REFERENCES `hg_employee_activity` (`id`);
-- ALTER TABLE `hg_employee_activity_employee` ADD CONSTRAINT `fk_activity_emp_employee` FOREIGN KEY (`employee_id`) REFERENCES `hg_employee` (`id`);
-- ALTER TABLE `hg_employee_activity_coupon` ADD CONSTRAINT `fk_activity_coupon_activity` FOREIGN KEY (`activity_id`) REFERENCES `hg_employee_activity` (`id`);
-- ALTER TABLE `hg_th_member_coupon` ADD CONSTRAINT `fk_member_coupon_employee` FOREIGN KEY (`employee_id`) REFERENCES `hg_employee` (`id`);
-- ALTER TABLE `hg_th_member_coupon` ADD CONSTRAINT `fk_member_coupon_activity` FOREIGN KEY (`activity_id`) REFERENCES `hg_employee_activity` (`id`);
