package color

type Color struct {
	R, G, B uint8
}

func New(r, g, b uint8) *Color {
	return &Color{r, g, b}
}
