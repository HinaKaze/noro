package filter

import (
	"github.com/astaxie/beego/context"
	muser "github.com/hinakaze/noro/models/user"
)

func FilterLogin(ctx *context.Context) {
	if user, ok := IsLogin(ctx); ok {
		ctx.Output.Session("user", user)
	} else {
		ctx.Redirect(302, "/login")
	}
}

func IsLogin(ctx *context.Context) (*muser.User, bool) {
	if username, ok := ctx.GetSecureCookie("noro_", "_n"); ok {
		if user := muser.GetUserByName(username); user != nil {
			loginToken, ok := ctx.GetSecureCookie("noro_", "_t")
			if !ok {
				return nil, false
			}
			loginSeq, ok := ctx.GetSecureCookie("noro_", "_s")
			if !ok {
				return nil, false
			}
			if user.CheckLoginSeq(loginSeq) && user.CheckLoginToken(loginToken) {
				return user, true
			} else {
				return nil, false
			}
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}
