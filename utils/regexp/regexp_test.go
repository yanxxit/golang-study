package regexp

import (
	"fmt"
	"testing"
)

func TestIsMobile(t *testing.T) {
	fmt.Println(IsMobile("17721021494"))
	fmt.Println(IsMobile("27721014942"))
}

func TestIsOpenid(t *testing.T) {
	fmt.Println("校验openID")
	fmt.Println(IsOpenid("oKXUCj1MOddnp-sCpGi1J1dg3Ty+"))
	fmt.Println(IsOpenid("oKXUCjyMRHD12cQLDxZiLpxtEdJEs"))
	fmt.Println(IsOpenid("oKXUCjyMRHD12cQLDxZiLpxtEdJ"))
}

func TestIsEmail(t *testing.T) {
	fmt.Println("校验邮箱")
	fmt.Println(IsEmail("1@qq.cn"))
}
func TestFormatNumber(t *testing.T) {
	fmt.Println("格式话设备号")
	fmt.Println(FormatNumber("17721021494"))
	fmt.Println(FormatNumber("12345678"))
	fmt.Println(FormatNumber("AD1234567812"))
}

func TestFormatName(t *testing.T) {
	fmt.Println("格式化姓名")
	fmt.Println(FormatName("张三"))
	fmt.Println(FormatName("周大装"))
	fmt.Println(FormatName("上海电信"))
}
