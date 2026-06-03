package logic_basics

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/location"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/convert"
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"sync"
)

type sBasicsBlacklist struct {
	sync.RWMutex
	list map[string]struct{}
}

func NewBasicsBlacklist() *sBasicsBlacklist {
	return &sBasicsBlacklist{
		list: make(map[string]struct{}),
	}
}

func init() {
	service.RegisterBasicsBlacklist(NewBasicsBlacklist())
}

func (s *sBasicsBlacklist) Delete(ctx context.Context, in *input_basics.BlacklistDeleteInp) (err error) {
	defer s.VariableLoad(ctx, err)
	_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Delete()
	return
}

func (s *sBasicsBlacklist) Edit(ctx context.Context, in *input_basics.BlacklistEditInp) (err error) {
	defer s.VariableLoad(ctx, err)
	if in.Ip == "" {
		err = gerror.New("ip不能为空")
		return
	}

	if in.Id > 0 {
		_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Data(in).Update()
		return
	}

	_, err = dao.SysBlacklist.Ctx(ctx).Data(in).OmitEmptyData().Insert()
	return
}

func (s *sBasicsBlacklist) Status(ctx context.Context, in *input_basics.BlacklistStatusInp) (err error) {
	defer s.VariableLoad(ctx, err)

	_, err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Data("status", in.Status).Update()
	return
}

func (s *sBasicsBlacklist) View(ctx context.Context, in *input_basics.BlacklistViewInp) (res *input_basics.BlacklistViewModel, err error) {
	err = dao.SysBlacklist.Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

func (s *sBasicsBlacklist) List(ctx context.Context, in *input_basics.BlacklistListInp) (list []*input_basics.BlacklistListModel, totalCount int, err error) {
	mod := dao.SysBlacklist.Ctx(ctx)
	cols := dao.SysBlacklist.Columns()

	if in.Ip != "" {
		mod = mod.WhereLike(cols.Ip, "%"+in.Ip+"%")
	}

	if in.Remark != "" {
		mod = mod.WhereLike(cols.Remark, "%"+in.Remark+"%")
	}

	if in.Status > 0 {
		mod = mod.Where(cols.Status, in.Status)
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(cols.CreatedAt, gtime.New(in.CreatedAt[0]), gtime.New(in.CreatedAt[1]))
	}

	totalCount, err = mod.Count()
	if err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Page(in.Page, in.PerPage).Order("id desc").Scan(&list); err != nil {
		return
	}
	return
}

func (s *sBasicsBlacklist) VariableLoad(ctx context.Context, err error) {
	if err == nil {
		s.Load(ctx)
	}
}

func (s *sBasicsBlacklist) Load(ctx context.Context) {
	s.RLock()
	defer s.RUnlock()

	s.list = make(map[string]struct{})

	array, err := dao.SysBlacklist.Ctx(ctx).
		Fields(dao.SysBlacklist.Columns().Ip).
		Where(dao.SysBlacklist.Columns().Status, consts.StatusEnabled).
		Array()
	if err != nil {
		g.Log().Errorf(ctx, "load blacklist fail：%+v", err)
		return
	}

	for _, v := range array {
		list := convert.IpFilterStrategy(v.String())
		if len(list) > 0 {
			for k := range list {
				s.list[k] = struct{}{}
			}
		}
	}
}

func (s *sBasicsBlacklist) VerifyRequest(r *ghttp.Request) (err error) {
	if len(s.list) == 0 {
		return
	}

	if _, ok := s.list[location.GetClientIp(r)]; ok {
		err = gerror.NewCode(gcode.New(gcode.CodeServerBusy.Code(), "请求异常，已被封禁，如有疑问请联系管理员！", nil))
		return
	}
	return
}

func (s *sBasicsBlacklist) ClusterSync(ctx context.Context, message *gredis.Message) {
	s.Load(ctx)
}
