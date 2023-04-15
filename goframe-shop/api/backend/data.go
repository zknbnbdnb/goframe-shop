package backend

import "github.com/gogf/gf/v2/frame/g"

type DataHeadReq struct {
	g.Meta `path:"/data/head" method:"get" tags:"数据data" desc:"数据大屏的头部信息"`
}

type DataHeadRes struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单数"`
	DAU             int `json:"dau" dc:"今日访客数"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}

type DataEChartsReq struct {
	g.Meta `path:"/data/echarts" method:"get" tags:"数据data" desc:"数据大屏的echarts"`
}

type DataEChartsRes struct {
	OrderTotal           []int `json:"order_total" dc:"订单量"`
	SalePriceTotal       []int `json:"sale_price_total" dc:"销售价格"`
	ConsumptionPerPerson []int `json:"consumption_per_person" dc:"人均消费"`
	NewOrder             []int `json:"new_order" dc:"新增订单"`
}
