# 员工福利系统设计文档
# Employee Benefits System Design Document

## 1. 系统概述 (System Overview)

员工福利系统是一个用于管理企业员工福利发放的综合性系统，主要包括员工管理、员工部门管理、活动管理以及礼品券发放等功能模块。

### 1.1 主要功能模块

- **员工管理**: 员工信息的增删改查
- **员工部门**: 多级部门结构的增删改查  
- **活动管理**: 福利活动的创建和管理
- **礼品券关联**: 活动与礼品券的关联管理
- **领取记录**: 员工礼品券领取记录跟踪

## 2. 数据库设计 (Database Design)

### 2.1 表结构设计

#### 2.1.1 员工管理表 (`hg_employee`)

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | bigint(20) | 员工ID | 主键，自增 |
| name | varchar(100) | 员工姓名 | 非空 |
| phone_area | varchar(10) | 电话区号 | 默认'+86' |
| phone | varchar(20) | 电话号码 | 非空 |
| member_id | bigint(20) | 绑定用户ID | 可空 |
| department_id | bigint(20) | 员工部门ID | 非空 |
| employee_no | varchar(50) | 员工编号 | 唯一 |
| status | tinyint(4) | 状态 | 1-正常 2-禁用 |
| remark | text | 备注 | 可空 |
| created_at | datetime | 创建时间 | 默认当前时间 |
| updated_at | datetime | 更新时间 | 自动更新 |

**索引设计:**
- 主键: `id`
- 唯一索引: `employee_no`, `phone_area + phone`
- 普通索引: `department_id`, `member_id`, `status`
- 索引: `idx_department_status` (`department_id`, `status`);

#### 2.1.2 员工部门表 (`hg_employee_department`)

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | bigint(20) | 部门ID | 主键，自增 |
| name | varchar(100) | 部门名称 | 非空 |
| parent_id | bigint(20) | 上级部门ID | 可空，NULL表示顶级部门 |
| level | int(11) | 部门层级 | 1表示顶级 |
| path | varchar(500) | 部门路径 | 如：1,2,3 |
| manager_id | bigint(20) | 部门负责人 ID | 可空 |
| description | text | 部门描述 | 可空 |
| sort | int(11) | 排序 | 默认0 |
| status | tinyint(4) | 状态 | 1-正常 2-禁用 |
| created_at | datetime | 创建时间 | 默认当前时间 |
| updated_at | datetime | 更新时间 | 自动更新 |

**索引设计:**
- 主键: `id`
- 普通索引: `parent_id`, `level`, `status`, `sort`, `manager_id`

#### 2.1.3 活动管理表 (`hg_employee_activity`)

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | bigint(20) | 活动ID | 主键，自增 |
| name | varchar(255) | 活动名称（多语言） | 非空 |
| cover | varchar(500) | 活动封面 | 可空 |
| description | text | 活动描述（多语言） | 可空 |
| validity_type | tinyint(4) | 有效期类型 | 1-指定时间 2-长期有效 |
| start_time | datetime | 开始时间 | 可空 |
| end_time | datetime | 结束时间 | 可空 |
| status | tinyint(4) | 状态 | 1-未开始 2-进行中 3-已结束 |
| manual_closed | tinyint(4) | 是否手动关闭 | 0-否 1-是 |
| restriction_type | tinyint(4) | 限制类型 | 1-不做任何限制 2-限制指定部门 3-限制指定员工 |
| rule | varchar(500) | 活动规则 | 可空 |
| sort | int(11) | 排序 | 默认0 |
| remark | text | 备注 | 可空 |
| is_enabled | tinyint(4) | 状态 | 1-正常 2-禁用 |
| created_at | datetime | 创建时间 | 默认当前时间 |
| updated_at | datetime | 更新时间 | 自动更新 |

**索引设计:**
- 主键: `id`
- 普通索引: `status`, `manual_closed`, `restriction_type`, `validity_type`, `start_time + end_time`, `sort`
- 索引: `idx_validity_status` (`validity_type`, `status`);
- 索引: `idx_time_status` (`start_time`, `end_time`, `status`);
- 索引: `idx_restriction_status` (`restriction_type`, `status`);

