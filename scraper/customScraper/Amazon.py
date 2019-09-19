
class AmazonScraper():
	def __init__(self):
		self. imageType = 'jpg'
		self.urlCheck = 'images/I'

	def getPrice(self, bs):
		try:
			bs.find('span', class_="a-price").get_text()
		except Exception as e:
			print(e)
			return ""

	def getProductTitle(self, bs):
		try:
			for node in bs.find_all('div', class_="a-row r4m-sou-product-name"):
				print(node.get_text())
				return node.get_text()
			return ""
		except:
			return ""

	def getProductBrand(self, bs):
		try:
			return bs.find(class_="prod-brandName").get_text()
		except:
			return ""

	def getProductSeller(self, bs):
		try:
			return bs.find('a', id="bylineInfo").get_text()
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

