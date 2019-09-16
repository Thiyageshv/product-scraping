
class WalmartScraper():
	def __init__(self):
		self. imageType = 'jpeg'
		self.urlCheck = ""

	def getProductTitle(self, bs):
		return bs.find(class_="ProductTitle").get_text()

	def getProductBrand(self, bs):
		return bs.find(class_="prod-brandName")

	def getProductSeller(self, bs):
		return bs.find('a', class_="seller-name").get_text()

	def getProductDescription(self, bs):
		return bs.find(class_="product-short-description-wrapper").get_text()

	def getPrice(self, bs):
		return bs.find(class_="price-group").get_text()

