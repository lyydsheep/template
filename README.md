# Go 项目模板

本项目是一个基于 Gin 框架构建的高效 Go 项目模板，深度集成数据库连接、Redis 缓存、日志记录和链路追踪等核心功能。同时，借助 `wire` 实现依赖注入管理，助力开发者快速启动并搭建高质量的 Go 项目。

## 项目结构
```plaintext
/Users/lyydsheep/workspace/Go/projects/template
├── api/                    # API 相关代码，涵盖控制器、响应和请求数据结构定义，以及路由配置
│   ├── controller/         # 处理 HTTP 请求的控制器层
│   ├── reply/              # 定义响应数据结构
│   ├── request/            # 定义请求数据结构
│   └── router/             # 配置路由
├── common/                 # 公共组件，包含应用工具、枚举常量、错误码管理、日志记录和中间件等
│   ├── app/                # 应用相关工具，如分页和响应处理
│   ├── enum/               # 定义枚举常量
│   ├── errcode/            # 统一管理错误码
│   ├── logger/             # 日志记录模块
│   ├── middleware/         # Gin 中间件，如链路追踪和请求日志
│   └── util/               # 通用工具函数
├── config/                 # 配置文件及加载逻辑，可根据不同环境配置数据库、Redis 等信息
├── dal/                    # 数据访问层，负责数据库和 Redis 操作
│   ├── cache/              # Redis 缓存操作
│   ├── dao/                # 数据库操作
│   └── model/              # 定义数据库模型
├── docker-compose.yaml     # Docker 配置文件，用于快速启动数据库和 Redis
├── event/                  # 事件相关代码
├── go.mod                  # Go 模块依赖文件
├── go.sum                  # Go 模块依赖校验文件
├── imageData/              # 图片数据目录
├── init.sh                 # 项目初始化脚本，可自动替换占位符并初始化 Go 模块
├── library/                # 封装第三方库
├── log/                    # 存储日志文件
├── logic/                  # 业务逻辑层，包含领域模型、仓库接口和业务服务实现
│   ├── domain/             # 定义领域模型
│   ├── repository/         # 定义仓库接口
│   └── service/            # 实现业务服务
├── main.go                 # 项目入口文件
├── resources/              # 资源文件目录
├── wire.go                 # 依赖注入配置文件
└── wire_gen.go             # 依赖注入生成文件
```

## 快速开始
### 1. 克隆项目
```bash
git clone <your-repo-url> your-project-name
cd your-project-name
```
### 2. 初始化项目
运行初始化脚本 `init.sh` ，按照提示输入项目名称和模块名称，脚本会自动替换项目中的占位符并初始化 Go 模块。
```bash
chmod +x init.sh
./init.sh
```
### 3. 配置环境
根据实际需求修改 `config` 目录下的配置文件，如 `application.dev.yaml`、`application.prod.yaml` 和 `application.test.yaml`，配置数据库、Redis 等信息。
### 4. 启动服务
使用 Docker 快速启动数据库和 Redis：
```bash
docker-compose up -d
```
运行项目：
```bash
go run main.go
```
项目启动后，默认监听 `http://localhost:8080`。