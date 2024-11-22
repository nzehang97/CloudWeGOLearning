package main

import (
	"log"
	item "my_shop/kitex_gen/my/shop/item/itemservice"
)

func main() {
	itemServiceImpl := new(ItemServiceImpl)
	stockCli, err := NewStockClient("localhost:8890")
	if err != nil {
		log.Fatal(err)
	}
	itemServiceImpl.stockCli = stockCli

	svr := item.NewServer(itemServiceImpl)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
