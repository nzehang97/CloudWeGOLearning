package main

import (
	"log"
	stock "my_shop/kitex_gen/my/shop/stock/stockservice"
)

func main() {
	svr := stock.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
