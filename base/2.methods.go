package base

type react struct {
	Length int `json:"length"`
	Height int `json:"height"`
}

var React *react

func MewReact(length int, height int) *react {
	React = &react{
		Length: length,
		Height: height,
	}

	return React
}

func (r *react) GetArea() int {
	return r.Length * r.Height
}
