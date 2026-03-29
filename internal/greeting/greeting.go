package greeting

import (
	"fmt"
	"strings"
)

// Build 负责根据输入姓名生成问候语。
// 这是一个纯函数：不依赖外部状态，不做 IO，便于测试和在 CI 中稳定验证。
func Build(name string) (string, error) {
	// 先做预处理：去掉首尾空白字符（空格、换行、制表符等），
	// 避免用户输入 `"   "` 这类“看起来有内容、实际为空”的值。
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		// 返回业务错误而不是 panic，让调用方决定如何展示和处理失败。
		return "", fmt.Errorf("name must not be empty")
	}

	// 组装最终文案并返回。
	// `nil` 表示本次调用没有发生错误。
	return fmt.Sprintf("Hello %s, your CI pipeline can build and test this app!", cleanName), nil
}
