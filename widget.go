package termeter

import "github.com/gizak/termui"

type Widget interface {
	SetX(x int)
	SetY(y int)
	SetWidth(w int)
	SetHeight(h int)
	Buffer() []termui.Point
}
