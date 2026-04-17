package main

import (
	"example/component"
	"flag"

	// 引入 start 包 让包内 init 函数生效
	_ "example/start"

	"github.com/team-ide/framework"
)

func main() {

	framework.OutInfo()

	// 如果是 -v、-version 表示只是查看版本，直接退出
	if framework.IsVersion() {
		return
	}

	// 解析参数
	flag.Parse()

	err := component.Starter.Start()
	if err != nil {
		panic(err)
	}

	if component.Starter.ShouldWait() {
		component.Starter.Wait()
	}
}