#### 2.1.4 活动部门关联表 (`hg_employee_activity_department`)

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|---------|
| id | bigint(20) | 关联ID | 主键，自增 |
| activity_id | bigint(20) | 活动ID | 非空 |
| department_id | bigint(20) | 部门ID | 非空 |
| created_at | datetime | 创建时间 | 默认当前时间 |

**索引设计:**
- 主键: `id`
- 唯一索引: `activity_id + department_id`
- 普通索引: `activity_id`, `department_id`

#### 2.1.5 活动员工关联表 (`hg_employee_activity_employee`)

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|---------|
| id | bigint(20) | 关联ID | 主键，自增 |
| activity_id | bigint(20) | 活动ID | 非空 |
| employee_id | bigint(20) | 员工ID | 非空 |
| created_at | datetime | 创建时间 | 默认当前时间 |

**索引设计:**
- 主键: `id`
- 唯一索引: `activity_id + employee_id`
- 普通索引: `activity_id`, `employee_id`

#### 2.1.6 活动礼品券关联表 (`hg_employee_activity_coupon`)

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| id | bigint(20) | 关联ID | 主键，自增 |
| activity_id | bigint(20) | 活动ID | 非空 |
| coupon_id | bigint(20) | 礼品券ID | 非空 |
| available_quantity | int(11) | 可领取数量 | 0表示无限制 |
| total_received | int(11) | 总领取数量 | 默认0 |
| total_used | int(11) | 已核销数量 | 默认0 |
| status | tinyint(4) | 状态 | 1-正常 2-禁用 |
| created_at | datetime | 创建时间 | 默认当前时间 |
| updated_at | datetime | 更新时间 | 自动更新 |

**索引设计:**
- 主键: `id`
- 唯一索引: `activity_id + coupon_id`
- 普通索引: `activity_id`, `coupon_id`, `status`

#### 2.1.7 会员礼品券表扩展 (`hg_th_member_coupon`)

**新增字段说明:**

| 字段名 | 类型 | 说明 | 约束 |
|--------|------|------|------|
| activity_id | bigint(20) | 活动ID | 可空，员工福利活动关联 |
| employee_id | bigint(20) | 员工ID | 可空，员工福利发放关联 |

**扩展后的索引:**
- 新增索引: `idx_activity_id`, `idx_employee_id`
- 建议复合索引: `idx_employee_activity` (`employee_id`, `activity_id`)
- 建议复合索引: `idx_activity_source` (`activity_id`, `source`)

**说明:** 复用现有的 `hg_th_member_coupon` 表作为员工礼品券领取记录表，通过新增的 `activity_id` 和 `employee_id` 字段来支持员工福利系统。当这两个字段有值时，表示该记录来源于员工福利活动。

### 2.2 表关系设计

```
hg_employee_department (1) -----> (N) hg_employee_department (父子关系)
hg_employee_department (1) -----> (N) hg_employee (通过department_id)
hg_employee_activity (1) -----> (N) hg_employee_activity_department
hg_employee_department (1) -----> (N) hg_employee_activity_department
hg_employee_activity (1) -----> (N) hg_employee_activity_employee
hg_employee (1) -----> (N) hg_employee_activity_employee
hg_employee_activity (1) -----> (N) hg_employee_activity_coupon
hg_th_coupon (1) -----> (N) hg_employee_activity_coupon
hg_employee (1) -----> (N) hg_th_member_coupon (通过employee_id)
hg_employee_activity (1) -----> (N) hg_th_member_coupon (通过activity_id)
hg_th_coupon (1) -----> (N) hg_th_member_coupon (通过coupon_id)
```

## 3. 业务流程 (Business Process)

### 3.1 员工管理流程

1. **新增员工**: 录入员工基本信息，选择部门
2. **编辑员工**: 修改员工信息，可调整部门
3. **删除员工**: 软删除或硬删除员工记录
4. **查询员工**: 支持按部门、状态等条件查询

