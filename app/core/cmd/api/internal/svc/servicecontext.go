package svc

import (
	"zero-zone/app/core/cmd/api/internal/config"
	"zero-zone/app/core/cmd/api/internal/middleware"
	"zero-zone/app/core/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config             config.Config
	Redis              *redis.Redis
	PermMenuAuth       rest.Middleware
	SysUserModel       model.SysUserModel
	SysPermMenuModel   model.SysPermMenuModel
	SysRoleModel       model.SysRoleModel
	SysDeptModel       model.SysDeptModel
	SysJobModel        model.SysJobModel
	SysProfessionModel model.SysProfessionModel
	SysDictionaryModel model.SysDictionaryModel
	SysLogModel        model.SysLogModel
	FeatTdFirmModel    model.TdFirmModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:             c,
		Redis:              redisClient,
		PermMenuAuth:       middleware.NewPermMenuAuthMiddleware(redisClient).Handle,
		SysUserModel:       model.NewSysUserModel(mysqlConn, c.Cache),
		SysPermMenuModel:   model.NewSysPermMenuModel(mysqlConn, c.Cache),
		SysRoleModel:       model.NewSysRoleModel(mysqlConn, c.Cache),
		SysDeptModel:       model.NewSysDeptModel(mysqlConn, c.Cache),
		SysJobModel:        model.NewSysJobModel(mysqlConn, c.Cache),
		SysProfessionModel: model.NewSysProfessionModel(mysqlConn, c.Cache),
		SysDictionaryModel: model.NewSysDictionaryModel(mysqlConn, c.Cache),
		SysLogModel:        model.NewSysLogModel(mysqlConn, c.Cache),
		FeatTdFirmModel:    model.NewTdFirmModel(mysqlConn, c.Cache),
	}
}
