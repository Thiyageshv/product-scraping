package cassandra

import (
	"time"
	util "product-scraping/lib_utilities"
)



type ProductMetaInfo struct {
	PID 			int64 	  `json:"pid"`
	PURLID 			int64 	  `json:"purlid"`
	MetaInfoID		int64 	  `json:"metainfoid"`
	Description     string 	  `json:"description"`
	Retailer        string 	  `json:"retailer"`
	Price 			string   `json:"price"`
	Seller 			string 	  `json:"seller"`
	CreatedOn 		time.Time `json:"createdon"`
	ModifiedOn      time.Time `json:"modifiedon"`
}

func (c *CasDb) GetProductMetaInfo(purlid int64, pid int64) (ProductMetaInfo, error) {
	var pinfo ProductMetaInfo
	err := c.Session.Query(prepareQuery(getProductMetaInfo, purlid, pid)).Scan(&pinfo.PID, &pinfo.PURLID, &pinfo.MetaInfoID, &pinfo.Description, &pinfo.Retailer, &pinfo.Price, &pinfo.Seller, &pinfo.CreatedOn, &pinfo.ModifiedOn)
	return pinfo, err
}

func (c * CasDb) AddProductMetaInfo(pid int64, purlid int64, description string, retailer string, price string, seller string) error {
	metainfoid := util.XXHash(description + retailer + price + seller) + purlid 
	err := c.Session.Query(prepareQuery(addProductMetaInfo, pid, purlid, metainfoid, description, retailer, price, seller)).Exec()
	return err
}
