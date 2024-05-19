package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TdFirmModel = (*customTdFirmModel)(nil)

type (
	// TdFirmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTdFirmModel.
	TdFirmModel interface {
		tdFirmModel
	}

	customTdFirmModel struct {
		*defaultTdFirmModel
	}
)

// NewTdFirmModel returns a model for the database table.
func NewTdFirmModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TdFirmModel {
	return &customTdFirmModel{
		defaultTdFirmModel: newTdFirmModel(conn, c, opts...),
	}
}
