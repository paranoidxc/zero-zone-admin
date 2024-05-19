func (l *{{ .Name }}ListLogic) {{ .Name }}List(req *types.{{ .Name }}ListRequest) (resp *types.{{ .Name }}ListResponse, err error) {
    where := " 1 "
    sys{{ .Name }}List, err := l.svcCtx.Sys{{ .Name }}Model.FindAllByWhere(l.ctx, where)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var item types.{{ .Name }}
	{{ .Name }}List := make([]*types.{{ .Name }}, 0)
	for _, v := range sys{{ .Name }}List {
		err := copier.Copy(&item, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		{{ .Name }}List = append({{ .Name }}List, &item)
	}

	return &types.{{ .Name }}ListResponse{
		List: {{ .Name }}List,
	}, nil
}