package model

import (
	"fmt"
	"golang-study/gin/config"
	"testing"
)

func TestConf_GetConf(t *testing.T) {
	fmt.Println(config.Conf.Port)
}
