
class GenericScraper():
	def __init__(self):
		self. imageType = 'jpg'
		self.urlCheck = ""

	def getPrice(self, bs):
		return bs.find(class_="price-group").get_text()