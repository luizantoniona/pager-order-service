package route

import (
	"net/http"
	"pager-order-service/handler"
)

func SetupRoutes() {
	http.HandleFunc("/orders", handler.HandleOrders)
	http.HandleFunc("/pagers", handler.HandlePagers)
}
