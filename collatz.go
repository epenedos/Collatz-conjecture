package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)


var initnumber int64
var html string
var intSlice []int
var xslice []string

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

	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	nn, err := strconv.ParseInt(params.Get("nhosts"), 10, 0)
	initnumber = int64(nn)

	compute(initnumber)

	html = f"<html lang='en'>"
	w.Write([]byte(html))
	html = "<head><title>Collatz Conjecture</title> <link rel='stylesheet' href='/assets/style.css' /></head>"
	w.Write([]byte(html))

	html = "<body>"
	w.Write([]byte(html))

	html = "<div class='fixed-header'>"
	w.Write([]byte(html))

	html = "<img id='logo' src='/assets/img/collatz.png' width='75' height='75'> <span id='title'>Collatz Conjecture</span>"
	w.Write([]byte(html))
	html = "</div>"
	w.Write([]byte(html))

	line := charts.NewLine()

	// Set global options
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:         "Collatz Conjecture",
		TitleStyle:    &opts.TextStyle{},
		Link:          "",
		Subtitle:      fmt.Sprintf("%v", " "),
		SubtitleStyle: &opts.TextStyle{},
		SubLink:       "",
		Target:        "",
		Top:           "",
		Bottom:        "",
		Left:          "",
		Right:         "",
	}))

	// Put data into instance
	line.SetXAxis(xslice).
		AddSeries("Collatz", generateLineItems(false))

	line.Render(w)

	html = "<div class='fixed-footer'>KAM Software Solutions</div>"
	w.Write([]byte(html))
	html = "</body>"
	w.Write([]byte(html))
	html = "</html>"
	w.Write([]byte(html))

}

func generateLineItems(raw bool) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(intSlice); i++ {
		items = append(items, opts.LineData{Value: intSlice[i]})
	}
	return items
}

func compute(initnumber int64) {

	intSlice = nil
	xslice = nil
	intSlice = append(intSlice, int(initnumber))
	for {

		n := coll(initnumber)

		initnumber = n
		intSlice = append(intSlice, int(initnumber))
		if initnumber == 1 {
			break
		}

	}

	for i := 0; i < len(intSlice); i++ {
		xslice = append(xslice, fmt.Sprint(i))
	}

	fmt.Printf("intSlice: %v\n", intSlice)
	fmt.Printf("xslice: %v\n", xslice)

}

func coll(r int64) (res int64) {

	if r%2 == 0 {
		res = r / 2
	} else {
		res = r*3 + 1
	}

	return res
}
