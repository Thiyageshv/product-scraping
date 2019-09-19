package cassandra_test

import (
	"testing"
	"log"
	cas "product-scraping/lib_cassandra"
	util "product-scraping/lib_utilities"
)

var purlid = int64(123)
var url = "https://www.amaon.com/Samsung-QN65Q6FN-FLAT-QLED-Smart/dp/B0hfth9V1MSQ1"
var urlid = util.XXHash(url)

func Test_add(t *testing.T) {
	config := cas.LocalCasaConfig()
	casdb, err := cas.InitializeCasaConn(config)
	if err != nil {
		log.Println(err)
		return
	}
	err = casdb.AddProductImageInfo(purlid, url)
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
	pinfo, err := casdb.GetProductImageInfo(urlid, purlid)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pinfo)
}