package initial

import "fmt"

func init() {
	if err := InitConfig(); err != nil {
		panic(fmt.Sprintf("解析配置文件发生异常%v", err))
	}
}
