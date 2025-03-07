package exception

import (
	"FarhanKurnia/simple-golang-restful-api/helper"
	"FarhanKurnia/simple-golang-restful-api/model/web"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	internalServerError(writer, request, err)
}

func internalServerError(writer http.ResponseWriter, _, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: false,
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
