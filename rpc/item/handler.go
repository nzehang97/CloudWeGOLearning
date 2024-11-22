package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"my_shop/kitex_gen/my/shop/item"
	"my_shop/kitex_gen/my/shop/stock"
	"my_shop/kitex_gen/my/shop/stock/stockservice"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct {
	stockCli stockservice.Client
}

func NewStockClient(addr string) (stockservice.Client, error) {
	return stockservice.NewClient("my.shop.stock", client.WithHostPorts(addr))
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemReq) (resp *item.GetItemResp, err error) {
	// TODO: Your code here...
	resp = item.NewGetItemResp()
	resp.Item = item.NewItem()
	resp.Item.Id = req.Id
	resp.Item.Title = "my shop"
	resp.Item.Description = "my kitex learning on shop"

	stockReq := stock.NewGetItemStockReq()
	stockReq.ItemId = req.GetId()
	stockResp, err := s.stockCli.GetItemStock(context.Background(), stockReq)
	if err != nil {
		log.Println(err)
		stockResp.Stock = 0
	}
	resp.Item.Stock = stockResp.GetStock()
	return
}
