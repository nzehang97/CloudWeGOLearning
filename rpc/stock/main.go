package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	stock "my_shop/kitex_gen/my/shop/stock/stockservice"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8890")
	svr := stock.NewServer(new(StockServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
