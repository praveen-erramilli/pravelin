package response

import (
	"encoding/json"
	"net/http"
	"pravelin/logging"
)

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Data       string `json:"data"`
}

func WriteJSONResponse(writer *http.ResponseWriter, response *Response) {
	(*writer).WriteHeader(response.StatusCode)
	encodeError := json.NewEncoder(*writer).Encode(response)
	if encodeError != nil {
		logging.ErrorLog.Println("unable to encode response")
		logging.ErrorLog.Println(encodeError)
	}
}

func WriteJSONErrorResponse(writer *http.ResponseWriter, err error) {
	(*writer).WriteHeader(http.StatusBadRequest)
	encodeError := json.NewEncoder(*writer).Encode(&Response{"failure", http.StatusBadRequest, err.Error()})
	if encodeError != nil {
		logging.ErrorLog.Println("unable to encode response")
		logging.ErrorLog.Println(encodeError)
	}
}
