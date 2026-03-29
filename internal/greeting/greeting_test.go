package greeting

import "testing"

// TestBuild 使用表驱动测试覆盖 Build 的主要输入路径：
// 1) 合法姓名应返回正确问候语
// 2) 空白姓名应返回错误
func TestBuild(t *testing.T) {
	// 标记当前测试可并行执行，缩短整体测试耗时。
	t.Parallel()

	// 定义测试用例结构体，统一管理输入、期望输出和期望错误。
	tests := []struct {
		name      string
		input     string
		wantErr   bool
		wantValue string
	}{
		{
			// 正常场景：合法输入应得到确定性输出。
			name:      "valid name",
			input:     "Go Dev",
			wantErr:   false,
			wantValue: "Hello Go Dev, your CI pipeline can build and test this app!",
		},
		{
			// 异常场景：仅空白字符应被识别为非法输入。
			name:    "blank name",
			input:   "   ",
			wantErr: true,
		},
	}

	// 遍历测试用例并为每个用例创建子测试，便于单独观察失败结果。
	for _, tc := range tests {
		// 重新绑定循环变量，避免并行子测试读到同一个最终迭代值。
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// 子测试并行执行，加快 CI 测试速度。
			t.Parallel()

			// 调用待测函数，获取实际结果。
			got, err := Build(tc.input)
			if tc.wantErr {
				// 期望报错的场景下，若未报错则测试失败。
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				// 该分支只关心“是否报错”，命中后可提前返回。
				return
			}

			// 不期望报错的场景下，任何错误都应判定为失败。
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			// 比较实际输出与期望输出，确保业务文案没有被意外改坏。
			if got != tc.wantValue {
				t.Fatalf("unexpected greeting, got %q, want %q", got, tc.wantValue)
			}
		})
	}
}
