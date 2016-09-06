package main

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) perim() float64 {
	return 2 * r.Width + 2 * r.Height
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}