package flappy_trade

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldSell(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body
		var transaction Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			t.Fatalf("Error decoding request: %v", err)
		}

		// Based on the transaction, decide the mock response
		var response Prediction
		if transaction.Amount > 1000 { // Mock condition for selling
			response = Prediction{Action: "sell", ExpectedValue: 0.5}
		} else {
			response = Prediction{Action: "hold", ExpectedValue: 0.3}
		}

		// Send a JSON response back
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Test cases
	tests := []struct {
		name        string
		transaction *Transaction
		want        bool
	}{
		{"TestSell", &Transaction{Amount: 2000}, true},
		{"TestHold", &Transaction{Amount: 500}, false},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldSell(tt.transaction, server.URL)
			if got != tt.want {
				t.Errorf("should_sell() = %v, want %v", got, tt.want)
			}
		})
	}
}
