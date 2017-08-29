package main

import "fmt"

func main() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("-------------------------")
	fmt.Println(kvs)
	fmt.Println("--")
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	fmt.Println("项目启动")

	show()
}

func show() {
	var myMap map[string]string
	/* 创建集合 */
	myMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	myMap["France"] = "Paris"
	myMap["Italy"] = "Rome"
	myMap["Japan"] = "Tokyo"
	myMap["India"] = "New Delhi"
	myMap["Linux"] = "sh"

	/* 使用 key 输出 map 值 */
	for country := range myMap {
		fmt.Println("Capital of", country, "is", myMap[country])
	}
	fmt.Println("--------------------")
	delete(myMap, "Paris");
	for country := range myMap {
		fmt.Println("Capital of", country, "is", myMap[country])
	}
	/* 查看元素在集合中是否存在 */
	captial, ok := myMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if (ok) {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}
