package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type details struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allDetails []details

var Data = allDetails{
	{
		ID:          " XI2471",
		Title:       " golang",
		Description: "learning",
	},
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "starting learning")
}
func createDetails(w http.ResponseWriter, r *http.Request) {
	var newData details
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error")
	}
	defer r.Body.Close()

	json.Unmarshal(reqBody, &newData)
	//Data = append(Data, newData)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newData) //newData
}

/*func getData(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleData := range Data {
		if singleData.ID == eventID {
			json.NewEncoder(w).Encode(singleData)
		}
	}
}
router.HandleFunc("/Data/{id}", getData).Methods("GET")
*/
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", start)
	router.HandleFunc("/Data", createDetails).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router)) //http://localhost:8080/
}
