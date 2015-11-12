package termeter

import "gopkg.in/gizak/termui.v1"

type ParWidget struct {
	Par *termui.Par
}

func NewParWidget() *ParWidget {
	return &ParWidget{
		Par: termui.NewPar(""),
	}
}

func (p *ParWidget) SetX(x int) {
	p.Par.X = x
}

func (p *ParWidget) SetY(y int) {
	p.Par.Y = y
}

func (p *ParWidget) SetWidth(w int) {
	p.Par.Width = w
}

func (p *ParWidget) SetHeight(h int) {
	p.Par.Height = h
}

func (p *ParWidget) Buffer() []termui.Point {
	defer func() {
		recover()
	}()
	return p.Par.Buffer()
}

func (p *ParWidget) Update(text string) {
	p.Par.Text = text
}
