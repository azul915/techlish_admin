package main

import (
	"encoding/json"
	"log"
	"net/http"

	vocabulary "github.com/azul915/techlish_admin/back_end/api"
)

func handleAddVocabulary(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	ws := r.Form["word"]
	if ws == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("BadRequest: parameter[word] is empty")
		return
	}

	cs := r.Form["category"]
	if cs == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("BadRequest: parameter[category] is empty")
		return
	}

	ms := r.Form["mean"]
	if ms == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("BadRequest: parameter[mean] is empty")
		return
	}

	as := r.Form["any"]

	voc := vocabulary.Vocabulary{
		Word:     ws[0],
		Category: cs[0],
		Mean:     ms[0],
		Any:      as[0],
	}

	_, res, err := vocabulary.AddVocabulary(&voc)
	json, err := json.Marshal(res)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	w.Write(json)
}

func handleRequests() {

	http.HandleFunc("/vocabulary", handleAddVocabulary)
	log.Fatal(http.ListenAndServe(":1998", nil))

}

func main() {

	handleRequests()
}
