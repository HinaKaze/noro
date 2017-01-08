package models

import (
	"fmt"
)

type Friendship struct {
	Id           int
	UserId1      int
	UserId2      int
	Relationship int
}

func (f *Friendship) AddRelationship(increment int) {
	f.Relationship += increment
}

func (f *Friendship) ToT() (t TFriendship) {
	if u, ok := GetUser(f.UserId2); ok {
		t.Friend = u.ToT(false)
	} else {
		panic(fmt.Sprintf("Target friend id [%d] invalid", f.UserId2))
	}

	t.Relationship = f.Relationship
	return
}

type TFriendship struct {
	Friend       TUser
	Relationship int
}
