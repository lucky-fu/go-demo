package base

type geometry interface {
	GetArea1() int
	GetPerim() int
}

type react1 struct {
	Length int `json:"length"`
	Height int `json:"height"`
}

var React1 *react1

func NewReact(length int, height int) *react {
	React := &react{
		Length: length,
		Height: height,
	}

	return React
}

func (r *react) GetArea1() int {
	return r.Height * r.Length
}

func (r *react) GetPerim() int {
	return 2 * (r.Length + r.Height)
}
