
class AmazonScraper():
	def __init__(self):
		self. imageType = 'jpg'
		self.urlCheck = 'images/I'

	def getPrice(self, bs):
		try:
			bs.find(id="priceblock_ourprice").get_text()
		except:
			return ""

	def getProductTitle(self, bs):
		try:
			return bs.find(class_="ProductTitle").get_text()
		except:
			return ""

	def getProductBrand(self, bs):
		try:
			return bs.find(class_="prod-brandName").get_text()
		except:
			return ""

	def getProductSeller(self, bs):
		try:
			return bs.find('a', class_="seller-name").get_text()
		except:
			return ""

	def getProductDescription(self, bs):
		try:
			return bs.find(class_="product-short-description-wrapper").get_text()
		except:
			return ""


	def getProductDetails(self, bs):
		for x in bs.find_all('table', id='productDetailsTable'):
			for tag in x.find_all('li'):
				print(tag.get_text())

