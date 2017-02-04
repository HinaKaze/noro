package user

type Friendship struct {
	Id           int
	User1        *User `orm:"rel(fk)"`
	User2        *User `orm:"rel(fk)"`
	Relationship int
}

func (f *Friendship) AddRelationship(increment int) {
	f.Relationship += increment
}

func (f *Friendship) ToT() (t TFriendship) {
	t.Friend = f.User2.ToT(false)

	t.Relationship = f.Relationship
	return
}

type TFriendship struct {
	Friend       TUser
	Relationship int
}
