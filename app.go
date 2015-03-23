package termeter

type App struct {
	*VBox
	panels []Panel
}

func NewApp() *App {
	return &App{
		VBox:   NewVBox(),
		panels: []Panel{},
	}
}

func (app *App) GetPanel(i int) Panel {
	return app.panels[i]
}

func (app *App) AddPanel(label string, panelType int, options map[string]string) {
	var panel Panel
	switch panelType {
	case LINE:
		panel = NewLineChartPanel(label)
		lineChartPanel := panel.(*LineChartPanel)
		lineChartPanel.SetMode(options["line-mode"])
	case CDF:
		panel = NewCDFPanel(label)
		cdfChartPanel := panel.(*CDFPanel)
		cdfChartPanel.SetMode(options["line-mode"])
	case COUNTER:
		panel = NewCounterPanel(label)
		switch options["sort-mode"] {
		case "alphabetical":
			panel.(*CounterPanel).SetSortMode(SORT_ALPHABETICAL)
		case "numerical":
			panel.(*CounterPanel).SetSortMode(SORT_NUMERICAL)
		}
	default:
		return
	}
	app.VBox.AddBoxes(panel)
	app.panels = append(app.panels, panel)
}

func (app *App) Render() {
	for _, panel := range app.panels {
		switch panel.GetType() {
		case LINE:
			panel.(*LineChartPanel).Update()
		case CDF:
			panel.(*CDFPanel).Update()
		}
	}
	Render(app.VBox)
}
