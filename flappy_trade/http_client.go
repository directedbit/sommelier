package flappy_trade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Prediction struct {
	Action        string  `json:"action"`
	ExpectedValue float64 `json:"expected_value"`
}

func ShouldSell(transaction *Transaction, url string) bool {
	// Data to send
	// Convert data to JSON
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	// Make HTTP POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		// Read response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Response:", string(body))

		var prediction Prediction
		err = json.Unmarshal([]byte(body), &prediction)
		if err != nil {
			panic(err)
		}
		if prediction.Action == "sell" {
			return true
		}
	case 400:
		fmt.Println("Bad request")
	case 500:
		fmt.Println("Internal server error")
	default:
		fmt.Println("Unknown error", resp.StatusCode)
	}
	return false
}
