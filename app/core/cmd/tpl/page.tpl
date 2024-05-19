func (l *{{ .Name }}PageLogic) {{ .Name }}Page(req *types.{{ .Name }}PageRequest) (resp *types.{{ .Name }}PageResponse, err error) {
    where := " 1 "
    sys{{ .Name }}Page, err := l.svcCtx.Sys{{ .Name }}Model.FindPageByWhere(l.ctx, where, req.Page, req.Limit)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var item types.{{ .Name }}
	{{ .Name }}Page := make([]*types.{{ .Name }}, 0)
	for _, v := range sys{{ .Name }}Page {
		err := copier.Copy(&item, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		{{ .Name }}Page = append({{ .Name }}Page, &item)
	}

	total, err := l.svcCtx.Sys{{ .Name }}Model.FindPageByWhereCount(l.ctx, where)
    if err != nil {
         return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
    }

    pagination := types.Pagination{
         Page:  req.Page,
         Limit: req.Limit,
         Total: total,
    }

	return &types.{{ .Name }}PageResponse{
		List: {{ .Name }}Page,
   		Pagination: pagination,
	}, nil
}