
class AlibabaScraper():
	def __init__(self):
		self. imageType = 'jpg'
		self.urlCheck = ''

	def getPrice(self, bs):
		try:
			return bs.find('span', class_="pre-inquiry-price").get_text()
		except:
			return ""

	def getProductTitle(self, bs):
		try:
			return bs.find('h1', class_="ma-title").get_text()
		except:
			return ""

	def getProductBrand(self, bs):
		try:	
			for l in bs.find_all('dl'):
				content = l.find('span', class_='attr-name J-attr-name').get_text()
				print(content)
				if content == "Brand Name:":
					return l.find('div', class_='ellipsis').get_text()
		except:
			return self.getProductSeller(bs)

	def getProductSeller(self, bs):
		try:
			return bs.find('a', class_="company-name").get_text().strip("\n").strip(" ")
		except:
			return ""

	def getProductDescription(self, bs):
		try:
			return bs.find(class_="richtext richtext-detail rich-text-description").get_text()
		except:
			return ""

