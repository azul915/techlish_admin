package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	vocabulary "github.com/azul915/techlish_admin/back_end/api"
)

func handleAddVocabulary(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading JSON data:", err)
		return
	}

	var voc vocabulary.Vocabulary
	json.Unmarshal(jsonData, &voc)

	_, res, err := vocabulary.AddVocabulary(&voc)
	json, err := json.Marshal(res)
	if err != nil {
		return
	}

	log.Println(string(json))
	w.Write(json)

}

func handleRequests() {

	http.HandleFunc("/vocabulary", handleAddVocabulary)
	log.Fatal(http.ListenAndServe(":1998", nil))

}

func main() {

	handleRequests()
}