### 3.2 部门管理流程

1. **创建部门**: 设置部门名称、上级部门和描述
2. **编辑部门**: 修改部门信息，调整部门层级
3. **删除部门**: 删除前需确保部门下无员工和子部门
4. **查询部门**: 支持按层级、上级部门查询，展示部门树结构

### 3.3 活动管理流程

1. **创建活动**: 设置活动基本信息和有效期
2. **选择参与范围**: 可选择不做限制、限制指定部门或限制指定员工
3. **配置礼品券**: 为活动关联礼品券，设置数量限制
4. **发布活动**: 激活活动，符合条件的员工可开始领取
5. **结束活动**: 手动或自动结束活动

### 3.4 礼品券领取流程

1. **员工查看活动**: 浏览适用于自己的活动列表
2. **权限验证**: 系统验证员工是否符合活动参与条件（部门或个人限制）
3. **选择礼品券**: 从活动中选择要领取的礼品券
4. **确认领取**: 系统验证数量限制后发放礼品券
5. **使用礼品券**: 在指定商户核销使用

## 4. 数据统计 (Data Statistics)

### 4.1 核心统计指标

- **活动参与率**: 参与活动的员工数 / 总员工数
- **礼品券领取率**: 已领取礼品券数 / 可领取礼品券总数
- **礼品券使用率**: 已使用礼品券数 / 已领取礼品券数
- **部门活跃度**: 各部门员工的参与情况统计

### 4.2 统计查询示例

```sql
-- 活动参与统计
SELECT 
    a.name AS activity_name,
    COUNT(DISTINCT mc.employee_id) AS participant_count,
    COUNT(mc.id) AS total_received,
    SUM(CASE WHEN mc.state = 3 THEN 1 ELSE 0 END) AS total_used
FROM hg_employee_activity a
LEFT JOIN hg_th_member_coupon mc ON a.id = mc.activity_id
GROUP BY a.id, a.name;

-- 部门参与统计
SELECT 
    d.name AS department_name,
    d.level AS department_level,
    COUNT(DISTINCT e.id) AS total_employees,
    COUNT(DISTINCT mc.employee_id) AS active_employees,
    COUNT(mc.id) AS total_received
FROM hg_employee_department d
LEFT JOIN hg_employee e ON d.id = e.department_id
LEFT JOIN hg_th_member_coupon mc ON e.id = mc.employee_id AND mc.activity_id IS NOT NULL
GROUP BY d.id, d.name, d.level
ORDER BY d.level, d.sort;
```

## 5. 系统扩展 (System Extension)

### 5.1 可扩展功能

1. **权限管理**: 不同角色的操作权限控制
2. **消息通知**: 活动发布、礼品券到期提醒
3. **数据导出**: 员工信息、活动统计数据导出
4. **API接口**: 提供移动端或第三方系统集成接口
5. **审批流程**: 大额礼品券发放需要审批

### 5.2 性能优化建议

1. **数据分区**: 按时间对大表进行分区
2. **缓存策略**: 活动信息、员工信息等热点数据缓存
3. **异步处理**: 大批量礼品券发放使用队列异步处理
4. **读写分离**: 查询统计使用只读从库

## 6. 安全考虑 (Security Considerations)

### 6.1 数据安全

- 员工手机号等敏感信息加密存储
- 操作日志记录，支持审计追踪
- 数据备份和恢复机制

### 6.2 业务安全

- 防止重复领取同一礼品券
- 礼品券数量限制和库存控制
- 异常操作监控和告警

## 7. 部署说明 (Deployment Instructions)

### 7.1 数据库初始化

```bash
# 执行SQL脚本创建表结构
mysql -u username -p database_name < employee_benefits_schema.sql
```

### 7.2 配置说明

- 确保数据库字符集为 `utf8mb4`
- 建议使用InnoDB存储引擎
- 根据业务量调整相关索引

### 7.3 监控建议

- 监控表空间使用情况
- 监控慢查询日志
- 定期检查数据一致性
