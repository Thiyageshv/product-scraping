from customScraper.Amazon import * 
from customScraper.Walmart import *
from customScraper.Alibaba import * 
from customScraper.Generic import *
from customScraper.EBay import *

class ScrapeSwitchController():
	def __init__(self):
		self.scrapSwitch = {
			'AMAZON': AmazonScraper(),
			'WALMART': WalmartScraper(),
			'ALIBABA': AlibabaScraper(),
			'EBAY': EBayScraper()
		}

	def fetchVendorScraper(self, vendor):
		return self.scrapSwitch.get(vendor.upper(), GenericScraper())

	