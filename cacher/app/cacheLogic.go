package app

import (
	"log"
	"time"
	"encoding/json"
	util "product-scraping/lib_utilities"
)

type Product struct {
	Name string `json:"name"`
	Retailer string `json:"retailer"`
	ImageURL []string `json:"imageurl"`
	MetaInfo struct {
		Description string `json:"description"`
		Price 		string `json:"price"`
		Seller 		string `json:"seller"`
		Brand 		string `json:"brand"`
	} `json:"metainfo"`
}

type ScrapeResponse struct {
	MIMEType string `json:"mimetype"`
	Response Product `json:"response"`
}

func (a *App) scrape(purl string) (Product, error) {
	var pinfo ScrapeResponse
	bodyparams := map[string]string{}
	bodyparams["producturl"] = purl
	log.Println("Posting request to", a.Conf.ScrapeEndpoint)
	_, result, err := util.SendPostRequest(a.Conf.ScrapeEndpoint, bodyparams)
	if err != nil {
		return pinfo.Response, err
	}
	err = json.Unmarshal([]byte(result), &pinfo)
	return pinfo.Response, err
}

func (a *App) storeMetaInfo(pid int64, purlid int64, product Product) (int64, error) {
	metainfoid := util.XXHash(product.MetaInfo.Description + product.Retailer + product.MetaInfo.Price + product.MetaInfo.Seller) + purlid 
	err := a.CasCursor.AddProductMetaInfo(pid, purlid, metainfoid, product.MetaInfo.Description, product.Retailer, product.MetaInfo.Price, product.MetaInfo.Seller)
	return metainfoid, err
}

func (a *App) storeImageInfo(purlid int64, iurls []string) ([]int64, error) {
	idList := []int64{}
	for _, iurl := range iurls {
		iurlid := util.XXHash(iurl)
		err := a.CasCursor.AddProductImageInfo(purlid, iurl, iurlid)
		if err != nil {
		 	return idList, err
		}
		idList = append(idList, iurlid)
	}
	return idList, nil
}

func (a *App) storeProductInfo(pid int64, purlid int64, metainfoid int64, idlist []int64) error {
	err := a.CasCursor.AddProductInfo(pid, purlid, metainfoid, idlist)
	return err
}

func (a *App) fetchAndStoreInformation(pid int64, purlid int64, purl string, pinfo Product) error {
	metainfoid, err := a.storeMetaInfo(pid, purlid, pinfo)
	if err != nil {
		return err
	}
	idList, err := a.storeImageInfo(purlid, pinfo.ImageURL)
	if err != nil {
		return err
	}
	err = a.storeProductInfo(pid, purlid, metainfoid, idList)
	return err
}

func (a *App) updateMetrics(pid int64, purlid int64, pname string, totaltry int64, failcount int64, isFailed bool) error {
	isExpired := 0	
	if !isFailed {
		return a.CasCursor.UpdateProductMetrics(pid, purlid, pname, isExpired, totaltry+1, failcount)
	}
	if ((failcount +1) / (totaltry + 1)) * 100 > int64(a.Conf.ExpiryThreshold) {
		isExpired = 1
	}
	return a.CasCursor.UpdateProductMetrics(pid, purlid, pname, isExpired, totaltry+1, failcount+1)
}

func (a *App) scrapeAndCache() error {
	productPageInfo, err := a.CasCursor.GetProductPageInfo()
	if err != nil {
		log.Println(err)
		return err 
	}
	for _, page := range productPageInfo {
		isFailed := false
		pid, pname, purlid, purl  := page.PID, page.PName, page.PURLID, page.PURL 
		pinfo, err := a.scrape(purl)
		if purl == "" {
			continue
		}
		if err != nil {
			isFailed = true
			a.updateMetrics(pid, purlid, pname, page.TotalTries, page.TotalMisses, isFailed)
			return err 
		} else {
			a.updateMetrics(pid, purlid, pname, page.TotalTries, page.TotalMisses, isFailed)
		}
		err = a.fetchAndStoreInformation(pid, purlid, purl, pinfo)
		if err != nil {
			return err
		}
	}
	return nil
}


func (a *App) startScrapeJob() {
	func() {
		ticker := time.NewTicker(time.Duration(a.Conf.ScrapeInterval) * time.Second)
		log.Println("Starting scrape job")
		for {
			select {
			case <-ticker.C:
				a.scrapeAndCache()
			case _, ok := <-a.Quit:
				if !ok {
					ticker.Stop()
					log.Println("Scrape job ticker stopped!")
					return
				}
			}
		}
	}()
}