
func (l *{{ .Name }}ListLogic) {{ .Name }}List(req *types.{{ .Name }}ListReq) (resp *types.{{ .Name }}ListResp, err error) {
    where := " 1 "
    feat{{ .Name }}List, err := l.svcCtx.Feat{{ .Name }}Model.FindAllByWhere(l.ctx, where)
	if err != nil {
		return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
	}

	var item types.{{ .Name }}
	{{ .Name }}List := make([]*types.{{ .Name }}, 0)
	for _, v := range feat{{ .Name }}List {
		err := copier.Copy(&item, &v)
		if err != nil {
			return nil, errorx2.NewSystemError(errorx2.ServerErrorCode, err.Error())
		}
		{{ .Name }}List = append({{ .Name }}List, &item)
	}

	return &types.{{ .Name }}ListResp{
		List: {{ .Name }}List,
	}, nil
}