from vendorScraper import * 
from urllib.request import urlopen
from urllib.parse import urlencode
from bs4 import BeautifulSoup
import re
import tldextract
import requests

class Product():
	def __init__(self):
		self.pname = ""
		self.purl = ""
		self.description = ""
		self.price = ""
		self.seller = ""
		self.retailer = ""
		self.brand = "" 

	def setPrice(self, price):
		self.price = price

	def setSeller(self, seller):
		self.seller = seller

	def setRetailer(self, retailer):
		self.retailer = retailer

	def setDescription(self, description):
		self.description = description

	def setBrand(self, brand):
		self.brand = brand

	def setURL(self, url):
		self.purl = url

	def setTitle(self, title):
		self.pname = title

	def getJSON(self):
		return {"name": self.pname, "imageurl": self.purl, "retailer": self.retailer, "metainfo": {
										"description" : self.description,
										"price": self.price,
										"seller": self.seller,
										"brand": self.brand }
			    }


class Scraper():
	def __init__(self):
		self.scrapSwitch = ScrapeSwitchController()
		self.headers = {'User-agent': 'Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36'}
		self.headerdata = urlencode(self.headers)
	
	def scrape(self, url):
		try:
			req = requests.get(url, headers=self.headers)
			bs = BeautifulSoup(req.text, 'html.parser')
			return bs, None
		except Exception as e:
			return None, e

	def getRetailer(self, url):
		extracted = tldextract.extract(url)
		return extracted.domain

	def fetchVendorScraperObject(self, vendor):
		return self.scrapSwitch.fetchVendorScraper(vendor)

	def fetchImage(self, bs, vendorScraper):
		images = bs.find_all('img', {'src':re.compile('.' + vendorScraper.imageType)})
		for image in images: 
			if vendorScraper.urlCheck == "" or vendorScraper.urlCheck in image['src']:
				return image['src']
		return ""


	def fetchInformation(self, url):
		retailer = self.getRetailer(url)
		product = Product()
		product.setRetailer(retailer)
		vo = self.fetchVendorScraperObject(retailer)
		soup, err = self.scrape(url)
		if err is not None:
			return product, err
		product.setURL(self.fetchImage(soup, vo))
		product.setPrice(vo.getPrice(soup))
		product.setSeller(vo.getProductSeller(soup))
		product.setDescription(vo.getProductDescription(soup))
		product.setTitle(vo.getProductTitle(soup))
		product.setBrand(vo.getProductBrand(soup))
		return product, None

	def fetchInformationEntry(self, url):
		product, err = self.fetchInformation(url)
		return product.getJSON(), err 

if __name__ == '__main__':	
	url = "https://www.amazon.com/Samsung-QN65Q6FN-FLAT-QLED-Smart/dp/B079V1MSQ1"
	#"https://www.walmart.com/ip/Bounty-Full-Sheet-Paper-Towels-White-12-Super-Rolls-22-Regular-Rolls/459722147?athcpid=459722147&athpgid=athenaItemPage&athcgid=null&athznid=PWVUB&athieid=v0&athstid=CS004&athguid=6054e5d7-1d0-16d37a71cfc7a3&athancid=null&athena=true"
	# "https://www.walmart.com/ip/AT-T-Apple-iPhone-11-Pro-Max-512GB-Silver-Upgrade-Only/745930459"
	#"https://www.amazon.com/Samsung-QN65Q6FN-FLAT-QLED-Smart/dp/B079V1MSQ1"
	#"https://www.amazon.com/PlayStation-4-Slim-1TB-Console/dp/B071CV8CG2/ref=sr_1_1?keywords=playstation&qid=1568515297&s=electronics&sr=1-1"
	s = Scraper()
	print(s.fetchInformation(url))