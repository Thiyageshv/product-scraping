from customScraper.Amazon import * 
from customScraper.Walmart import * 
from customScraper.Generic import * 

class ScrapeSwitchController():
	def __init__(self):
		self.scrapSwitch = {
			'AMAZON': AmazonScraper(),
			'WALMART': WalmartScraper(),
		}

	def fetchVendorScraper(self, vendor):
		return self.scrapSwitch.get(vendor.upper(), GenericScraper())

	