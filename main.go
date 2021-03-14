package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Result struct {
	Function
}

func main () {
	router := mux.NewRouter()
	router.HandleFunc("/api/calculator/add", Add())

func Calc(w http.ResponseWriter, r *http.Request){
	var countResult []countResult
	var firstNumber []firstNumber
	var secondNumber []secondNumber

	// handle panic error
	if r.Method != "POST" {
		wrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close() // i dont know what does it for
	if err != nil {
		wrapAPIError(w, r, "can't read the body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &firstNumber)
	if err != nil {
		wrapAPIError(w, r, "error unmarshal : " + err.Errpr(), http.StatusInternalServerError)
		return
	}

	for _, v := range firstNumber {
		countResult = append(countResult, Result{
			
		})
	}

}