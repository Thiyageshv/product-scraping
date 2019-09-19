package app 

import (
	"log"
	util "product-scraping/lib_utilities"
	cas "product-scraping/lib_cassandra"
)

type ScrapeSummary struct {
	ImageURL []cas.ProductImageInfo `json:"imageurls"`
	MetaInfo cas.ProductMetaInfo `json:"metainfo"`
}


func (a *App) getImages(purlid int64, pid int64) ([]cas.ProductImageInfo, error) {
	results := []cas.ProductImageInfo{}
	pinfo, err := a.CasCursor.GetProductInfo(purlid, pid)
	if err != nil {
		return results, err
	}
	for _, iurlid := range pinfo.IURLGroup {
		pimginfo, err := a.CasCursor.GetProductImageInfo(iurlid, purlid)
		if err != nil {
			continue
		}
		results = append(results, pimginfo)
	}
	return results, nil
}

func (a *App) getMetaInfo(purlid int64, pid int64) (cas.ProductMetaInfo, error) {
	return a.CasCursor.GetProductMetaInfo(purlid, pid)
}

func (a *App) getInfo() ([]ScrapeSummary, error) {
	log.Println("Api entry")
	results := []ScrapeSummary{}
	pages, err := a.CasCursor.GetProductPageInfo()
	for _, page := range pages {
		result := ScrapeSummary{}
		result.ImageURL, err = a.getImages(page.PURLID, page.PID)
		result.MetaInfo, err = a.getMetaInfo(page.PURLID, page.PID)
		if err != nil || (result.MetaInfo.PID == 0 && result.MetaInfo.PURLID == 0 && result.MetaInfo.MetaInfoID == 0) {
			continue
		}
		if len(result.ImageURL) == 0 {
			result.MetaInfo.Title = "Unable to scrape the product page in " + result.MetaInfo.Retailer
		}
		results = append(results, result)
	}
	return results, nil
}

func (a *App) getProductMetrics() ([]cas.ProductPageInfo, error) {
	results, err := a.CasCursor.GetProductPagesMetrics()
	return results, err
}

func (a *App) addProductPage(purl string, pname string, user string) error {
	pid := int64(util.Hash32(pname))
	purlid := int64(util.Hash32(purl))
	err := a.CasCursor.AddProductPageInfo(pid, purlid, pname, purl, user)
	return err
}	