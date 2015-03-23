package termeter

import (
	"fmt"
	"sort"

	"github.com/VividCortex/gohistogram"
)

// panel types
const (
	LINE = iota
	COUNTER
	CDF
)

// sort mode of counter panel
const (
	SORT_NONE = iota
	SORT_ALPHABETICAL
	SORT_NUMERICAL
)

// panel parameters
const (
	HISTOGRAM_BIN_COUNT = 50
	LIST_WIDTH          = 30
	BUFFER_SIZE         = 1000
)

type Panel interface {
	LayoutBox
	SetWidget(w Widget)
	GetType() int
}

type PanelBase struct {
	*HBox
	label     string
	list      *ListWidget
	panelType int
	widget    Widget
}

func NewPanelBase(label string) *PanelBase {
	listBox := NewWidgetBox()
	listBox.SetWidth(LIST_WIDTH)
	widgetBox := NewWidgetBox()
	hbox := NewHBox()
	hbox.AddBoxes(listBox, widgetBox)

	listWidget := NewListWidget()
	listBox.SetWidget(listWidget)
	listWidget.Add(label)

	return &PanelBase{
		HBox:  hbox,
		label: label,
		list:  listWidget,
	}
}

func (p *PanelBase) SetWidget(w Widget) {
	p.widget = w
	p.boxes[1].(*WidgetBox).SetWidget(w)
}

func (p *PanelBase) GetType() int {
	return p.panelType
}

type LineChartPanel struct {
	*PanelBase
	histogram *gohistogram.NumericHistogram
	min       float64
	max       float64
}

func NewLineChartPanel(label string) *LineChartPanel {
	panel := &LineChartPanel{
		PanelBase: NewPanelBase(label),
		histogram: gohistogram.NewHistogram(HISTOGRAM_BIN_COUNT),
	}
	panel.panelType = LINE
	panel.SetWidget(NewLineChartWidget(BUFFER_SIZE))
	return panel
}

func (p *LineChartPanel) SetMode(mode string) {
	chartWidget := p.widget.(*LineChartWidget)
	chartWidget.Mode = mode
}

func (p *LineChartPanel) Add(x float64, dataLabel string) {
	chartWidget := p.widget.(*LineChartWidget)

	chartWidget.Add(x, dataLabel)
	p.histogram.Add(x)

	if x < p.min {
		p.min = x
	}
	if p.max < x {
		p.max = x
	}

}

func (p *LineChartPanel) Update() {
	p.list.Update([]string{
		p.label,
		fmt.Sprintf("Count : %.0f", p.histogram.Count()),
		fmt.Sprintf("Mean  : %.2f", p.histogram.Mean()),
		fmt.Sprintf("Max   : %.2f", p.max),
		fmt.Sprintf("Min   : %.2f", p.min),
		fmt.Sprintf("Var   : %.2f", p.histogram.Variance()),
		fmt.Sprintf("Q1    : %.2f", p.histogram.Quantile(0.25)),
		fmt.Sprintf("Q2    : %.2f", p.histogram.Quantile(0.50)),
		fmt.Sprintf("Q3    : %.2f", p.histogram.Quantile(0.75)),
	})
}

type CDFPanel struct {
	*PanelBase
	histogram *gohistogram.NumericHistogram
	min       float64
	max       float64
}

func NewCDFPanel(label string) *CDFPanel {
	panel := &CDFPanel{
		PanelBase: NewPanelBase(label),
		histogram: gohistogram.NewHistogram(HISTOGRAM_BIN_COUNT),
	}
	panel.panelType = CDF
	panel.SetWidget(NewLineChartWidget(BUFFER_SIZE))
	return panel
}

func (p *CDFPanel) SetMode(mode string) {
	chartWidget := p.widget.(*LineChartWidget)
	chartWidget.Mode = mode
}

func (p *CDFPanel) Add(x float64) {
	p.histogram.Add(x)

	if x < p.min {
		p.min = x
	}
	if p.max < x {
		p.max = x
	}
	p.Update()
}

func (p *CDFPanel) Update() {
	chartWidget := p.widget.(*LineChartWidget)
	chartWidth := chartWidget.LineChart.Width
	npoint := 2 * chartWidth
	if chartWidget.Mode == "dot" {
		npoint = chartWidth
	}

	chartWidget.Clear()
	w := (p.max - p.min) / float64(npoint)
	for i := 0; i < npoint; i++ {
		x := p.min + float64(i)*w
		chartWidget.Add(p.histogram.CDF(x), fmt.Sprintf("%.2f", x))
	}

	p.list.Update([]string{
		p.label,
		fmt.Sprintf("Count : %.0f", p.histogram.Count()),
		fmt.Sprintf("Mean  : %.2f", p.histogram.Mean()),
		fmt.Sprintf("Max   : %.2f", p.max),
		fmt.Sprintf("Min   : %.2f", p.min),
		fmt.Sprintf("Var   : %.2f", p.histogram.Variance()),
		fmt.Sprintf("Q1    : %.2f", p.histogram.Quantile(0.25)),
		fmt.Sprintf("Q2    : %.2f", p.histogram.Quantile(0.50)),
		fmt.Sprintf("Q3    : %.2f", p.histogram.Quantile(0.75)),
	})
}

type CounterPanel struct {
	*PanelBase
	labels   []string
	counter  map[string]int
	total    int
	sortMode int
}

func NewCounterPanel(label string) *CounterPanel {
	panel := &CounterPanel{
		PanelBase: NewPanelBase(label),
		labels:    []string{},
		counter:   map[string]int{},
		total:     0,
		sortMode:  SORT_NONE,
	}
	panel.panelType = COUNTER

	chartWidget := NewBarChartWidget()
	chartWidget.BarChart.BarWidth = 5
	panel.SetWidget(chartWidget)
	return panel
}

func (p *CounterPanel) update() {
	chartWidget := p.widget.(*BarChartWidget)

	freqs := make([]int, 0, len(p.labels))
	items := make([]string, 0, len(p.labels)+2)
	items = append(items, p.label, fmt.Sprintf("Total : %d", p.total))

	switch p.sortMode {
	case SORT_ALPHABETICAL:
		sort.Strings(p.labels)
	case SORT_NUMERICAL:
		numericalSort(p.labels)
	}

	for _, label := range p.labels {
		freqs = append(freqs, p.counter[label])
		if p.total > 0 {
			items = append(items,
				fmt.Sprintf("%s : %d (%.2f%%)",
					label, p.counter[label],
					100*float64(p.counter[label])/float64(p.total)))
		}
	}
	chartWidget.Update(p.labels, freqs)
	p.list.Update(items)
}

func (p *CounterPanel) SetSortMode(mode int) {
	p.sortMode = mode
}

func (p *CounterPanel) AddLabel(label string) {
	_, ok := p.counter[label]
	if !ok {
		p.labels = append(p.labels, label)
		p.counter[label] = 0
	}
	p.update()
}

func (p *CounterPanel) Add(label string) {
	_, ok := p.counter[label]
	if !ok {
		p.labels = append(p.labels, label)
		p.counter[label] = 0
	}
	p.counter[label] += 1
	p.total += 1
	p.update()
}
