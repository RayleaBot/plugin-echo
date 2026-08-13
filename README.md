# Echo

RayleaBot 官方插件 · `raylea.echo`

把你发给机器人的内容原样回过来。用来确认协议接入、命令前缀、权限和消息回复是否通畅。

## 功能

- 响应 `/echo`，把命令后面的文字原样发回当前群聊或私聊
- 没有附带内容时回复 `(空消息)`，用来确认命令本身已被识别
- 群聊、私聊都可以用，不需要超级管理员

## 安装

本插件独立发布，不随 RayleaBot 主程序打包。安装后默认停用，需要在插件列表里**启用**才会响应命令。

### 插件商店

1. 打开 Web 管理面，进入 [插件商店](https://github.com/RayleaBot/RayleaBot/blob/main/docs/user/management-surface.md)（`/plugins/store`）。
2. 找到 **Echo**，安装与当前系统匹配的版本。
3. 安装前确认：插件会作为本机原生程序运行。
4. 到 **插件列表**（`/plugins`）启用 `raylea.echo`。

### 本地安装包

也可以在插件列表中安装本仓库 [GitHub Release](https://github.com/RayleaBot/plugin-echo/releases) 里对应平台的 ZIP：

| 平台 | 资源 |
| --- | --- |
| Windows x64 | `windows-x64` |
| Linux x64 | `linux-x64` |
| macOS arm64 | `macos-arm64` |

## 使用方法

命令前缀以管理面 **插件设置** 为准，下面按默认前缀 `/` 书写。

| 命令 | 权限 | 说明 |
| --- | --- | --- |
| `/echo [内容]` | 所有人 | 复读收到的内容；内容可省略 |

示例：

```text
你：/echo 你好
机器人：你好

你：/echo
机器人：(空消息)
```

多段文字会按空格拼回去，例如 `/echo 今天 天气 不错` 会回复 `今天 天气 不错`。

## 说明

- 这是链路验证插件，不会改配置、不会访问外网、也不会写存储。
- 若完全没有回复，先确认插件已启用，再核对命令前缀、黑白名单和权限策略。
- 若只有部分群能用，检查该群是否在黑名单中，或白名单是否已开启且未包含该群。

## 开发

插件以独立 Go 模块发布。生产包由本仓库 GitHub Actions 构建，RayleaBot 主程序不打包这份源码。

### 目录结构

```text
plugin-echo/
  cmd/echo/              进程入口
  internal/plugin/       协议处理、业务逻辑和测试
  tools/build/           调用 RayleaBot SDK 构建 artifact
  info.json              插件能力与发布元数据
```

### 本地联调

1. 将本仓库路径写入 RayleaBot 根目录下被 Git 忽略的 `plugin-workspace.local.json`：

```json
{
  "workspace_version": "1",
  "plugins": [
    {
      "id": "raylea.echo",
      "path": "../RayleaBotPlugins/plugin-echo"
    }
  ]
}
```

2. 在 **RayleaBot 主仓库根目录** 启动：

```powershell
$env:RAYLEA_PLUGIN_DEV = "watch"
$env:RAYLEA_SERVER_RELOAD = "watch"
.\start.bat
```

启动器会连接本地 Go SDK、构建当前平台 artifact，并通过离线 `plugin dev-sync` 同步到 `plugins/installed/`。构建失败时继续使用上一个已安装产物。

### 测试与构建

```powershell
go test -race ./...
go run ./tools/build -target windows-x64
```

### 发布

`v*` 标签对应的发布工作流使用固定的 RayleaBot SDK 引用构建 Windows x64、Linux x64 和 macOS arm64 包，并创建 GitHub Release。[plugin-catalog](https://github.com/RayleaBot/plugin-catalog) 记录各平台 ZIP 与 manifest 的 SHA-256；RayleaBot 插件商店只消费签名目录。

本地联调与商店分发说明见 [插件商店与独立开发](https://github.com/RayleaBot/RayleaBot/blob/main/docs/plugin/store-and-development.md)。

## License

[MIT](./LICENSE)
