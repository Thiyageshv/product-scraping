package cassandra_test

import (
	"testing"
	"log"
	cas "product-scraping/lib_cassandra"
	util "product-scraping/lib_utilities"
)

var pid = int64(01)
var purlid = int64(123)
var description = "Don't let spills and messes get in your way. Lock in confidence with Bounty, the Quicker Picker Upper*. This pack contains Bounty white Select-A-Size paper towels that are 2X more absorbent* and strong when wet, so you can get the job done quickly. *vs. leading ordinary brand"
var retailer = "Rakuten"
var price = "0.01$"
var seller = "bounty"

func Test_add(t *testing.T) {
	config := cas.LocalCasaConfig()
	casdb, err := cas.InitializeCasaConn(config)
	if err != nil {
		log.Println(err)
		return
	}
	metainfoid := util.XXHash(description + retailer + price + seller) + purlid
	err = casdb.AddProductMetaInfo(pid, purlid, metainfoid, description, retailer, price, seller)
	if err != nil {
		log.Println(err)
	}
}

func Test_basic(t *testing.T) {
	config := cas.LocalCasaConfig()
	casdb, err := cas.InitializeCasaConn(config)
	if err != nil {
		log.Println(err)
		return
	} 
	pinfo, err := casdb.GetProductMetaInfo(purlid, pid)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pinfo)
}