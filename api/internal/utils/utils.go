package utils

import (
	structure "blog-backend/api/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ThrowMethodNotAllowedError(resWriter http.ResponseWriter, request *http.Request, expectedMethod string) bool {
	if request.Method != expectedMethod {
		CreateResponse(resWriter, http.StatusMethodNotAllowed, "Method Not Allowed")
		return true
	} else {
		return false
	}
}

func CreateResponse(resWriter http.ResponseWriter, resCode int, message string, op ...string) {
	resErr := structure.Error{}
	resErr.Message = message
	if len(op) > 0 {
		resErr.Op = op[0]
	}
	resWriter.Header().Set("Content-Type", "Application/json")
	resWriter.WriteHeader(resCode)
	json.NewEncoder(resWriter).Encode(resErr)
}

func ParseFlexibleDate(input string) (string, error) {
	layouts := []string{
		"2006-01-02",
		"02-01-2006",
		"02/01/2006",
		"01/02/2006",
		"2 Jan 2006",
		"2 January 2006",
		"January 2, 2006",
		"2006/01/02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, input); err == nil {
			// Success: convert to yyyy-mm-dd
			return t.Format("2006-01-02"), nil
		}
	}

	return "", fmt.Errorf("invalid date format: %s", input)
}
