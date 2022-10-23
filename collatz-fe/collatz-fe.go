package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var initnumber int64
var html string
var intSlice []int
var xslice []string

var allSlice []int

func main() {

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/graph", httpserver)
	mux.HandleFunc("/", httpserver_home)
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func httpserver_home(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("www/index.html"))
	tpl.Execute(w, nil)

}

func httpserver(w http.ResponseWriter, r *http.Request) {

	var tpl = template.Must(template.ParseFiles("www/graph-header.html"))
	tpl.Execute(w, nil)

	allSlice = nil

	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	nn, err := strconv.ParseInt(params.Get("nhosts"), 10, 0)
	initnumber = int64(nn)

	//compute(initnumber)

	response, err := http.Get("http://localhost:8081/collatz/?init=" + params.Get("nhosts"))
	responseData, err := ioutil.ReadAll(response.Body)

	//BuildGraph(w)
	fmt.Fprintf(w, "%s", string(responseData))
	//initnumber2 := int64(nn) - 1
	//for i := initnumber2; i > 0; i-- {
	//	compute((i))

	//}

	//BuildGraphLim0(w)

	fmt.Printf("response: %v\n", response)

	tpl = template.Must(template.ParseFiles("www/graph-footer.html"))
	tpl.Execute(w, nil)
}

func generateLineItems(raw bool) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(intSlice); i++ {
		items = append(items, opts.LineData{Value: intSlice[i]})
	}
	return items
}

func generateAllLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := len(allSlice) - 1; i > 0; i-- {
		items = append(items, opts.LineData{Value: allSlice[i]})
	}
	return items
}

func compute(number int64) {

	intSlice = nil
	xslice = nil
	intSlice = append(intSlice, int(number))
	for {

		n := coll(number)

		number = n
		intSlice = append(intSlice, int(number))
		if number == 1 {
			break
		}

	}

	for i := 1; i <= len(intSlice); i++ {
		xslice = append(xslice, fmt.Sprint(i))
	}

	//fmt.Printf("intSlice: %v\n", intSlice)
	//fmt.Printf("xslice: %v\n", len(intSlice))
	//fmt.Println("")
	allSlice = append(allSlice, len(intSlice))

}

func coll(r int64) (res int64) {

	if r%2 == 0 {
		res = r / 2
	} else {
		res = r*3 + 1
	}

	return res
}

func BuildGraph(w http.ResponseWriter) {

	line := charts.NewLine()

	line.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		PageTitle:       "teste",
		Width:           "2000",
		Height:          "1000",
		BackgroundColor: "",
		ChartID:         "1",
		AssetsHost:      "",
		Theme:           "",
	}))

	// Set global options
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Collatz Conjecture",
		TitleStyle: &opts.TextStyle{
			Color:      "",
			FontStyle:  "",
			FontSize:   0,
			FontFamily: "",
			Padding:    nil,
			Normal: &opts.TextStyle{
				Color:      "",
				FontStyle:  "",
				FontSize:   0,
				FontFamily: "",
				Padding:    nil,
				Normal:     &opts.TextStyle{},
			},
		},
		Link:     "",
		Subtitle: fmt.Sprintf("%v", " "),
		SubtitleStyle: &opts.TextStyle{
			Color:      "",
			FontStyle:  "",
			FontSize:   0,
			FontFamily: "",
			Padding:    nil,
			Normal:     &opts.TextStyle{},
		},
		SubLink: "",
		Target:  "",
		Top:     "",
		Bottom:  "",
		Left:    "",
		Right:   "",
	}))

	line.SetGlobalOptions(charts.WithToolboxOpts(opts.Toolbox{
		Show:   true,
		Orient: "",
		Left:   "",
		Top:    "",
		Right:  "",
		Bottom: "",
		Feature: &opts.ToolBoxFeature{
			SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{Show: true, Type: "", Name: "abc", Title: "Save"},
			DataZoom:    &opts.ToolBoxFeatureDataZoom{Show: true, Title: map[string]string{"zoom": "Select Area Zooming", "back": "Restore Area"}},
			DataView:    &opts.ToolBoxFeatureDataView{Show: true, Title: "View", Lang: []string{"data view", "turn off", "refresh"}, BackgroundColor: "green"},
			Restore:     &opts.ToolBoxFeatureRestore{Show: true, Title: "Reset"},
		},
	}))

	line.SetGlobalOptions(charts.WithDataZoomOpts(opts.DataZoom{
		Type:  "slider",
		Start: 1,
	}))

	line.SetGlobalOptions(charts.WithYAxisOpts(opts.YAxis{
		Name: "",
		Type: "log",
		Show: true,
	}))

	line.SetGlobalOptions(charts.WithTooltipOpts(opts.Tooltip{
		Show:        true,
		Trigger:     "axis",
		TriggerOn:   "",
		Formatter:   "Element: {b0} <br/> Valor: {c0}",
		AxisPointer: &opts.AxisPointer{Type: "line", Snap: false},
	}))

	// Put data into instance
	line.SetXAxis(xslice).
		AddSeries("Valor", generateLineItems(false))

	line.SetGlobalOptions(charts.WithDataZoomOpts(opts.DataZoom{
		Type:  "slider",
		Start: 1,
	}))

	line.Render(w)

}

