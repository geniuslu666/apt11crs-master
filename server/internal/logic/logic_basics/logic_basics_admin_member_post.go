package logic_basics

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sBasicsAdminMemberPost struct{}

func NewBasicsAdminMemberPost() *sBasicsAdminMemberPost {
	return &sBasicsAdminMemberPost{}
}

func init() {
	service.RegisterBasicsAdminMemberPost(NewBasicsAdminMemberPost())
}

// UpdatePostIds 更新用户岗位
func (s *sBasicsAdminMemberPost) UpdatePostIds(ctx context.Context, memberId int64, postIds []int64) (err error) {
	if _, err = dao.AdminMemberPost.Ctx(ctx).Where(dao.AdminMemberPost.Columns().MemberId, memberId).Delete(); err != nil {
		err = gerror.Wrap(err, "清理用户旧岗位数据失败，请稍后重试！")
		return
	}

	for i := 0; i < len(postIds); i++ {
		_, err = dao.AdminMemberPost.Ctx(ctx).OmitEmptyData().Insert(entity.AdminMemberPost{
			MemberId: memberId,
			PostId:   postIds[i],
		})
		if err != nil {
			err = gerror.Wrap(err, "加入用户岗位数据失败，请稍后重试！")
			return err
		}
	}
	return
}
