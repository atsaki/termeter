package termeter

import (
	"gopkg.in/gizak/termui.v1"
	"github.com/nsf/termbox-go"
)

type Box interface {
	GetWidth() int
	GetHeight() int
	SetWidth(w int)
	SetHeight(h int)
	Bufferers() []termui.Bufferer
	Render(x, y, w, h int)
}

type LayoutBox interface {
	Box
	AddBoxes(boxes ...Box)
}

type BoxBase struct {
	width  int
	height int
}

func (bb *BoxBase) GetWidth() int {
	return bb.width
}

func (bb *BoxBase) GetHeight() int {
	return bb.height
}

func (bb *BoxBase) SetWidth(w int) {
	bb.width = w
}

func (bb *BoxBase) SetHeight(h int) {
	bb.height = h
}

type LayoutBoxBase struct {
	BoxBase
	boxes []Box
}

func (lb *LayoutBoxBase) AddBoxes(boxes ...Box) {
	lb.boxes = append(lb.boxes, boxes...)
}

func (lb *LayoutBoxBase) Bufferers() []termui.Bufferer {
	bufs := []termui.Bufferer{}
	for _, box := range lb.boxes {
		for _, buf := range box.Bufferers() {
			bufs = append(bufs, buf)
		}
	}
	return bufs
}

type WidgetBox struct {
	BoxBase
	Widget Widget
}

func NewWidgetBox() *WidgetBox {
	return &WidgetBox{}
}

func (wb *WidgetBox) SetWidget(w Widget) {
	wb.Widget = w
}

func (wb *WidgetBox) Bufferers() []termui.Bufferer {
	return []termui.Bufferer{wb.Widget}
}

func (wb *WidgetBox) Render(x, y, w, h int) {
	if wb.Widget == nil {
		return
	}
	wb.Widget.SetX(x)
	wb.Widget.SetY(y)
	wb.Widget.SetWidth(w)
	wb.Widget.SetHeight(h)
	for _, v := range wb.Widget.Buffer() {
		termbox.SetCell(v.X, v.Y, v.Ch,
			termbox.Attribute(v.Fg),
			termbox.Attribute(v.Bg))
	}
}

type VBox struct {
	LayoutBoxBase
}

func NewVBox() *VBox {
	vb := new(VBox)
	vb.boxes = []Box{}
	return vb
}

func (vb *VBox) Render(x, y, w, h int) {
	fixedHeight := 0
	numFixedBox := 0
	for _, box := range vb.boxes {
		if box.GetHeight() > 0 {
			fixedHeight += box.GetHeight()
			numFixedBox += 1
		}
	}
	boxDefaultHeight := 0
	if len(vb.boxes) > numFixedBox && h > fixedHeight {
		boxDefaultHeight = (h - fixedHeight) / (len(vb.boxes) - numFixedBox)
	}

	for _, box := range vb.boxes {
		boxHeight := box.GetHeight()
		if boxHeight <= 0 {
			boxHeight = boxDefaultHeight
		}
		box.Render(x, y, w, boxHeight)
		y += boxHeight
	}
}

type HBox struct {
	LayoutBoxBase
}

func NewHBox() *HBox {
	hb := new(HBox)
	hb.boxes = []Box{}
	return hb
}

func (hb *HBox) Render(x, y, w, h int) {
	fixedWidth := 0
	numFixedBox := 0
	for _, box := range hb.boxes {
		if box.GetWidth() > 0 {
			fixedWidth += box.GetWidth()
			numFixedBox += 1
		}
	}

	boxDefaultWidth := 0
	if len(hb.boxes) > numFixedBox && w > fixedWidth {
		boxDefaultWidth = (w - fixedWidth) / (len(hb.boxes) - numFixedBox)
	}

	for _, box := range hb.boxes {
		boxWidth := box.GetWidth()
		if boxWidth <= 0 {
			boxWidth = boxDefaultWidth
		}
		box.Render(x, y, boxWidth, h)
		x += boxWidth
	}
}
