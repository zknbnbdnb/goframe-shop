package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/controller"
	"goframe-shop/internal/controller/backend"
	"goframe-shop/internal/controller/frontend"
	"goframe-shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			// 后台路由
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse 不使用官方的中间件
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由
				group.Bind(
					controller.Admin.Create, // 管理员 创建
					controller.Login,        // 登录
				)
				// 需要登录的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) // this middleware is for jwt
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Data,         // 数据大屏
						controller.Role,         // 角色管理
						controller.Permission,   // 权限管理
						controller.Admin.Info,   // 查询当前管理员信息
						controller.Admin.Update, // 管理员 更新
						controller.Admin.List,   // 管理员 列表
						controller.Admin.Delete, // 管理员 删除
						controller.Rotation,     // 轮播图
						controller.Position,     // 手工位
						controller.File,         // 从0到1实现文件入库
						controller.Upload,       // 实现可跨项目使用文件上传的工具类
						controller.Category,     // 商品分类管理
						controller.Coupon,       // 优惠券管理
						controller.UserCoupon,   // 用户优惠券管理
						controller.Goods,        // 商品管理
						controller.GoodsOptions, // 商品规模管理
						controller.Address,      // 地址管理
						//这么写是为了避免前后端重复注册相同的路由和方法
						controller.Order.List,   //订单列表
						controller.Order.Detail, //订单详情
						backend.Article,         // 文章管理&CMS
					)
				})
			})

			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			// 前台路由
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS, // 跨域请求中间件
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由
				group.Bind(
					controller.User.Register, // 用户注册
					controller.Goods,         // 商品管理
				)
				// 需要登录的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.User.Info,           // 获取用户信息
						controller.User.UpdatePassword, // 修改用户密码
						controller.Collection,          // 收藏管理
						controller.Praise,              // 点赞管理
						controller.Comment,             // 评论管理
						controller.Cart,                // 购物车管理
						controller.Order.Add,           // 添加订单
						controller.OrderGoodsComments,  // 订单商品评论
						frontend.Article,               // 文章
						frontend.Refund,                // 售后
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