func BuildGraphLim0(w http.ResponseWriter) {

	lineAll := charts.NewLine()
	lineAll.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		PageTitle:       "teste1",
		Width:           "2000",
		Height:          "1000",
		BackgroundColor: "",
		ChartID:         "2",
		AssetsHost:      "",
		Theme:           "",
	}))

	lineAll.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:      "Number of Elements ",
		TitleStyle: &opts.TextStyle{},
		Link:       "",
		Subtitle:   fmt.Sprintf("%v", " "),
		SubtitleStyle: &opts.TextStyle{
			Color:      "",
			FontStyle:  "",
			FontSize:   0,
			FontFamily: "",
			Padding:    nil,
			Normal:     &opts.TextStyle{},
		},
		SubLink: "",
		Target:  "",
		Top:     "",
		Bottom:  "",
		Left:    "",
		Right:   "",
	}))

	// Put data into instance
	xslice = nil
	for i := 0; i < len(allSlice); i++ {
		xslice = append(xslice, fmt.Sprint(i+1))
	}

	// fmt.Printf("all xslice: %v\n", xslice)

	lineAll.SetXAxis(xslice).
		AddSeries("# Elements", generateAllLineItems())

	lineAll.SetGlobalOptions(charts.WithTooltipOpts(opts.Tooltip{
		Show:      true,
		Trigger:   "axis",
		TriggerOn: "",
		Formatter: "Element: {b0} <br/> Valor: {c0}",
		AxisPointer: &opts.AxisPointer{
			Type: "line",
			Snap: true,
		},
	}))

	lineAll.SetGlobalOptions(charts.WithToolboxOpts(opts.Toolbox{
		Show:   true,
		Orient: "",
		Left:   "",
		Top:    "",
		Right:  "",
		Bottom: "",
		Feature: &opts.ToolBoxFeature{
			SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
				Show:  true,
				Type:  "",
				Name:  "",
				Title: "Save",
			},
			DataZoom: &opts.ToolBoxFeatureDataZoom{
				Show:  true,
				Title: map[string]string{"zoom": "Select Area Zooming", "back": "Restore Area"},
			},
			DataView: &opts.ToolBoxFeatureDataView{
				Show:            false,
				Title:           "View",
				Lang:            []string{"data view", "turn off", "refresh"},
				BackgroundColor: "green",
			},
			Restore: &opts.ToolBoxFeatureRestore{
				Show:  true,
				Title: "Reset",
			},
		},
	}))

	lineAll.SetGlobalOptions(charts.WithDataZoomOpts(opts.DataZoom{
		Type:  "slider",
		Start: 1,
	}))

	lineAll.Render(w)

}