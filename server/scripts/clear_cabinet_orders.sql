-- 储物柜订单数据清理脚本
-- 警告：此脚本将清空所有储物柜订单数据，包括订单表、日志表和相关业务表
-- 执行前请确保已备份重要数据

-- 开始事务
START TRANSACTION;

-- 1. 清空储物柜相关的支付交易退款记录
DELETE FROM hg_pms_transaction_refund WHERE order_sn IN (SELECT order_sn FROM hg_cabinet_order);

-- 2. 清空储物柜相关的支付交易记录
DELETE FROM hg_pms_transaction WHERE scene = 'CABINET';

-- 3. 清空储物柜相关的会员余额变动记录
DELETE FROM hg_pms_balance_change WHERE scene = 'CABINET';

-- 4. 清空储物柜相关的优惠券记录
DELETE FROM hg_pms_coupon WHERE scene = 5;

-- 5. 清空储物柜订单导出记录
DELETE FROM hg_order_export WHERE scene = 5;

-- 6. 清空储物柜订单日志表
TRUNCATE TABLE hg_cabinet_order_log;

-- 7. 清空储物柜订单表  
TRUNCATE TABLE hg_cabinet_order;

-- 提交事务
COMMIT;

-- 显示清理结果
SELECT 
    'hg_cabinet_order' as table_name,
    COUNT(*) as record_count
FROM hg_cabinet_order
UNION ALL
SELECT 
    'hg_cabinet_order_log' as table_name,
    COUNT(*) as record_count  
FROM hg_cabinet_order_log;

-- 显示清理完成信息
SELECT '储物柜订单数据清理完成！' as message;
