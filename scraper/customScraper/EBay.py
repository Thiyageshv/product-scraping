
class EBayScraper():
	def __init__(self):
		self. imageType = 'jpg'
		self.urlCheck = ""

	def getProductTitle(self, bs):
		try:
			return bs.find(id="vi-lkhdr-itmTitl").get_text()
		except:
			return ""

	def getProductBrand(self, bs):
		try:
			return bs.find(class_="prod-brandName").get_text()
		except:
			return ""

	def getProductSeller(self, bs):
		try:
			return bs.find(class_="mbg-nw").get_text()
		except:
			return ""

	def getProductDescription(self, bs):
		description = ""
		try:
			description += bs.find(id="vi-cond-addl-info").get_text()
			description += bs.find('span', id="hiddenContent").get_Text()
			return description
		except:
			return description

	def getPrice(self, bs):
		try:
			return bs.find(id="prcIsum_bidPrice").get_text()
		except:
			return ""

