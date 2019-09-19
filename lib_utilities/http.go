
package utilities

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"bytes"
	"strings"
	// "log"
	"io/ioutil"
)

var HTTPSuccess = 200
var HTTPInternalError = 500
var HTTPForbidden = 403
var HTTPPageNotFound = 404
var HTTPInvalidRequest = 400
var HTTPPreConditionFailure = 412
var HTTPTimeout = 504

var InvalidRequestMessage = "Invalid Request. Please check the API method and parameters"
var PageNotFoundMessage = "Page/Endpoint not found"
var InternalErrorMessage = "Internal Server Error. Please contact support"
var InvalidResponseStatus = "Received an invalid status response for the request"
var ForbiddenError = "Transaction is not authorized. Try again with right set of headers"

type JsonSuccessResponse struct {
	Statuscode int             `json:"status"`
	Message    string          `json:"message"`
	Response   json.RawMessage `json:"response"`
}

type JsonErrorResponse struct {
	Statuscode int    `json:"status"`
	Message    string `json:"message"`
	Logs       string `json:"logs"`
}


func SendJsonSuccessResponse(w http.ResponseWriter, statuscode int, msg string, responseJSON json.RawMessage) {
	outputJSONstruct := JsonSuccessResponse{Statuscode: statuscode, Message: msg, Response: responseJSON}
	outputJSON, _ := json.Marshal(outputJSONstruct)
	w.Write([]byte(outputJSON))
}

func SendJsonErrorResponse(w http.ResponseWriter, statuscode int, errorJSON string, comments string) {
	outputJSONstruct := JsonErrorResponse{Statuscode: statuscode, Message: errorJSON, Logs: comments}
	outputJSON, _ := json.Marshal(outputJSONstruct)
	http.Error(w, string(outputJSON), statuscode)
}

func getIntStatus(status string) (int, error) {
	if status == "" {
		return HTTPInvalidRequest, errors.New(InvalidResponseStatus)
	}
	status = strings.Split(status, " ")[0]
	intstatus, err := strconv.Atoi(status)
	if err != nil {
		return HTTPInternalError, err
	}
	return intstatus, nil
}

func SendPostRequest(apiurl string, params map[string]string) (int, string, error) {
	var result string
	reqdata, err := json.Marshal(params)
	if err != nil {
		return HTTPInvalidRequest, result, err
	}

	resp, err := http.Post(apiurl, "application/json", bytes.NewBuffer(reqdata))
	if err != nil {
		return HTTPInternalError, result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return 200, string(body), err
}
