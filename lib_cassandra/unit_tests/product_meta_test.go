package cassandra_test

import (
	"testing"
	"log"
	cas "product-scraping/lib_cassandra"
)

var pid = int64(01)
var purlid = int64(123)
var description = "Lorum Epsum"
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
	err = casdb.AddProductMetaInfo(pid, purlid, description, retailer, price, seller)
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