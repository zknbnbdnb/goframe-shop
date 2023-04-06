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
