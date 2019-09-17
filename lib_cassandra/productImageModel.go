package cassandra

import (
	"time"
)



type ProductImageInfo struct {
	PURLID     		int64 	  `json:"iurlgroupid"`
	IURLID 			int64 	  `json:"iurlid"`
	IURL 			string 	  `json:"iurl"`
	CreatedOn 		time.Time `json:"createdon"`
	ModifiedOn      time.Time `json:"modifiedon"`
}


func (c * CasDb) AddProductImageInfo(groupid int64, iurl string, iurlid int64) error {
	err := c.Session.Query(prepareQuery(addProductImage, groupid, iurlid, iurl)).Exec()
	return err
}

func (c *CasDb) GetProductImageInfo(iurlid int64, purlid int64) (ProductImageInfo, error) {
	var pinfo ProductImageInfo
	err := c.Session.Query(prepareQuery(getProductImageInfo, iurlid, purlid)).Scan(&pinfo.PURLID, &pinfo.IURLID, &pinfo.IURL, &pinfo.CreatedOn, &pinfo.ModifiedOn)
	return pinfo, err
} 


func (c *CasDb) GetProductImagesInfo(purlid int64) ([]ProductImageInfo, error) {
	var results []ProductImageInfo
	var pinfo ProductImageInfo
	iter := c.Session.Query(prepareQuery(getProductImagesInfo, purlid)).Iter()
	for iter.Scan(&pinfo.PURLID, &pinfo.IURLID, &pinfo.IURL, &pinfo.CreatedOn, &pinfo.ModifiedOn) {
		results = append(results, pinfo)
	}
	err := iter.Close()
	return results, err
} 