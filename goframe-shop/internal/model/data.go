package model

type DataHeadOutput struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单数"`
	DAU             int `json:"dau" dc:"今日访客数"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}
