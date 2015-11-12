package termeter

import "gopkg.in/gizak/termui.v1"

type Widget interface {
	SetX(x int)
	SetY(y int)
	SetWidth(w int)
	SetHeight(h int)
	Buffer() []termui.Point
}
