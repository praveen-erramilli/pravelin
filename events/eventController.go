package events

import (
	"net/http"
	"pravelin/logging"
	"pravelin/response"
)

func (eventRepository *EventRepository) EventRequestHandler(writer http.ResponseWriter, request *http.Request) {
	logging.InfoLog.Println("Request received for eventRequestHandler")

	setHeaders(&writer)
	switch request.Method {
	case http.MethodOptions:
		writer.WriteHeader(http.StatusOK)
	case http.MethodPost:
		logging.InfoLog.Println("Handling POST request")

		input, err := parseEventInput(request)
		if err == nil {
			err = input.validate()
		}
		if err != nil {
			logging.ErrorLog.Println(err)
			response.WriteJSONErrorResponse(&writer, err)
			return
		}

		event := eventRepository.GetOrCreateEvent(input)
		err = event.ProcessEvent(input)
		if err != nil {
			logging.ErrorLog.Println(err)
			response.WriteJSONErrorResponse(&writer, err)
			return
		}

		response.WriteJSONResponse(&writer, &response.Response{Status: "success", StatusCode: http.StatusOK, Data: "Successfully processed the event"})
	default:
		response.WriteJSONResponse(&writer, &response.Response{Status: "failure", StatusCode: http.StatusMethodNotAllowed, Data: http.StatusText(http.StatusMethodNotAllowed)})
	}
}

func setHeaders(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
	(*writer).Header().Set("Access-Control-Allow-Headers", "*")
	(*writer).Header().Set("Access-Control-Allow-Methods", http.MethodPost+", "+http.MethodOptions)
	(*writer).Header().Set("Content-Type", "application/json; charset=utf-8")
}
