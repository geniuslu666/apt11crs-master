package notify

import (
	"APT/internal/dao"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"strings"

	"APT/api/notify/kefu"
)

func (c *ControllerKefu) FindMemberInfo(ctx context.Context, req *kefu.FindMemberInfoReq) (res *kefu.FindMemberInfoRes, err error) {
	res = new(kefu.FindMemberInfoRes)
	if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, req.Id).WithAll().Scan(&res); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	return
}
func (c *ControllerKefu) FindMemberInfos(ctx context.Context, req *kefu.FindMemberInfosReq) (res *kefu.FindMemberInfosRes, err error) {
	res = new(kefu.FindMemberInfosRes)
	mod := dao.PmsMember.Ctx(ctx)

	if !g.IsEmpty(req.Ids) {
		mod = mod.WherePri(strings.Split(req.Ids, ","))
	}
	if !g.IsEmpty(req.KeyWords) {
		mod = mod.Where("( member_no = ? OR id = ? OR full_name LIKE ? OR phone LIKE ?)", req.KeyWords, req.KeyWords, "%"+req.KeyWords+"%", "%"+req.KeyWords+"%")
	}

	if err = mod.WithAll().Scan(&res.List); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	return
}
