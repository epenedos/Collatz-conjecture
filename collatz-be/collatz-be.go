package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var initnumber int64
var html string
var intSlice []int
var xslice []string

var allSlice []int

func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/collatz/", Collatz).Methods("GET")

	// serve the app
	fmt.Println("Server at 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func Collatz(w http.ResponseWriter, r *http.Request) {
	init := r.FormValue("init")
	i, err := strconv.ParseInt(init, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(init)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", init, "")
	compute(int64(i))
	fmt.Printf("xslice: %v\n", xslice)

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