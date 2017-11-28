package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/robfig/cron"
)

//,事件操作,事件总数,唯一访客事件数,事件价值,平均价值
//,总计,28245,18428,0,0
//1,报障首页,11597,7224,0,0
//2,手机报障,203,139,0,0
//3,固话报障,1145,651,0,0
//4,宽带报障,3585,2017,0,0
//5,IPTV报障,700,443,0,0
//6,进度查询,2216,902,0,0
//7,故障小知识,5968,5308,0,0
//8,一键报障,1846,1091,0,0
//9,一键修复,985,653,0,0
func dd() {
	fmt.Println("数据读取开始")
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	var maps []orm.Params
	num, err := o.Raw("SELECT operation, count(*) as countall , count(distinct(openid)) as count FROM dx_trouble_records WHERE date(DATE_FORMAT(FROM_UNIXTIME(createtime),'%Y-%m-%d')) =(select date_sub(curdate(),interval 1 day)) group by operation order by type asc").Values(&maps)

	if err != nil {
		fmt.Println(err.Error())
		//TODO 发送错误提示
	}

	txtStr := ",事件操作,事件总数,唯一访客事件数,事件价值,平均价值\n"

	fmt.Println("查询到结果行数：" + string(num))
	sum1 := 0
	sum2 := 0
	// 求和
	for _, tt := range maps {
		if tt["operation"] != nil {
			var s1 int
			var s2 int
			s1, _ = strconv.Atoi(tt["countall"].(string))
			s2, _ = strconv.Atoi(tt["count"].(string))
			sum1 += s1
			sum2 += s2
		}
	}

	txtStr = txtStr + ",总计," + strconv.Itoa(sum1) + "," + strconv.Itoa(sum2) + ",0,0\n"

	i := 1

	for _, term := range maps {
		if term["operation"] != nil {
			istr := strconv.Itoa(i) + "," + term["operation"].(string) + "," + term["countall"].(string) + "," + term["count"].(string) + ",0,0\n"
			txtStr = txtStr + istr
			i = i + 1
		}
	}
	fmt.Println(txtStr)
	writeFile(txtStr)
	defer fmt.Println("数据写入结束")
}

func writeFile(str string) {
	ymd := time.Now().Format("20060102")
	dstFile, err := os.Create("./loginfo/KDBZ_" + ymd + ".csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(str + "\n")
}

func StartStatistics() {
	i := 0
	c := cron.New()
	spec := "0 30 1 * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
		go dd()
	})
	c.Start()
}