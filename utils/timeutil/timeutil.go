package timeutil

import (
	"time"
	"math/rand"
	"strconv"
)

var num = 0

func RandomString(length int) string {
	time.Sleep(1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := "1"
	for i := 0; i < length; i++ {
		t += "0"
	}
	num++
	k := strconv.Itoa(num)
	if num > 10 {
		num = 1
	}
	n, _ := strconv.Atoi(t)
	v := strconv.Itoa(r.Intn(n))
	if len(v) != length {
		return RandomString(length)
	} else {
		return v + k
	}
}
