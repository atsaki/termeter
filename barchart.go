package termeter

import "gopkg.in/gizak/termui.v1"

type BarChartWidget struct {
	BarChart *termui.BarChart
}

func NewBarChartWidget() *BarChartWidget {
	return &BarChartWidget{
		BarChart: termui.NewBarChart(),
	}
}

func (bc *BarChartWidget) SetX(x int) {
	bc.BarChart.X = x
}

func (bc *BarChartWidget) SetY(y int) {
	bc.BarChart.Y = y
}

func (bc *BarChartWidget) SetWidth(w int) {
	bc.BarChart.Width = w
}

func (bc *BarChartWidget) SetHeight(h int) {
	bc.BarChart.Height = h
}

func (bc *BarChartWidget) Buffer() []termui.Point {
	defer func() {
		recover()
	}()
	return bc.BarChart.Buffer()
}

func (bc *BarChartWidget) Add(label string, freq int) {
	bc.BarChart.Data = append(bc.BarChart.Data, freq)
	bc.BarChart.DataLabels = append(bc.BarChart.DataLabels, label)
}

func (bc *BarChartWidget) Update(labels []string, freqs []int) {
	bc.BarChart.Data = make([]int, len(freqs), len(freqs))
	bc.BarChart.DataLabels = make([]string, len(labels), len(labels))
	copy(bc.BarChart.Data, freqs)
	copy(bc.BarChart.DataLabels, labels)
}
