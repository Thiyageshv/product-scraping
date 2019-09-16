from vendorScraper import * 
from urllib.request import urlopen
from urllib.parse import urlencode
from bs4 import BeautifulSoup
import re
import tldextract
import requests

class Scraper():
	def __init__(self):
		self.scrapSwitch = ScrapeSwitchController()
		self.headers = {'User-agent': 'Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.120 Safari/537.36'}
		self.headerdata = urlencode(self.headers)
	
	def scrape(self, url):
		req = requests.get(url, headers=self.headers)
		bs = BeautifulSoup(req.text, 'html.parser')
		return bs

	def getVendor(self, url):
		extracted = tldextract.extract(url)
		return extracted.domain

	def fetchVendorScraperObject(self, vendor):
		return self.scrapSwitch.fetchVendorScraper(vendor)

	def fetchImage(self, bs, vendorScraper):
		print(vendorScraper.imageType)
		images = bs.find_all('img', {'src':re.compile('.' + vendorScraper.imageType)})
		print(images)
		for image in images: 
			print(image['src'])
			if vendorScraper.urlCheck == "" or vendorScraper.urlCheck in image['src']:
				return image['src']


	def fetchInformation(self, url):
		vendor = self.getVendor(url)
		print(vendor)
		vo = self.fetchVendorScraperObject(vendor)
		soup = self.scrape(url)
		iurl = self.fetchImage(soup, vo)
		price = vo.getPrice(soup)
		brand = vo.getProductSeller(soup)
		title = vo.getProductTitle(soup)
		return iurl, price, title, brand



url = "https://www.amazon.com/Samsung-QN65Q6FN-FLAT-QLED-Smart/dp/B079V1MSQ1"
#"https://www.walmart.com/ip/Bounty-Full-Sheet-Paper-Towels-White-12-Super-Rolls-22-Regular-Rolls/459722147?athcpid=459722147&athpgid=athenaItemPage&athcgid=null&athznid=PWVUB&athieid=v0&athstid=CS004&athguid=6054e5d7-1d0-16d37a71cfc7a3&athancid=null&athena=true"
# "https://www.walmart.com/ip/AT-T-Apple-iPhone-11-Pro-Max-512GB-Silver-Upgrade-Only/745930459"
#"https://www.amazon.com/Samsung-QN65Q6FN-FLAT-QLED-Smart/dp/B079V1MSQ1"
#"https://www.amazon.com/PlayStation-4-Slim-1TB-Console/dp/B071CV8CG2/ref=sr_1_1?keywords=playstation&qid=1568515297&s=electronics&sr=1-1"
s = Scraper()
print(s.fetchInformation(url))