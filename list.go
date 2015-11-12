package termeter

import "gopkg.in/gizak/termui.v1"

type ListWidget struct {
	List *termui.List
}

func NewListWidget() *ListWidget {
	return &ListWidget{
		List: termui.NewList(),
	}
}

func (ls *ListWidget) SetX(x int) {
	ls.List.X = x
}

func (ls *ListWidget) SetY(y int) {
	ls.List.Y = y
}

func (ls *ListWidget) SetWidth(w int) {
	ls.List.Width = w
}

func (ls *ListWidget) SetHeight(h int) {
	ls.List.Height = h
}

func (ls *ListWidget) Buffer() []termui.Point {
	defer func() {
		recover()
	}()
	return ls.List.Buffer()
}

func (ls *ListWidget) Add(items ...string) {
	for _, s := range items {
		ls.List.Items = append(ls.List.Items, s)
	}
}

func (ls *ListWidget) Update(items []string) {
	ls.List.Items = make([]string, 0, len(items))
	for _, s := range items {
		ls.List.Items = append(ls.List.Items, s)
	}
}
