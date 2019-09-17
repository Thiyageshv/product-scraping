from flask import jsonify, request
from scrapeLogic import *

INTERNALSERVERERROR = 510
UNAUTHORIZED = 401
NOTFOUND = 404
BADREQUEST = 400

class APIResponse():
	def __init__(self):
		self.errorMessages = {
			'TIMEOUT': 'Failed Scraping! Product page timeout',
			'NOTFOUND': 'Failed Scraping! Product page not found',
			'INVALIDREQUEST': 'Invalid Request! Please check methods and parameters',
			'INTERNALERROR': 'Scraping failed! '
		}
		self.errorResponse = {
			'status': 0,
			'message': "",
		} 
		self.successResponse = {
			"status": 200,
			"mimetype": 'application/json',
			"response": {}
		}
	def constructErrorMessage(self, status, msg):
		error  = self.errorResponse
		error['status'] = status
		error['message'] = msg
		return error

	def constructSuccessResponse(self, resp):
		success = self.successResponse
		success['response'] = resp
		return success

class APIController():
	def __init__(self):
		self.sbl = Scraper()
		self.response = APIResponse()

	def fetchProductInfo(self):
		if request.method != "POST":
			return jsonify(self.response.constructErrorMessage(BADREQUEST, self.response.errorMessages['INVALIDREQUEST']))
		purl = request.get_json(silent=True)['producturl']
		resp, err = self.sbl.fetchInformationEntry(purl)
		if err is not None:
			errmsg = self.response.constructErrorMessage(INTERNALSERVERERROR, self.response.errorMessages['INTERNALERROR'] + str(err))
			return jsonify(errmsg)
		success = self.response.constructSuccessResponse(resp)
		return jsonify(success)