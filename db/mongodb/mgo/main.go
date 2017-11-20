package main

//参考 http://blog.csdn.net/wangshubo1989/article/details/75105397?locationNum=4&fps=1
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const URL = "127.0.0.1:27017" //mongodb的地址

func main() {
	session, err := mgo.Dial(URL) //连接服务器
	if err != nil {
		panic(err)
	}

	c := session.DB("ChatRoom").C("account") //选择ChatRoom库的account表

	c.Insert(map[string]interface{}{"id": 7, "name": "tongjh", "age": 25}) //增

	objid := bson.ObjectIdHex("55b97a2e16bc6197ad9cad59")

	c.RemoveId(objid) //删除

	c.UpdateId(objid, map[string]interface{}{"id": 8, "name": "aaaaa", "age": 30}) //改

	var one map[string]interface{}
	c.FindId(objid).One(&one) //查询符合条件的一行数据
	fmt.Println(one)

	var result []map[string]interface{}
	c.Find(nil).All(&result) //查询全部
	fmt.Println(result)
}
