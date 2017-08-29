package bean

func GetOk()(string){
	return "ok"
}

func GetWhoOk(name string)(rs string){
	rs = name +"ok"
	return
}