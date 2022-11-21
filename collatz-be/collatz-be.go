package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var initnumber int64
var intSlice []int
var xslice []int

var allSlice []int

type resultscollatz struct {
	Value        int   `json:"value"`
	List_results []int `json:"list_results"`
}

type JsonResponse struct {
	Type    string         `json:"type"`
	Data    resultscollatz `json:"data"`
	Message string         `json:"message"`
}

var myresults resultscollatz
var response JsonResponse

func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/collatz/{id}", Collatz).Methods("GET")

	// serve the app
	fmt.Println("Server at 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func Collatz(w http.ResponseWriter, r *http.Request) {
	init := mux.Vars(r)["id"]
	i, err := strconv.ParseInt(init, 10, 64)
	if err != nil {
		panic(err)
	}

	compute(int64(i))

	myresults = resultscollatz{Value: int(initnumber), List_results: intSlice}
	response = JsonResponse{
		Type:    "Result",
		Data:    myresults,
		Message: "Successeful",
	}
	json.NewEncoder(w).Encode(response)

}

func compute(number int64) {
	initnumber = number
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
		xslice = append(xslice, i)
	}

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
