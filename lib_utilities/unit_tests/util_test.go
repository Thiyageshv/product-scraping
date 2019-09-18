package util_test

import (
	"testing"
	"log"
	util "product-scraping/lib_utilities"
)

func Test_basic(t * testing.T) {
	var id = 'RGVsaWNpb3VzIENob2NvbGF0ZSBDaGlwIEJpc2N1aXQgQ29va2llcyBmb3IgR2lmdA'
	decode, err := util.B64Decode(id)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(decode)
}