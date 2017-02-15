package user

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"

	"github.com/hinakaze/noro/common"
	muser "github.com/hinakaze/noro/models/user"
)

type WardrobeController struct {
	beego.Controller
}

func (this *WardrobeController) Get() {
	this.TplName = "user/wardrobe.html"
}

func (this *WardrobeController) Put() {
	defer this.ServeJSON()
	result := common.Result{}
	defer func() {
		this.Data["json"] = string(result.ToJSON())
	}()

	body := this.Ctx.Input.RequestBody
	log.Println(len(body), string(body))
	newShow := muser.Show{}
	err := json.Unmarshal(body, &newShow)
	if err != nil {
		result.Type = common.ResultFailed
		result.Info = err.Error()
		return
	}
	user, ok := this.GetSession("user").(*muser.User)
	if !ok {
		this.Redirect("/login", 302)
		return
	}
	if user.Show.Equal(newShow) {
		result.Type = common.ResultSuccess
		return
	} else {
		newShow.User = user
		muser.UpdateShow(&newShow)
		result.Type = common.ResultSuccess
		return
	}
}
