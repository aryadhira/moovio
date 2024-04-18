package helper

import (
	"encoding/json"
	"moovio/libs/models"
	"net/http"
)

func HttpGetRequest(url string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	var client = &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func WriteJSON(w http.ResponseWriter, status int, message string, data any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	response := models.ResponseModels{}
	if status == http.StatusOK {
		response.Success = true
	} else {
		response.Success = false
	}
	response.Message = message
	response.Data = data

	return json.NewEncoder(w).Encode(response)
}
