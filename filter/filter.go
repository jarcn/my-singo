package filter

type Request interface{}
type Responese interface{}

type Filter interface {
	Process(data Request) (Responese, error)
}
