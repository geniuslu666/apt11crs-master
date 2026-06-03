package app

import (
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/nxcloud"
	"APT/internal/library/telecom"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/service"
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/member"
)

func (c *ControllerMember) MemberIntentionEdit(ctx context.Context, req *member.MemberIntentionEditReq) (res *member.MemberIntentionEditRes, err error) {
	var (
		tx gdb.TX
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var (
		MemberUser   = contexts.GetMemberUser(ctx)
		orm          = g.Model(dao.PmsMemberIntention.Table()).Ctx(ctx).Safe()
		updateResult sql.Result
		updateRow    int64
	)

	updateData := entity.PmsMemberIntention{
		MemberId:  MemberUser.Id,
		Name:      req.Name,
		Sex:       req.Sex,
		Country:   req.Country,
		Phone:     req.Phone,
		PhoneArea: req.PhoneArea,
		Language:  req.Language,
		Mail:      req.Mail,
		Remark:    req.Remark,
		WechatNo:  req.WechatNo,
		LineId:    req.LineId,
	}
	// 新增
	if updateResult, err = orm.
		Data(updateData).OmitEmptyData().Insert(); err != nil {
		return
	}
	if updateRow, err = updateResult.RowsAffected(); err != nil {
		return
	}
	if updateRow == 0 {
		// 新增会员意向表失败，请稍后重试
		err = gerror.New(gi18n.T(ctx, "failed_to_add_new_member_intention_form"))
		return
	}
	if updateRow != 1 {
		// 新增会员意向表失败，请稍后重试
		err = gerror.New(gi18n.T(ctx, "failed_to_add_new_member_intention_form"))
		return
	}

	// 是否发放短信
	var (
		MemberIntentionConfig *model.MemberIntentionConfig
	)
	if MemberIntentionConfig, err = service.BasicsConfig().GetMemberIntentionConfig(ctx); err != nil {
		return
	}

	if MemberIntentionConfig.IsSendSms == 1 {
		sendSmsPhoneArr := strings.Split(MemberIntentionConfig.SendSmsPhone, ",")
		// 开始新人礼奖励
		for _, sendPhone := range sendSmsPhoneArr {
			if !g.IsEmpty(sendPhone) {
				sendArr := strings.Split(sendPhone, "-")
				if len(sendArr) > 1 && !g.IsEmpty(sendArr[0]) {
					areaNo := sendArr[0]
					phone := sendArr[1]
					sexStr := "未知"
					if req.Sex == 1 {
						sexStr = "男"
					} else if req.Sex == 2 {
						sexStr = "女"
					}

					sendContent := "新咨询：" + req.Name + "/" + sexStr + "/" + req.PhoneArea + "-" + req.Phone + "/" + req.Mail +
						"，来自" + req.Country + "/" + req.Language + "/微信：" + req.WechatNo + "/LINE：" + req.LineId +
						"，留言：" + req.Remark

					if areaNo == "86" {
						_ = telecom.SendSms(ctx, phone, sendContent)

					} else {
						Phone := fmt.Sprintf("%s%s", areaNo, phone)
						_ = nxcloud.SendSms(ctx, Phone, sendContent)
					}
				}

			}
		}
	}

	return
}
func (c *ControllerMember) GetMemberIntentionConfig(ctx context.Context, _ *member.GetMemberIntentionConfigReq) (res *member.GetMemberIntentionConfigRes, err error) {
	res = new(member.GetMemberIntentionConfigRes)
	if res.MemberIntentionConfig, err = service.BasicsConfig().GetMemberIntentionConfig(ctx); err != nil {
		return
	}

	return
}
