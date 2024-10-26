package wxlogin

import (
	"context"
	"errors"
	"math/rand/v2"
	"strconv"
	"wxlogin/model"
	"wxlogin/wx"

	"github.com/google/uuid"
)

func g6() string {
	i := rand.IntN(888889)
	return strconv.Itoa(i + 100000)
}

type SMSSender interface {
	SendSMS(string, string) error
}

type Cache interface {
	Store(context.Context, model.AuthInfo) error
	Load(context.Context, string) (model.AuthInfo, error)
}

type User struct {
	Name string
	Code string
}

type UserStore interface {
	LoadByWxOpenID(openID string) (User, bool, error)
	LoadByPhone(context.Context, string) (User, bool, error)
	BindUser(context.Context, string, User) error
}

type WxServer struct {
	wxSnsApi  *wx.WeChatClient
	userStore UserStore
	secret    []string
	cache     Cache
	smsSender SMSSender
}

type AuthData struct {
	AuthCode string
	Phone    string
	Code     string
	UID      string
	Status   string
}

func (x *WxServer) Login(ctx context.Context, jsCode string) (AuthData, error) {
	res, err := x.wxSnsApi.WxCode2Session(jsCode)
	if err != nil {
		return AuthData{}, err
	}
	u, ok, err := x.userStore.LoadByWxOpenID(res.OpenID)
	if err != nil {
		return AuthData{}, err
	}

	if ok {
		return AuthData{Code: u.Code, Status: "S100"}, nil
	}

	uid := uuid.NewString()

	authData := model.AuthInfo{
		JsCode:   jsCode,
		WxOpenID: res.OpenID,
		Code:     u.Code,
		UID:      uid,
	}

	x.cache.Store(ctx, authData)

	return AuthData{Code: u.Code, UID: uid}, nil
}

func (x *WxServer) SendSmsX(ctx context.Context, uid, phone string) error {
	user, ok, err := x.userStore.LoadByPhone(ctx, phone)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("未找到用户")
	}

	authCode := g6()

	u, _ := x.cache.Load(ctx, uid)
	u.AuthCode = authCode
	u.PhoneNo = phone
	u.Code = user.Code

	x.cache.Store(ctx, u)

	x.smsSender.SendSMS(phone, authCode)

	return nil
}

func (x *WxServer) BindX(ctx context.Context, uid, authCode string) error {

	u, _ := x.cache.Load(ctx, uid)
	if u.AuthCode == authCode {
		x.userStore.BindUser(ctx, u.WxOpenID, User{Code: u.Code})
	}

	return nil
}
