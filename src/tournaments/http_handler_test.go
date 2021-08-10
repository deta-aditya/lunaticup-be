package tournaments

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreate(t *testing.T) {
	tests := []struct {
		name   string
		body   map[string]interface{}
		expect func(rw *httptest.ResponseRecorder, rb map[string]interface{})
	}{
		{
			name: "it should respond successfully",
			body: map[string]interface{}{
				"name":          "Test",
				"method":        "SINGLE_ELIM",
				"method_config": nil,
				"players": []string{
					"Player 1",
					"Player 2",
				},
			},
			expect: func(rr *httptest.ResponseRecorder, rb map[string]interface{}) {
				if status := rr.Result().StatusCode; status != http.StatusOK {
					t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
				}

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.body)

			r, err := http.NewRequest("POST", "http://localhost:8080/tournaments", bytes.NewReader(jsonBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleCreate)
			handler.ServeHTTP(rr, r)

			var rb map[string]interface{}
			err = json.NewDecoder(rr.Body).Decode(&rb)
			if err != nil {
				t.Fatal(err)
			}

			r.Body.Close()

			tt.expect(rr, rb)
		})
	}
}
