package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"os"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type resultscollatz struct {
	Value        int   `json:"value"`
	List_results []int `json:"list_results"`
}

type JsonResponse struct {
	Type    string         `json:"type"`
	Data    resultscollatz `json:"data"`
	Message string         `json:"message"`
}

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
	if err != nil {
		log.Fatal(err)
	}
	initnumber = int64(nn)
	env:= os.Getenv("BACKEND")
	
	response, err := http.Get("http://" + env + ":8081/collatz/" + params.Get("nhosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var abc JsonResponse

	json.Unmarshal(responseData, &abc)

	//json.NewDecoder(response.Body).Decode(&abc)

	fmt.Fprintf(w, "<ol> ")
	for i, p := range abc.Data.List_results {
		fmt.Fprintln(w, "<li>"+strconv.Itoa(i+1)+":"+strconv.Itoa(p)+"</li>")
		xslice = append(xslice,strconv.Itoa(i+1))
	}
	fmt.Fprintf(w, "</ol> ")

	intSlice= abc.Data.List_results

    BuildGraph(w)

	//initnumber2 := int64(nn) - 1
	//for i := initnumber2; i > 0; i-- {
	//	compute((i))

	//}

	BuildGraphLim0(w)

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