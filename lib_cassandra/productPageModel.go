package cassandra

import (
	"time"
)


type ProductPageInfo struct {
	PID 			int64 		`json:"pid"`
	PURLID 			int64 		`json:"purlid"`
	PName 			string 		`json:"pname"`
	PURL 			string 		`json:"purl"`
	IsExpired 		int 		`json:"isExpired"`
	TotalTries 		int64 		`json:"totaltries"`
	TotalMisses 	int64 		`json:"totalmisses"`
	CreatedOn 		time.Time 	`json:"createdon"`
	ModifiedOn 		time.Time 	`json:"modifiedon"`
	ModifiedBy 		string 		`json:"modifiedby"`
}

func (c *CasDb) GetProductPagesBasic() ([]ProductPageInfo, error) {
	var results []ProductPageInfo
	var pinfo ProductPageInfo
	iter := c.Session.Query(getPageBasicInfo).Iter()
	for iter.Scan(&pinfo.PID, &pinfo.PURLID, &pinfo.PName, &pinfo.PURL) {
		results = append(results, pinfo)
	}
	err := iter.Close()
	return results, err
}


func (c *CasDb) GetProductPagesMetrics() ([]ProductPageInfo, error) {
	var results []ProductPageInfo
	var pinfo ProductPageInfo
	iter := c.Session.Query(getPageMetricInfo).Iter()
	for iter.Scan(&pinfo.PID, &pinfo.PURLID, &pinfo.PName, &pinfo.IsExpired, &pinfo.TotalTries, &pinfo.TotalMisses, &pinfo.CreatedOn, &pinfo.ModifiedOn, &pinfo.ModifiedBy) {
		results = append(results, pinfo)
	}
	err := iter.Close()
	return results, err
}

func (c *CasDb) UpdateProductMetrics(pid int64, purlid int64, isExpired int, totaltries int64, totalmisses int64, pname string) error {
	err := c.Session.Query(prepareQuery(updateMetrics, pid, purlid, isExpired, totaltries, totalmisses, pname)).Exec()
	return err
}