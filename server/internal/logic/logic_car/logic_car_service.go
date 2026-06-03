package logic_car

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"strings"
)

type sCarService struct{}

func NewCarService() *sCarService {
	return &sCarService{}
}

func init() {
	service.RegisterCarService(NewCarService())
}

// Model 服务ORM模型
func (s *sCarService) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.CarService.Ctx(ctx), option...)
}

// List 获取服务列表
func (s *sCarService) List(ctx context.Context, in *input_car.CarServiceListInp) (list []*input_car.CarServiceListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	// 字段过滤
	mod = mod.Fields(input_car.CarServiceListModel{})

	if !g.IsEmpty(in.ServiceType) {
		mod = mod.Where(dao.CarService.Columns().ServiceType, in.ServiceType)
	}

	// 分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	if !g.IsEmpty(in.Sort) {
		if in.Sort == "ascend" {
			mod = mod.Order(dao.CarService.Columns().Sort, "asc")
		} else {
			mod = mod.Order(dao.CarService.Columns().Sort, "desc")
		}
	} else {
		mod = mod.OrderDesc(dao.CarService.Columns().Sort).OrderDesc(dao.CarService.Columns().Id)
	}

	// 排序
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	// 查询数据
	if in.Pagination {
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务列表失败，请稍后重试！")
			return
		}
	} else {
		if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取服务列表失败，请稍后重试！")
			return
		}
	}
	return
}

// Edit 修改/新增服务
func (s *sCarService) Edit(ctx context.Context, in *input_car.CarServiceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var (
			StartServiceAddressInfo []*entity.CarServiceAddress
			EndServiceAddressInfo   []*entity.CarServiceAddress
		)

		// 修改
		if in.Id > 0 {

			if _, err = s.Model(ctx).
				Fields(input_car.CarServiceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改服务失败，请稍后重试！")
			}

			if !g.IsEmpty(in.StartIds) {
				for _, Start := range strings.Split(in.StartIds, ",") {
					if !g.IsEmpty(Start) {
						StartId, _ := strconv.ParseInt(Start, 10, 64)
						StartServiceAddressInfo = append(StartServiceAddressInfo, &entity.CarServiceAddress{
							ServiceId: int64(in.Id),
							AddressId: StartId,
							Type:      1,
						})
					}
				}
			}

			if !g.IsEmpty(StartServiceAddressInfo) {
				if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, in.Id).Where(dao.CarServiceAddress.Columns().Type, 1).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.CarServiceAddress.Ctx(ctx).OmitEmptyData().Insert(StartServiceAddressInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, in.Id).Where(dao.CarServiceAddress.Columns().Type, 1).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
					return
				}
			}

			if !g.IsEmpty(in.EndIds) {
				for _, End := range strings.Split(in.EndIds, ",") {
					if !g.IsEmpty(End) {
						EndId, _ := strconv.ParseInt(End, 10, 64)
						EndServiceAddressInfo = append(EndServiceAddressInfo, &entity.CarServiceAddress{
							ServiceId: int64(in.Id),
							AddressId: EndId,
							Type:      2,
						})
					}
				}
			}

			if !g.IsEmpty(EndServiceAddressInfo) {
				if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, in.Id).Where(dao.CarServiceAddress.Columns().Type, 2).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
					return
				}
				if _, err = dao.CarServiceAddress.Ctx(ctx).OmitEmptyData().Insert(EndServiceAddressInfo); err != nil {
					return err
				}
			} else {
				if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, in.Id).Where(dao.CarServiceAddress.Columns().Type, 2).Delete(); err != nil {
					err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
					return
				}
			}
			return
		}

		var (
			lastInsertId int64
		)
		// 新增
		if lastInsertId, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(input_car.CarServiceInsertFields{}).
			Data(in).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "新增服务失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "新增失败，请稍后重试！")
			return
		}

		if !g.IsEmpty(in.StartIds) {
			for _, Start := range strings.Split(in.StartIds, ",") {
				if !g.IsEmpty(Start) {
					StartId, _ := strconv.ParseInt(Start, 10, 64)
					StartServiceAddressInfo = append(StartServiceAddressInfo, &entity.CarServiceAddress{
						ServiceId: lastInsertId,
						AddressId: StartId,
						Type:      1,
					})
				}
			}
		}

		if !g.IsEmpty(StartServiceAddressInfo) {
			if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, lastInsertId).Where(dao.CarServiceAddress.Columns().Type, 1).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.CarServiceAddress.Ctx(ctx).OmitEmptyData().Insert(StartServiceAddressInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, lastInsertId).Where(dao.CarServiceAddress.Columns().Type, 1).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
				return
			}
		}

		if !g.IsEmpty(in.EndIds) {
			for _, End := range strings.Split(in.EndIds, ",") {
				if !g.IsEmpty(End) {
					EndId, _ := strconv.ParseInt(End, 10, 64)
					EndServiceAddressInfo = append(EndServiceAddressInfo, &entity.CarServiceAddress{
						ServiceId: lastInsertId,
						AddressId: EndId,
						Type:      2,
					})
				}
			}
		}

		if !g.IsEmpty(EndServiceAddressInfo) {
			if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, lastInsertId).Where(dao.CarServiceAddress.Columns().Type, 2).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.CarServiceAddress.Ctx(ctx).OmitEmptyData().Insert(EndServiceAddressInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.CarServiceAddress.Ctx(ctx).Where(dao.CarServiceAddress.Columns().ServiceId, lastInsertId).Where(dao.CarServiceAddress.Columns().Type, 2).Delete(); err != nil {
				err = gerror.Wrap(err, "清理服务地址旧数据失败，请稍后重试！")
				return
			}
		}
		return
	})
}

// Delete 删除服务
func (s *sCarService) Delete(ctx context.Context, in *input_car.CarServiceDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除服务失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取服务最大排序
func (s *sCarService) MaxSort(ctx context.Context, in *input_car.CarServiceMaxSortInp) (res *input_car.CarServiceMaxSortModel, err error) {
	if err = dao.CarService.Ctx(ctx).Fields(dao.CarService.Columns().Sort).OrderDesc(dao.CarService.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取服务最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(input_car.CarServiceMaxSortModel)
	}

	res.Sort = input_form.DefaultMaxSort(res.Sort)
	return
}

// View 获取服务指定信息
func (s *sCarService) View(ctx context.Context, in *input_car.CarServiceViewInp) (res *input_car.CarServiceViewModel, err error) {
	if !in.IsLanguage {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取服务信息，请稍后重试！")
			return
		}
	} else {
		if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
			err = gerror.Wrap(err, "获取服务信息，请稍后重试！")
			return
		}
	}

	return
}

// Status 更新服务状态
func (s *sCarService) Status(ctx context.Context, in *input_car.CarServiceStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.CarService.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新服务状态失败，请稍后重试！")
		return
	}
	return
}

// ServiceSort 编辑服务排序
func (s *sCarService) ServiceSort(ctx context.Context, in *input_car.CarServiceSortInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Update(g.Map{
		dao.CarService.Columns().Sort: in.Sort,
	}); err != nil {
		err = gerror.Wrap(err, "修改服务排序，请稍后重试！")
		return
	}
	return
}
