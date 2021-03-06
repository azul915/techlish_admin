package vocabulary

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/api/sheets/v4"
)

type Vocabulary struct {
	Word      string `json:"word"`
	Category  string `json:"category"`
	Mean      string `json:"mean"`
	Any       string `json:"any"`
	CreatedAt time.Time
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AddVocabulary(v *Vocabulary) (int, interface{}, error) {

	client, err := SpreadsheetInit()
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}

	sheetService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1J3nzfzaUj0Qu8T95R_oz_xpbotrW60pIlP4nI8Ny5Qw"
	readRange := "シート1!A:A"
	resp, err := sheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
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
				[]interface{}{
					v.Word,
					v.Category,
					v.Mean,
					v.Any,
					"1999/09/15 0:00:00",
				},
			},
		}

		_, err = sheetService.Spreadsheets.
			Values.
			Update(spreadsheetId, writeRange, valueRange).
			ValueInputOption("RAW").
			Do()

		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet. %v", err)
			res := Response{
				Code:    http.StatusInternalServerError,
				Message: "something went wrong",
			}
			return http.StatusInternalServerError, res, err
		}

	}

	res := Response{
		Code:    http.StatusOK,
		Message: "success",
	}

	return http.StatusOK, res, nil

}
