package service

type OrderService struct {
	Commands interface{}
	Queries  interface{}
}

func InitOrderService(c interface{}, q interface{}) *OrderService {
	return &OrderService{Commands: c, Queries: q}
}
