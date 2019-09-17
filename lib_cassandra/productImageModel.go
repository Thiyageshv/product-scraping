package cassandra

import (
	"time"
	util "product-scraping/lib_utilities"
)



type ProductImageInfo struct {
	PURLID     		int64 	  `json:"iurlgroupid"`
	IURLID 			int64 	  `json:"iurlid"`
	IURL 			string 	  `json:"iurl"`
	CreatedOn 		time.Time `json:"createdon"`
	ModifiedOn      time.Time `json:"modifiedon"`
}


func (c * CasDb) AddProductImageInfo(groupid int64, iurl string) error {
	iurlid := util.XXHash(iurl)
	err := c.Session.Query(prepareQuery(addProductImage, groupid, iurlid, iurl)).Exec()
	return err
}

func (c *CasDb) GetProductImageInfo(iurlid int64, purlid int64) ([]ProductImageInfo, error) {
	var results []ProductImageInfo
	var pinfo ProductImageInfo
	iter := c.Session.Query(prepareQuery(getProductImageInfo, iurlid, purlid)).Iter()
	for iter.Scan(&pinfo.PURLID, &pinfo.IURLID, &pinfo.IURL, &pinfo.CreatedOn, &pinfo.ModifiedOn) {
		results = append(results, pinfo)
	}
	err := iter.Close()
	return results, err
} 