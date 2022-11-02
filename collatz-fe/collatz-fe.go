package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	//"github.com/go-echarts/go-echarts/v2/charts"
	//"github.com/go-echarts/go-echarts/v2/opts"
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
	fmt.Println("http://collatz-be:8081/collatz/?init=" + params.Get("nhosts"))
	response, err := http.Get("http://collatz-be:8081/collatz/?init=" + params.Get("nhosts"))
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
