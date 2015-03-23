package termeter

import (
	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

func Render(box Box) {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.Attribute(termui.ColorDefault))
	box.Render(0, 0, w, h)
	termbox.Flush()
}
