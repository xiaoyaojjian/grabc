package permission

import (
	. "grabc/views/layout"
)

type Index struct {
	BaseTemplate
}

func (this *Index) Html() string {
	html := `
<div class="box box-primary">
        <div class="box-body">
            <a class="btn btn-primary" href="/permission/add" role="button">新增权限</a>
        </div>
    </div>
    <div class="box box-info">
        <div class="box-body">
            <table class="table table-bordered table-striped">
                <thead>
                    <tr>
                        <td>ID</td>
                        <td>权限名称</td>
                        <td>权限描述</td>
                        <td class="row_operate">操作</td>
                    </tr>
                </thead>
                <tbody>
                {{range $index,$permission:=.permissions}}
                    <tr>
                        <td>{{$permission.Id}}</td>
                        <td>{{$permission.Name}}</td>
                        <td>{{$permission.Description}}</td>
                        <td>
                            <a href="/permission/get?id={{$permission.Id}}" title="查看">
                                <span class="glyphicon glyphicon-eye-open"></span>
                            </a>
                            <a href="/permission/put?id={{$permission.Id}}" title="更新">
                                <span class="glyphicon glyphicon-pencil"></span>
                            </a>
                            <a href="/permission/delete?id={{$permission.Id}}" title="删除" data-confirm="您确定要删除此项吗？">
                                <span class="glyphicon glyphicon-trash"></span>
                            </a>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
            {{pagination}}
        </div>
    </div>
`
	return this.DealHtml(html)
}
