package logic_app_member

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/utility/encrypt"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetMemberQrCode 获取会员二维码
func (s *sAppMember) GetMemberQrCode(ctx context.Context) (code string, err error) {
	var memberInfo *model.MemberIdentity
	// 获取授权用户信息
	memberInfo = contexts.GetMemberUser(ctx)

	// 获取会员详细信息
	var member *entity.PmsMember
	if err = dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, memberInfo.Id).Scan(&member); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "failed_to_get_member_info"))
		return
	}

	if g.IsEmpty(member) {
		err = gerror.New(gi18n.T(ctx, "failed_to_get_member_info"))
		return
	}

	// 查询会员是否绑定员工
	var employee *entity.Employee
	if err = dao.Employee.Ctx(ctx).Where(dao.Employee.Columns().MemberId, memberInfo.Id).Scan(&employee); err != nil {
		g.Log().Warningf(ctx, "query employee failed: %v", err)
		// 查询员工失败不影响二维码生成，继续执行
		err = nil
	}

	// 生成动态二维码
	timestamp := gtime.Now().Unix()
	plainText := fmt.Sprintf("%s|%d", member.MemberNo, timestamp)
	codeAes := encrypt.MustAesECBEncryptToString(plainText, string(consts.RequestEncryptKey))
	code = fmt.Sprintf("%s|MEMBER", codeAes)

	return
}
