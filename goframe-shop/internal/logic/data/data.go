package data

import (
	"context"
	"fmt"
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

func (s *sData) DataECharts(ctx context.Context) (out *model.DataEChartsOutput, err error) {
	return &model.DataEChartsOutput{
		OrderTotal:           OrderTotal(ctx),
		SalePriceTotal:       SalePriceTotalRecentDays(ctx),
		ConsumptionPerPerson: OrderTotal(ctx),
		NewOrder:             OrderTotal(ctx),
	}, nil
}

func OrderTotal(ctx context.Context) (counts []int) {
	counts = []int{0, 0, 0, 0, 0, 0, 0}
	recent7Dates := utility.GetRecent7Date()
	var TodayTotals []model.TodayTotal
	err := dao.OrderInfo.Ctx(ctx).
		Where(dao.OrderInfo.Columns().CreatedAt+" >= ", utility.GetBefore7Date()).
		Fields("count(*) total,date_format(created_at, '%Y-%m-%d') today").
		Group("today").
		Scan(&TodayTotals)
	if err != nil {
		return nil
	}
	fmt.Printf("result:%v\n", TodayTotals)
	for i, data := range recent7Dates {
		for _, total := range TodayTotals {
			if data == total.Today {
				counts[i] = total.Total
			}
		}
	}
	if err != nil {
		return counts
	}
	return
}

func SalePriceTotalRecentDays(ctx context.Context) (totals []int) {
	totals = []int{0, 0, 0, 0, 0, 0, 0}
	recent7Dates := utility.GetRecent7Date()
	var TodayTotals []model.TodayTotal
	err := dao.OrderGoodsInfo.Ctx(ctx).
		Where(dao.OrderInfo.Columns().CreatedAt+" >= ", utility.GetBefore7Date()).
		Fields("sum(actual_price) total,date_format(created_at, '%Y-%m-%d') today").
		Group("today").
		Scan(&TodayTotals)
	if err != nil {
		return nil
	}
	fmt.Printf("result:%v\n", TodayTotals)
	for i, data := range recent7Dates {
		for _, total := range TodayTotals {
			if data == total.Today {
				totals[i] = total.Total
			}
		}
	}
	if err != nil {
		return totals
	}
	return
}
