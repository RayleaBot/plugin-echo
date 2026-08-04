# RayleaBot Echo Plugin

`raylea.echo` 为 RayleaBot 提供回显命令与消息链路验证能力。插件以独立 Go 模块发布，生产包由本仓库的 GitHub Actions 构建，不随 RayleaBot 主程序打包。

## 本地联调

将本仓库路径写入 RayleaBot 根目录下被忽略的 `plugin-workspace.local.json`，然后运行：

```powershell
$env:RAYLEA_PLUGIN_DEV = 'watch'
$env:RAYLEA_SERVER_RELOAD = 'watch'
node scripts/start-dev.mjs
```

启动脚本会连接本地 Go SDK、构建当前平台 artifact，并通过 Server 的离线 `plugin dev-sync` 命令原子同步到 `plugins/installed/`。构建失败时继续使用上一个已安装产物。

## 发布

推送 `v*` 标签后，工作流使用固定的 RayleaBot SDK 引用构建 Windows x64、Linux x64 和 macOS arm64 包，并创建 GitHub Release。发布后由插件目录仓库记录各平台 ZIP 与 manifest 的 SHA-256，RayleaBot 插件商店只消费签名目录。

License: MIT
