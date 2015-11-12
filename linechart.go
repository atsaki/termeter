package termeter

import "gopkg.in/gizak/termui.v1"

type LineChartWidget struct {
	Data       *float64RingBuffer
	DataLabels *stringRingBuffer
	Mode       string
	LineChart  *termui.LineChart
}

func NewLineChartWidget(bufsize int) *LineChartWidget {
	lc := &LineChartWidget{
		Data:       newFloat64RingBuffer(bufsize),
		DataLabels: newStringRingBuffer(bufsize),
		LineChart:  termui.NewLineChart(),
	}
	return lc
}

func (lc *LineChartWidget) SetX(x int) {
	lc.LineChart.X = x
}

func (lc *LineChartWidget) SetY(y int) {
	lc.LineChart.Y = y
}

func (lc *LineChartWidget) SetWidth(w int) {
	lc.LineChart.Width = w
	lc.updateLineChartData()
}

func (lc *LineChartWidget) SetHeight(h int) {
	lc.LineChart.Height = h
}

func (lc *LineChartWidget) Buffer() []termui.Point {
	defer func() {
		recover()
	}()
	return lc.LineChart.Buffer()
}

func (lc *LineChartWidget) Add(x float64, dataLabel string) {
	lc.Data.Add(x)
	if dataLabel != "" {
		lc.DataLabels.Add(dataLabel)
	}
	lc.updateLineChartData()
}

func (lc *LineChartWidget) Update(xs []float64, dataLabels []string) {
	lc.Data = newFloat64RingBuffer(lc.Data.Capacity())
	if xs != nil {
		for _, x := range xs {
			lc.Data.Add(x)
		}
	}
	lc.DataLabels = newStringRingBuffer(lc.DataLabels.Capacity())
	if dataLabels != nil {
		for _, dataLabel := range dataLabels {
			lc.DataLabels.Add(dataLabel)
		}
	}
	lc.updateLineChartData()
}

func (lc *LineChartWidget) Clear() {
	lc.Update(nil, nil)
}

func (lc *LineChartWidget) updateLineChartData() {
	lc.LineChart.Mode = lc.Mode
	if lc.Mode == "dot" {
		lc.LineChart.Data = lc.Data.Last(lc.LineChart.Width)
		lc.LineChart.DataLabels = lc.DataLabels.Last(lc.LineChart.Width)
	} else {
		lc.LineChart.Data = lc.Data.Last(2 * lc.LineChart.Width)
		lc.LineChart.DataLabels = lc.DataLabels.Last(2 * lc.LineChart.Width)
	}
}
