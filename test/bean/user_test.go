package bean

import "testing"

func TestUserOk(t *testing.T) {
	GetOk()
	t.Log("TestUserOk test")
}

func TestGetWhoOk(t *testing.T) {
	res := GetWhoOk("xiaoxiao")
	t.Log("TestGo ooo" + res)
}
