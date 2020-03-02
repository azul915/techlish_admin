package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

	"github.com/azul915/techlish_admin/back_end/api"
)

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return tok

}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
			return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getClient(config *oauth2.Config) *http.Client {

	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	return config.Client(context.Background(), tok)

}

func spreadsheetInit() (*http.Client, error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	return client, err

}

func handleAddVocabulary(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	params := r.Form

	wordSlice := params["word"]
	if wordSlice == nil {
		return
	}

	categorySlice := params["category"]
	if categorySlice == nil {
		return
	}

	meanSlice := params["mean"]
	if meanSlice == nil {
		return
	}

	// TODO: コンテナに入って[go get -u google.golang.org/api/sheets/v4]
	// TODO: コンテナに入って[go get -u golang.org/x/oauth2/google]
	// TODO: ParamsをAddVocabularyに渡して、SpreadSheeetを更新する
	// TODO: SpreadSheetのinitを別ファイル定義して呼び出す
	vocabulary.AddVocabulary(&vocabulary.Vocabulary{
		Word: WordSlice[0],
		Category: categorySlice[0],
		Mean: meanSlice[0],
		Any: params["any"][0]
	})
}

func handleRequests() {

	http.HandleFunc("/vocabulary", handleAddVocabulary)
	log.Fatal(http.ListenAndServe(":1998", nil))

}

func main() {
	
	client, err := spreadsheetInit()
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1J3nzfzaUj0Qu8T95R_oz_xpbotrW60pIlP4nI8Ny5Qw"
	readRange := "シート1!A:A"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {

		fmt.Println("No data found.")

	} else {

		newIdx := len(resp.Values) + 1
		writeRange := "シート1!A" + strconv.Itoa(newIdx)
		valueRange := &sheets.ValueRange{
			Values: [][]interface{}{
				[]interface{}{"blue", "名", "青", "補足", "1999/09/15 0:00:00"},
			},
		}

		_, err = srv.Spreadsheets
					.Values
					.Update(spreadsheetId, writeRange, valueRange)
					.ValueInputOption("RAW")
					.Do()

		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet. %v", err)
		}
	}

	handleRequests()

}