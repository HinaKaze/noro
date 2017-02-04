package chat

import (
	"time"

	"github.com/hinakaze/noro/common"
	"github.com/hinakaze/noro/models/user"
)

type ChatMessage struct {
	Id       int
	Type     int        //0 join,1 leave,2 message
	User     *user.User `orm:"rel(fk)"`
	Text     string
	Time     time.Time
	ChatRoom *ChatRoom `orm:"rel(fk)"`
}

type TChatMessage struct {
	Id   int
	Type int
	User user.TUser
	Text string
	Time string
}

func (c *ChatMessage) ToT() (t TChatMessage) {
	t.Id = c.Id
	t.Type = c.Type
	t.User = c.User.ToT(false)
	t.Text = c.Text
	t.Time = common.ToFormatTime(c.Time)
	return
}
