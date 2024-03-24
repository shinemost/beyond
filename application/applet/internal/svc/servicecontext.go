package svc

import (
	"github.com/shinemost/beyond/application/applet/internal/config"
	"github.com/shinemost/beyond/application/user/rpc/user"
	"github.com/shinemost/beyond/pkg/interceptors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器 处理gprc返回的status 如果有异常转换code包装异常信息并返回
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
		BizRedis: redis.MustNewRedis(c.BizRedis), //redis.New方法已经过时，请使用MustNewRedis 或者NewRedis 来替代
	}
}
