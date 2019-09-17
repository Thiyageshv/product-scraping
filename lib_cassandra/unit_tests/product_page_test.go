package cassandra_test

import (
	"testing"
	"log"
	cas "product-scraping/lib_cassandra"
)

var pid = int64(1790395674)
var purlid = int64(4279907335)
var isexpired = 0
var totaltries = int64(5)
var totalmisses = int64(3)
var pname = "tv"

func Test_update(t *testing.T) {
	config := cas.LocalCasaConfig()
	casdb, err := cas.InitializeCasaConn(config)
	if err != nil {
		log.Println(err)
		return
	} 
	err = casdb.UpdateProductMetrics(pid, purlid, isexpired, totaltries, totalmisses, createdon)
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
	pinfo, err := casdb.GetProductPagesBasic()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pinfo)
}