package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Snow-kal/ci-demo/internal/greeting"
)

func main() {
	// 定义命令行参数 `--name`。
	// 第一个参数是 flag 名称，第二个参数是默认值，第三个参数是帮助说明。
	name := flag.String("name", "CI/CD Learner", "name to show in greeting")
	// 解析命令行参数，把用户输入写回 `name` 指针指向的位置。
	flag.Parse()

	// 调用业务层函数生成问候语。
	// 这里把 `*name` 解引用为 string 传入，返回值包含正常结果和错误信息。
	message, err := greeting.Build(*name)
	if err != nil {
		// 出错时把错误信息打印到标准错误输出（stderr），便于脚本和 CI 区分普通输出。
		fmt.Fprintln(os.Stderr, err)
		// 以非 0 状态码退出进程，告诉调用方“本次执行失败”。
		os.Exit(1)
	}

	// 成功时把最终问候语输出到标准输出（stdout）。
	fmt.Println(message)
}
