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
	pinfo.Description, err = util.B64Decode(pinfo.Description)
	pinfo.Retailer, err = util.B64Decode(pinfo.Retailer)
	pinfo.Seller, err = util.B64Decode(pinfo.Seller)
	return pinfo, err
}

func (c * CasDb) AddProductMetaInfo(pid int64, purlid int64, metainfoid int64, description string, retailer string, price string, seller string) error {
	description = util.B64Encode(description)
	retailer = util.B64Encode(retailer)
	seller = util.B64Encode(seller)
	err := c.Session.Query(prepareQuery(addProductMetaInfo, pid, purlid, metainfoid, description, retailer, price, seller)).Exec()
	return err
}
