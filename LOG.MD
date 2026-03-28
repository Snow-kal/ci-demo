# Git 提交记录与每一步作用

下面是本项目建议的提交拆分方式（也是这次初始化使用的拆分），你可以用下面命令查看真实哈希：

```bash
git log --oneline --decorate --graph
```

## 提交 1

- 提交信息：`feat: bootstrap go app and unit tests`
- 作用：
  - 初始化 `go.mod`
  - 创建 `cmd/cicd-demo/main.go` 作为入口程序
  - 创建 `internal/greeting` 业务逻辑和单元测试
  - 形成“可编译 + 可测试”的最小闭环

## 提交 2

- 提交信息：`ci: add github actions workflows for ci and cd`
- 作用：
  - 新增 `.github/workflows/ci.yml`：提交即自动格式检查、静态检查、测试、构建
  - 新增 `.github/workflows/cd.yml`：版本标签触发自动打包并发布 Release
  - 让仓库具备自动化质量门禁与发布能力

## 提交 3

- 提交信息：`docs: document cicd workflow and commit log`
- 作用：
  - 编写 `README.MD` 解释结构和工作原理
  - 编写 `LOG.MD` 解释每次提交的目的
  - 让新同学可以快速理解并复现 CI/CD 流程
