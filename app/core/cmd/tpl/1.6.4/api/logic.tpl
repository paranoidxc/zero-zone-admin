package {{.pkgName}}

import (
	{{.imports}}
   	"github.com/jinzhu/copier"
   	"zero-zone/app/core/model"
   	errorx2 "zero-zone/app/pkg/errorx"
   	"github.com/zeromicro/go-zero/core/logx"
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

{{if .hasDoc}}{{.doc}}{{end}}
func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
