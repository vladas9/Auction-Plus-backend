package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		aErr := fmt.Errorf("Encoding of object of type %v failed", reflect.TypeOf(v))
		Logger.Error(aErr)
		return aErr
	}
	return nil
}
