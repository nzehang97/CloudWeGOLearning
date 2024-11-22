package main

import (
	"log"
	item "my_shop/kitex_gen/my/shop/item/itemservice"
)

func main() {
	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
