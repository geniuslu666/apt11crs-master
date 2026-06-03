package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBasicsServeLog struct{}

func NewBasicsServeLog() *sBasicsServeLog {
	return &sBasicsServeLog{}
}

func init() {
	service.RegisterBasicsServeLog(NewBasicsServeLog())
}

func (s *sBasicsServeLog) Model(ctx context.Context) *gdb.Model {
	return dao.SysServeLog.Ctx(ctx)
}

func (s *sBasicsServeLog) List(ctx context.Context, in *input_basics.ServeLogListInp) (list []*input_basics.ServeLogListModel, totalCount int, err error) {
	mod := dao.SysServeLog.Ctx(ctx)

	if in.TraceId != "" {
		mod = mod.Where(dao.SysServeLog.Columns().TraceId, in.TraceId)
	}

	if in.LevelFormat != "" {
		mod = mod.WhereLike(dao.SysServeLog.Columns().LevelFormat, in.LevelFormat)
	}

	if len(in.TriggerNs) == 2 {
		mod = mod.WhereBetween(dao.SysServeLog.Columns().TriggerNs, in.TriggerNs[0], in.TriggerNs[1])
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SysServeLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.SysServeLog.Table(), dao.SysServeLog.Columns().TraceId,
		dao.SysLog.Table(), "sysLog", dao.SysLog.Columns().ReqId,
	)...)

	totalCount, err = mod.Clone().Count(1)
	if err != nil || totalCount == 0 {
		return
	}

	fields, err := hgorm.GenJoinSelect(ctx, input_basics.ServeLogListModel{}, &dao.SysServeLog, []*hgorm.Join{
		{Dao: &dao.SysLog, Alias: "sysLog"},
	})

	if err != nil {
		return
	}

	err = mod.Fields(fields).Handler(handler.FilterAuth).Page(in.Page, in.PerPage).OrderDesc(dao.SysServeLog.Columns().Id).Scan(&list)
	return
}

func (s *sBasicsServeLog) Export(ctx context.Context, in *input_basics.ServeLogListInp) (err error) {
	list, _, err := s.List(ctx, in)
	if err != nil {
		return
	}

	tags, err := convert.GetEntityDescTags(input_basics.ServeLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出服务日志-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("服务日志")
		exports   []input_basics.ServeLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return err
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sBasicsServeLog) Delete(ctx context.Context, in *input_basics.ServeLogDeleteInp) (err error) {
	_, err = dao.SysServeLog.Ctx(ctx).Where(dao.SysServeLog.Columns().Id, in.Id).Delete()
	return
}

func (s *sBasicsServeLog) View(ctx context.Context, in *input_basics.ServeLogViewInp) (res *input_basics.ServeLogViewModel, err error) {
	err = dao.SysServeLog.Ctx(ctx).Where(dao.SysServeLog.Columns().Id, in.Id).Scan(&res)
	return
}

func (s *sBasicsServeLog) RealWrite(ctx context.Context, models entity.SysServeLog) (err error) {
	_, err = dao.SysServeLog.Ctx(ctx).FieldsEx(dao.SysLog.Columns().Id).Data(models).OmitEmptyData().Insert()
	return
}
