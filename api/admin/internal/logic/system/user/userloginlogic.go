package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"zero-demo/api/admin/internal/svc"
	"zero-demo/api/admin/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("用户名或密码不能为空")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errors.New("用户不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errors.New("密码错误")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	refreshSecret := l.svcCtx.Config.Auth.RefreshSecret

	jwtToken, _ := l.getJwtToken(accessSecret, userInfo.Username, now, accessExpire, userInfo.Id)
	refresh, err := l.getRefreshToken(refreshSecret)
	if err != nil {
		logc.Debug(l.ctx, "getRefreshToken", err)
	}

	return &types.LoginResp{
		UserName:     userInfo.Username,
		AccessToken:  jwtToken,
		RefreshToken: refresh,
		ExpireTime:   now + accessExpire,
	}, nil
}

type UserClaims struct {
	ID      int64  `json:"id"`
	Account string `json:"account"`
	Time    int64  `json:"time"`
	jwt.RegisteredClaims
}

func (l *UserLoginLogic) getJwtToken(secretKey, account string, iat, seconds, uid int64) (string, error) {
	claims := UserClaims{
		uid,
		account,
		iat,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(seconds))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (l *UserLoginLogic) getRefreshToken(secretKey string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
