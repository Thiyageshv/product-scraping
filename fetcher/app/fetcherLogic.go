package app 

import (
	"log"
)

type ProductInfo struct {
	Name string `json:"name"`
} 

func (a *App) getInfo() (ProductInfo, error) {
	log.Println("Api entry")
	return ProductInfo{Name: "Thiyagesh"}, nil
}