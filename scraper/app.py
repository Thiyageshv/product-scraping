from flask import Flask
from controller import *

app = Flask(__name__)


class App():
	def __init__(self):
		self.apiController = APIController()
		self.addRules()

	def addRules(self):
		app.add_url_rule('/scraper/api/v1/scrape', 'fetchProductInfo', self.apiController.fetchProductInfo, methods=["POST"])

if __name__ == '__main__':		
	appobj = App()
	app.run(debug=True)
