package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sCarAddressType struct{}

func NewCarAddressType() *sCarAddressType {
	return &sCarAddressType{}
}

func init() {
	service.RegisterCarAddressType(NewCarAddressType())
}

// Model 接送机地址类型ORM模型
func (s *sCarAddressType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarAddressType.Ctx(ctx), option...)
}

// List 获取接送机地址类型列表
func (s *sCarAddressType) List(ctx context.Context, in *input_car.CarAddressTypeListInp) (list []*input_car.CarAddressTypeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(input_car.CarAddressTypeListModel{})

	// 合作类型
	if !g.IsEmpty(in.TypeName) {
		mod = mod.WhereLike(dao.CarAddressType.Columns().TypeName, "%"+in.TypeName+"%")
	}

	// 查询状态1、启用 2、禁用
	if in.Status > 0 {
		mod = mod.Where(dao.CarAddressType.Columns().Status, in.Status)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.CarAddressType.Columns().Sort)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取接送机地址类型列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取接送机地址类型列表失败，请稍后重试！")
			return
		}
	}

	return
}
