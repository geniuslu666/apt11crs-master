package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	// 导入MySQL驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

// InitClearCabinetOrders 初始化脚本：清空储物柜订单表和储物柜订单日志表
func main() {
	var (
		ctx = gctx.New()
		err error
	)

	fmt.Println("开始执行储物柜订单数据清理脚本...")

	// 确认操作
	fmt.Print("警告：此操作将清空所有储物柜订单数据，包括订单表和日志表。是否继续？(y/N): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("操作已取消")
		return
	}

	// 开始事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 清空储物柜相关的支付交易退款记录
		fmt.Println("正在清空储物柜相关的退款交易记录...")
		_, err = tx.Exec("DELETE FROM hg_pms_transaction_refund WHERE order_sn IN (SELECT order_sn FROM hg_cabinet_order)")
		if err != nil {
			return fmt.Errorf("清空退款交易记录失败: %v", err)
		}
		fmt.Println("✓ 退款交易记录清空完成")

		// 2. 清空储物柜相关的支付交易记录
		fmt.Println("正在清空储物柜相关的支付交易记录...")
		_, err = tx.Exec("DELETE FROM hg_pms_transaction WHERE scene = 'CABINET'")
		if err != nil {
			return fmt.Errorf("清空支付交易记录失败: %v", err)
		}
		fmt.Println("✓ 支付交易记录清空完成")

		// 3. 清空储物柜相关的会员余额变动记录
		fmt.Println("正在清空储物柜相关的会员余额变动记录...")
		_, err = tx.Exec("DELETE FROM hg_pms_balance_change WHERE scene = 'CABINET'")
		if err != nil {
			return fmt.Errorf("清空会员余额变动记录失败: %v", err)
		}
		fmt.Println("✓ 会员余额变动记录清空完成")

		// 4. 清空储物柜相关的优惠券记录
		fmt.Println("正在清空储物柜相关的优惠券记录...")
		_, err = tx.Exec("DELETE FROM hg_pms_coupon WHERE scene = 5")
		if err != nil {
			return fmt.Errorf("清空优惠券记录失败: %v", err)
		}
		fmt.Println("✓ 优惠券记录清空完成")

		// 5. 清空储物柜订单导出记录
		fmt.Println("正在清空储物柜订单导出记录...")
		_, err = tx.Exec("DELETE FROM hg_order_export WHERE scene = 5")
		if err != nil {
			return fmt.Errorf("清空储物柜订单导出记录失败: %v", err)
		}
		fmt.Println("✓ 储物柜订单导出记录清空完成")

		// 6. 清空储物柜订单日志表
		fmt.Println("正在清空储物柜订单日志表 (hg_cabinet_order_log)...")
		_, err = tx.Exec("TRUNCATE TABLE hg_cabinet_order_log")
		if err != nil {
			return fmt.Errorf("清空储物柜订单日志表失败: %v", err)
		}
		fmt.Println("✓ 储物柜订单日志表清空完成")

		// 7. 清空储物柜订单表
		fmt.Println("正在清空储物柜订单表 (hg_cabinet_order)...")
		_, err = tx.Exec("TRUNCATE TABLE hg_cabinet_order")
		if err != nil {
			return fmt.Errorf("清空储物柜订单表失败: %v", err)
		}
		fmt.Println("✓ 储物柜订单表清空完成")

		return nil
	})

	if err != nil {
		fmt.Printf("❌ 数据清理失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🎉 储物柜订单数据清理完成！")

	// 显示清理结果
	showClearResult(ctx)
}

// showClearResult 显示清理结果
func showClearResult(ctx context.Context) {
	fmt.Println("\n=== 清理结果统计 ===")

	// 检查订单表记录数
	orderCount, err := g.DB().Model("cabinet_order").Count(ctx)
	if err != nil {
		fmt.Printf("查询订单表记录数失败: %v\n", err)
	} else {
		fmt.Printf("储物柜订单表当前记录数: %d\n", orderCount)
	}

	// 检查订单日志表记录数
	logCount, err := g.DB().Model("cabinet_order_log").Count(ctx)
	if err != nil {
		fmt.Printf("查询订单日志表记录数失败: %v\n", err)
	} else {
		fmt.Printf("储物柜订单日志表当前记录数: %d\n", logCount)
	}

	// 检查支付交易记录数
	transactionCount, err := g.DB().Model("pms_transaction").Where("scene = ?", "CABINET").Count(ctx)
	if err != nil {
		fmt.Printf("查询支付交易记录数失败: %v\n", err)
	} else {
		fmt.Printf("储物柜相关支付交易记录数: %d\n", transactionCount)
	}

	// 检查余额变动记录数
	balanceChangeCount, err := g.DB().Model("pms_balance_change").Where("scene = ?", "CABINET").Count(ctx)
	if err != nil {
		fmt.Printf("查询余额变动记录数失败: %v\n", err)
	} else {
		fmt.Printf("储物柜相关余额变动记录数: %d\n", balanceChangeCount)
	}

	// 检查优惠券记录数
	couponCount, err := g.DB().Model("pms_coupon").Where("scene = ?", 5).Count(ctx)
	if err != nil {
		fmt.Printf("查询优惠券记录数失败: %v\n", err)
	} else {
		fmt.Printf("储物柜相关优惠券记录数: %d\n", couponCount)
	}
}
