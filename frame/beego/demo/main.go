package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func main() {
	fmt.Println("数据读取开始")
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	var maps []orm.Params
	num, err := o.Raw("SELECT operation, count(*) as countall , count(distinct(openid)) as count FROM dx_trouble_records WHERE date(DATE_FORMAT(FROM_UNIXTIME(createtime),'%Y-%m-%d')) =(select date_sub(curdate(),interval 1 day)) group by operation order by type asc").Values(&maps)

}
