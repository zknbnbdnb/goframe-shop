package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Register 注册
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	// 处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	// 插入数据和返回ID
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(ctx.Value(consts.CtxUserId)).Scan(&userInfo) // 在db中取出用户信息
	if err != nil {
		return out, err
	}
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer { // 检查密保答案
		return out, gerror.New(consts.ErrSecretAnswerMsg)
	}
	UserSalt := grand.S(10)
	in.UserSalt = UserSalt
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	_, err = dao.UserInfo.Ctx(ctx).WherePri(ctx.Value(consts.CtxUserId)).Update(in) // 更新密码
	if err != nil {
		return out, err
	}
	return model.UpdatePasswordOutput{Id: userId}, err
}
