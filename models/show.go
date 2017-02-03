package models

import (
	"github.com/astaxie/beego/orm"
)

type TShow struct {
	Id       int64
	Body     int
	Hair     int
	Emotion  int
	Clothes  int
	Trousers int
	Shoes    int
}

type Show struct {
	Id       int64
	User     *User `orm:"null;rel(one);on_delete(set_null)"`
	Body     int
	Hair     int
	Emotion  int
	Clothes  int
	Trousers int
	Shoes    int
}

func (this *Show) ChangeBody(body int) {
	this.Body = body
}

func (this *Show) ChangeHair(hair int) {
	this.Hair = hair
}

func (this *Show) ChangeEmotion(emotion int) {
	this.Emotion = emotion
}

func (this *Show) ChangeClothes(clothes int) {
	this.Clothes = clothes
}

func (this *Show) ChangeTrousers(trousers int) {
	this.Trousers = trousers
}

func (this *Show) ChangeShoes(shoes int) {
	this.Shoes = shoes
}

func (this *Show) ToT() (t TShow) {
	t.Body = this.Body
	t.Hair = this.Hair
	t.Emotion = this.Emotion
	t.Clothes = this.Clothes
	t.Trousers = this.Trousers
	t.Shoes = this.Shoes
	return
}

func SaveShow(show Show) *Show {
	var err error
	show.Id, err = orm.NewOrm().Insert(&show)
	if err != nil {
		panic(err.Error())
	}
	return &show
}
