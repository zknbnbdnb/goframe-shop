package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

// s指的是simple的意思，这里的sData是一个简单的数据结构体，用于演示。
type sData struct {
}

// 注册服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{ // todo 暂时写死
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(500),
	}, nil
}

// 查询订单总数
func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.
		Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
