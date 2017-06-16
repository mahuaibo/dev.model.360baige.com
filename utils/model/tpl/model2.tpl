{{$exportModelName := .ModelName | FirstCharUpper}}

package {{.PackageName}}

type {{$exportModelName}} struct {
	{{range .TableSchema}} {{.COLUMN_NAME | ExportColumn}} {{.DATA_TYPE | TypeConvert}} {{.COLUMN_NAME | Tags}} // {{.COLUMN_COMMENT}}
	{{end}}
}

/**
 * ******* {{$exportModelName}}
 */
 // 新增
func (*{{$exportModelName}}) Add(args *{{$exportModelName}}, reply *{{$exportModelName}}) error {
	o := orm.NewOrm()
	o.Using("{{$exportModelName}}")
	id, err := o.Insert(args)
	if err == nil {
		{{range .TableSchema}}{{$column_name := .COLUMN_NAME | ExportColumn}}{{ if eq $column_name "Id" }}reply.{{$column_name}} = id{{ else }}reply.{{$column_name}} = args.{{$column_name}}{{ end }}
		{{end}}
	}
	return err
}

// 查询 by Id
func (*{{$exportModelName}}) FindById(args *{{$exportModelName}}, reply *{{$exportModelName}}) error {
	o := orm.NewOrm()
	o.Using("{{$exportModelName}}")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*{{$exportModelName}}) UpdateById(args *{{$exportModelName}}, reply *{{$exportModelName}}) error {
	o := orm.NewOrm()
	o.Using("{{$exportModelName}}")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
