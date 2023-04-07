package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/controller"
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
			// 不需要登录的路由组绑定
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse 不使用官方的中间件
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Rotation,     // 轮播图
					controller.Position,     // 手工位
					controller.Admin.Create, // 管理员 创建
					controller.Admin.Update, // 管理员 更新
					controller.Admin.List,   // 管理员 列表
					controller.Admin.Delete, // 管理员 删除
					controller.Login,        // 登录
					controller.Data,         // 数据大屏
					controller.Role,         // 角色管理
					controller.Permission,   // 权限管理
				)
				// 需要登录的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) // this middleware is for jwt
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/admin/info": controller.Admin.Info, // 管理员 详情,因为就他需要token,所以单独写
					})
					group.Bind(
						controller.File,       // 从0到1实现文件入库
						controller.Upload,     // 实现可跨项目使用文件上传的工具类
						controller.Category,   // 商品分类管理
						controller.Coupon,     // 优惠券管理
						controller.UserCoupon, // 用户优惠券管理
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
