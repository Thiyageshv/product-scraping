package cassandra_test

import (
	"testing"
	"log"
	cas "product-scraping/lib_cassandra"
)

var pid = int64(1)
var purlid = int64(123)
var metainfoid = int64(345)
var group= []int64{568}
var additionalID = int64(4242)

func Test_add(t *testing.T) {
	config := cas.LocalCasaConfig()
	casdb, err := cas.InitializeCasaConn(config)
	if err != nil {
		log.Println(err)
		return
	}
	err = casdb.AddProductInfo(pid, purlid, metainfoid, group)
	if err != nil {
		log.Println(err)
		return
	}
	err = casdb.AddProductImageID(pid, purlid, additionalID)
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
	pinfo, err := casdb.GetProductInfo(purlid, pid)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pinfo)
}