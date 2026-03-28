# Go CI/CD Demo

这个项目是一个最小可运行的 Go CI/CD 学习样例，目标是帮助你理解：

- CI（持续集成）如何在每次代码提交后自动检查代码质量。
- CD（持续交付）如何在打版本标签后自动打包并发布产物。

## 项目结构

```text
.
|- .github/
|  |- workflows/
|     |- ci.yml          # 持续集成：格式、静态检查、测试、构建
|     |- cd.yml          # 持续交付：打 tag 后构建多平台二进制并发版
|- cmd/
|  |- cicd-demo/
|     |- main.go         # 程序入口
|- internal/
|  |- greeting/
|     |- greeting.go     # 核心业务逻辑
|     |- greeting_test.go# 单元测试
|- .gitignore
|- go.mod
|- README.MD
|- LOG.MD
```

## 工作方式（从提交到发布）

1. 开发者把代码 push 到 `main` 或发起 PR 到 `main`。
2. GitHub Actions 触发 `CI` 工作流（`.github/workflows/ci.yml`）。
3. CI 依次执行：
   - `gofmt`：检查格式统一性
   - `go vet`：检查常见静态问题
   - `go test`：运行单元测试
   - `go build`：验证能否成功构建
4. 当你 push 一个版本标签（如 `v1.0.0`）时，触发 `CD` 工作流（`.github/workflows/cd.yml`）。
5. CD 会：
   - 再跑一次测试（保证发布前质量）
   - 构建 Linux / macOS / Windows 三个平台可执行文件
   - 自动创建 GitHub Release，并上传 `dist/` 下的产物

## 本地运行

```bash
go test ./...
go run ./cmd/cicd-demo --name "Snow"
go build ./...
```

## 触发一次完整 CD 的步骤

```bash
git tag v0.1.0
git push origin v0.1.0
```

推送 tag 后，进入仓库 `Actions` 页面可以看到 `CD` 流程；在 `Releases` 页面可以下载构建产物。

## 你可以重点观察什么

- CI 失败时会阻止有问题的代码进入主分支。
- CD 只在版本标签触发，避免每次普通提交都发布。
- 流水线本质上是“触发器（on）+ 作业（jobs）+ 步骤（steps）”的组合。
