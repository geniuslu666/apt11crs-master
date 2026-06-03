package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	ClearCabinetOrders = &gcmd.Command{
		Name:  "clear-cabinet-orders",
		Usage: "clear-cabinet-orders",
		Brief: "清理储物柜订单数据",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return clearCabinetOrdersFunc(ctx)
		},
	}
)

func clearCabinetOrdersFunc(ctx context.Context) (err error) {
	fmt.Println("开始执行储物柜订单数据清理脚本...")

	// 确认操作
	fmt.Print("警告：此操作将清空所有储物柜订单数据，包括订单表和日志表。是否继续？(y/N): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("操作已取消")
		return nil
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

		// 6. 清空储物柜订单日志
		fmt.Println("正在清空储物柜订单日志...")
		_, err = tx.Exec("TRUNCATE TABLE hg_cabinet_order_log")
		if err != nil {
			return fmt.Errorf("清空储物柜订单日志失败: %v", err)
		}
		fmt.Println("✓ 储物柜订单日志清空完成")

		// 7. 清空储物柜订单表
		fmt.Println("正在清空储物柜订单表...")
		_, err = tx.Exec("TRUNCATE TABLE hg_cabinet_order")
		if err != nil {
			return fmt.Errorf("清空储物柜订单表失败: %v", err)
		}
		fmt.Println("✓ 储物柜订单表清空完成")

		return nil
	})

	if err != nil {
		fmt.Printf("❌ 清理失败: %v\n", err)
		return err
	}

	fmt.Println("🎉 储物柜订单数据清理完成！")
	return nil
}
