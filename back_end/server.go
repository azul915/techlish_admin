package main

import (

	"log"
	"net/http"

	"github.com/azul915/techlish_admin/back_end/api"
)

func handleAddVocabulary(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	params := r.Form

	wordSlice := params["word"]
	if wordSlice == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("BadRequest: parameter[word] is empty")
		return
	}

	categorySlice := params["category"]
	if categorySlice == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("BadRequest: parameter[category] is empty")
		return
	}

	meanSlice := params["mean"]
	if meanSlice == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("BadRequest: parameter[mean] is emtpy")
		return
	}

	voc := vocabulary.Vocabulary{
		Word: wordSlice[0],
		Category: categorySlice[0],
		Mean: meanSlice[0],
		Any: params["any"][0],
	}

	statusCode, res, err := vocabulary.AddVocabulary(&voc)
	
}

func handleRequests() {

	http.HandleFunc("/vocabulary", handleAddVocabulary)
	log.Fatal(http.ListenAndServe(":1998", nil))

}

func main() {

	handleRequests()
}