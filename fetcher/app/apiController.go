package app

import (
	"encoding/json"
	"net/http"
	util "product-scraping/lib_utilities"
)



func (a *App) getInfoEntry(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(400), 400)
		util.SendJsonErrorResponse(w, util.HTTPInvalidRequest, util.InvalidRequestMessage, "")
		return
	}
	responseObj, err := a.getInfo()
	if err != nil {
		util.SendJsonErrorResponse(w, util.HTTPInternalError, util.InternalErrorMessage, err.Error())
		return
	}
	responseJSON, err := json.Marshal(&responseObj)
	if err != nil {
		util.SendJsonErrorResponse(w, util.HTTPInternalError, util.InternalErrorMessage, err.Error())
		return
	}
	util.SendJsonSuccessResponse(w, util.HTTPSuccess, "", responseJSON)
}


func (a *App) getMetricsEntry(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(400), 400)
		util.SendJsonErrorResponse(w, util.HTTPInvalidRequest, util.InvalidRequestMessage, "")
		return
	}
	responseObj, err := a.getProductMetrics()
	if err != nil {
		util.SendJsonErrorResponse(w, util.HTTPInternalError, util.InternalErrorMessage, err.Error())
		return
	}
	responseJSON, err := json.Marshal(&responseObj)
	if err != nil {
		util.SendJsonErrorResponse(w, util.HTTPInternalError, util.InternalErrorMessage, err.Error())
		return
	}
	util.SendJsonSuccessResponse(w, util.HTTPSuccess, "", responseJSON)
}


