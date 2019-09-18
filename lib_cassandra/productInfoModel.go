package cassandra

import (
	"time"
	// util "product-scraping/lib_utilities"
)


 
type ProductInfo struct {
	PID 			int64 	  `json:"pid"`
	PURLID 			int64 	  `json:"purlid"`
	MetaInfoID		int64 	  `json:"metainfoid"`
	IURLGroup     	[]int64   `json:"iurlgroup"`
	CreatedOn 		time.Time `json:"createdon"`
	ModifiedOn      time.Time `json:"modifiedon"`
}

type SimpleProductInfo struct {
	PID int64 `json:"pid"`
	PURLID int64 `json:"purlid"`
	IURL string `json:"iurl"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price string `json:"price"`
	Retailer string `json:"retailer"`
	Seller string `json:"seller"`
	CreatedOn string `json:"createdon"`
	ModifiedOn string `json:"modifiedon"`
}


func (c *CasDb) GetProductInfo(purlid int64, pid int64) (ProductInfo, error) {
	var pinfo ProductInfo
	err := c.Session.Query(prepareQuery(getProductInfo, purlid, pid)).Scan(&pinfo.PID, &pinfo.PURLID, &pinfo.MetaInfoID, &pinfo.IURLGroup, &pinfo.CreatedOn, &pinfo.ModifiedOn)
	return pinfo, err
}

func (c * CasDb) AddProductInfo(pid int64, purlid int64, metainfoid int64, group []int64) error {
	err := c.Session.Query(prepareQuery(addProductInfo, pid, purlid, metainfoid, group)).Exec()
	return err
}

func (c *CasDb) AddProductImageID(pid int64, purlid int64, iurlid int64) error {
	err := c.Session.Query(prepareQuery(addImageID, iurlid, purlid, pid)).Exec()
	return err
}