
class WalmartScraper():
	def __init__(self):
		self. imageType = 'jpeg'
		self.urlCheck = ""

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
			return bs.find(class_="product-short-description-wrapper").get_text().strip("&quot;")
		except:
			return ""

	def getPrice(self, bs):
		try:
			return bs.find(class_="price-group").get_text()
		except:
			return ""

