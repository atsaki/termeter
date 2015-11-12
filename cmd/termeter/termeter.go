package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andrew-d/go-termutil"
	"github.com/atsaki/termeter"
	"gopkg.in/gizak/termui.v1"
	"github.com/nsf/termbox-go"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	cli = kingpin.New("termeter", "Visualize data in the terminal")

	inputFile = cli.Arg("file", "Input File").ExistingFile()

	delimiter = cli.Flag(
		"delimiter",
		"Delimiter of input",
	).Short('d').Default("\t").String()

	panelTypes = cli.Flag(
		"types",
		"panel types of each column (L: line, C:counter, D:cdf, other: auto)",
	).Short('t').String()

	theme = cli.Flag(
		"theme",
		"Theme of charts",
	).Short('T').Default("default").String()

	lineMode = cli.Flag(
		"line-mode",
		"Mode of line chart. (braille, dot)",
	).Short('M').Default("braille").String()

	dataLabelMode = cli.Flag(
		"data-label",
		"Data labels used by line chart (count, first, time)",
	).Short('L').Default("count").String()

	sortMode = cli.Flag(
		"sort-mode",
		"Mode of counter label sorting (none, alphabetical, numeric)",
	).Short('S').Default("none").String()
)

var (
	tickInterval = 100 * time.Millisecond
)

func update(app *termeter.App, dataLabelMode string, record []string) {
	label := ""
	switch dataLabelMode {
	case "first":
		label = record[0]
		record = record[1:]
	case "time":
		now := time.Now()
		label = fmt.Sprintf("%02d:%02d:%02d",
			now.Hour(), now.Minute(), now.Second())
	}

	for i := 0; i < len(record); i++ {
		panel := app.GetPanel(i)
		panelType := panel.GetType()
		switch panelType {
		case termeter.LINE:
			x, _ := strconv.ParseFloat(record[i], 64)
			panel.(*termeter.LineChartPanel).Add(x, label)
		case termeter.CDF:
			x, _ := strconv.ParseFloat(record[i], 64)
			panel.(*termeter.CDFPanel).Add(x)
		case termeter.COUNTER:
			panel.(*termeter.CounterPanel).Add(record[i])
		}
	}
}

func main() {

	var err error

	cli.Parse(os.Args[1:])
	options := map[string]string{}
	options["line-mode"] = *lineMode
	options["sort-mode"] = *sortMode

	// read file or stdin
	var reader *csv.Reader
	if *inputFile != "" {
		f, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		reader = csv.NewReader(f)
		defer f.Close()
	} else if !termutil.Isatty(os.Stdin.Fd()) {
		reader = csv.NewReader(os.Stdin)
	} else {
		return
	}
	reader.Comma = []rune(*delimiter)[0]

	// initialize UI
	err = termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	termui.UseTheme(*theme)
	app := termeter.NewApp()

	// read first line as labels
	labels, err := reader.Read()
	if err != nil {
		panic(err)
	}
	// read second line to determine panel type
	record, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}

	if *dataLabelMode == "first" {
		labels = labels[1:]
		record = record[1:]
	}

	types := make([]rune, len(labels), len(labels))
	copy(types, []rune(strings.ToUpper(*panelTypes)))
	for i, label := range labels {
		var panelType int
		switch types[i] {
		case 'L':
			panelType = termeter.LINE
		case 'D':
			panelType = termeter.CDF
		case 'C':
			panelType = termeter.COUNTER
		default:
			_, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				panelType = termeter.COUNTER
			}
		}
		app.AddPanel(label, panelType, options)
	}
	update(app, *dataLabelMode, record)
	app.Render()

	// start polling
	evt := make(chan termbox.Event)
	tick := time.Tick(tickInterval)
	dataChan := make(chan []string)
	go func() {
		for {
			evt <- termbox.PollEvent()
		}
	}()
	go func() {
		for {
			r, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			dataChan <- r
		}
	}()

	for {
		select {
		case e := <-evt:
			if e.Type == termbox.EventKey && e.Ch == 'q' {
				return
			}
			if e.Type == termbox.EventResize {
				app.Render()
			}
		case record := <-dataChan:
			update(app, *dataLabelMode, record)
		case <-tick:
			app.Render()
		}
	}
}
