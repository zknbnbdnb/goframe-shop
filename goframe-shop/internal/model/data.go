package model

type DataHeadInput struct{}

type DataHeadOutput struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单数"`
	DAU             int `json:"dau" dc:"今日访客数"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}

type DataEChartsInput struct{}

type DataEChartsOutput struct {
	OrderTotal           []int
	SalePriceTotal       []int
	ConsumptionPerPerson []int
	NewOrder             []int
}

type TodayTotal struct {
	Today string `json:"today"`
	Total int    `json:"total"`
}
