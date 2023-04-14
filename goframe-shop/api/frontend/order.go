package frontend

import "github.com/gogf/gf/v2/frame/g"

type OrderAddReq struct {
	g.Meta `path:"/order/add" method:"post" summary:"添加订单" tags:"订单"`
	// 主订单维度
	Price            int    `json:"price"            description:"订单金额 单位分"`
	CouponPrice      int    `json:"coupon_price"      description:"优惠券金额 单位分"`
	ActualPrice      int    `json:"actual_price"      description:"实际支付金额 单位分"`
	ConsigneeName    string `json:"consignee_name"    description:"收货人姓名"`
	ConsigneePhone   string `json:"consignee_phone"   description:"收货人手机号"`
	ConsigneeAddress string `json:"consignee_address" description:"收货人详细地址"`
	Remark           string `json:"remark"           description:"备注"`
	// 商品订单维度
	OrderAddGoodsInfo []*OrderAddGoodsInfo `json:"order_add_goods_info"`
}

type OrderAddRes struct {
	Id uint
}

type OrderAddGoodsInfo struct {
	GoodsId        int    `json:"goods_id"        description:"商品id"`
	GoodsOptionsId int    `json:"goods_options_id" description:"商品规格id sku id"`
	Count          int    `json:"count"          description:"商品数量"`
	Remark         string `json:"remark"         description:"备注"`
	Price          int    `json:"price"          description:"订单金额 单位分"`
	CouponPrice    int    `json:"coupon_price"    description:"优惠券金额 单位分"`
	ActualPrice    int    `json:"actual_price"    description:"实际支付金额 单位分"`
}
