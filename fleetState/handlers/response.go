package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

// SuccessArrRespond -> response formatter
func SuccessArrRespond(fields []interface{}, writer http.ResponseWriter) {
	// var fields["status"] := "success"
	_, err := json.Marshal(fields)
	type data struct {
		Data       []interface{}
		StatusCode int
		Message    string
	}
	temp := &data{Data: fields, StatusCode: 200, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessRespond -> response formatter
func SuccessRespond(fields interface{}, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)
	type data struct {
		Data       interface{}
		StatusCode int
		Message    string
	}
	temp := &data{Data: fields, StatusCode: 200, Message: "success"}
	if err != nil {
		log.Println(fields)
		ServerErrResponse(err.Error(), writer)
		return
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessResponse -> success formatter
func SuccessResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		StatusCode int
		Message    string
	}
	temp := &errdata{StatusCode: 200, Message: msg}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// ErrorResponse -> error formatter
func ErrorResponse(error string, writer http.ResponseWriter) {
	type errdata struct {
		StatusCode int
		Message    string
	}
	temp := &errdata{StatusCode: 400, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}

// ServerErrResponse -> server error formatter
func ServerErrResponse(error string, writer http.ResponseWriter) {
	type servererrdata struct {
		StatusCode int
		Message    string
	}
	temp := &servererrdata{StatusCode: 500, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(temp)
}

// ValidationResponse -> user input validation
func ValidationResponse(fields map[string][]string, writer http.ResponseWriter) {
	//Create a new map and fill it
	response := make(map[string]interface{})
	response["Errors"] = fields
	response["StatusCode"] = 422
	response["Message"] = "validation error"

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(writer).Encode(response)
}
