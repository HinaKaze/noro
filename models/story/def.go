package story

const (
	StoryModeSingle uint8 = iota
	StoryModeMulti
)

type Story struct {
	Id         int64
	Name       string
	Desc       string
	Mode       uint8
	FirstScene *StoryScene
	Characters []*StoryCharacter `orm:"reverse(many)"`
}

type StoryScene struct {
	Id        int64
	Story     *Story `orm:"rel(one)"`
	Name      string
	Contents  []*StoryContent `orm:"reverse(many)"`
	StoryFlag *StoryFlag      `orm:"reverse(one)"`
	End       bool
}

type StoryContent struct {
	Id    int64
	Scene *StoryScene `orm:"rel(fk)"`
	Who   string
	Text  string
	Delay int // seconds
}

type StoryFlag struct {
	Id        int64
	Scene     *StoryScene `orm:"rel(one)"`
	Content   string
	NextScene *StoryScene `orm:"rel(one)"`
}

type StoryCharacter struct {
	Id     int64
	Story  *Story `orm:"rel(fk)"`
	Name   string
	Avatar string //avatar url
	Desc   string
}
