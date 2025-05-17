# Go 项目模板

这是一个基于 Gin 框架构建的 Go 项目模板，集成了数据库连接、Redis 缓存、日志记录、链路追踪等功能，同时使用 `wire` 进行依赖注入管理，能帮助开发者快速启动新的 Go 项目。

## 项目结构
```plaintext
/Users/lyydsheep/workspace/Go/projects/template
├── api/                    # API 相关代码
│   ├── controller/         # 控制器层，处理 HTTP 请求
│   ├── reply/              # 响应数据结构定义
│   ├── request/            # 请求数据结构定义
│   └── router/             # 路由配置
├── common/                 # 公共组件
│   ├── app/                # 应用相关工具，如分页、响应处理
│   ├── enum/               # 枚举常量定义
│   ├── errcode/            # 统一错误码管理
│   ├── logger/             # 日志记录模块
│   ├── middleware/         # Gin 中间件，如链路追踪、请求日志
│   └── util/               # 通用工具函数
├── config/                 # 配置文件及加载逻辑
├── dal/                    # 数据访问层
│   ├── cache/              # Redis 缓存操作
│   ├── dao/                # 数据库操作
│   └── model/              # 数据库模型定义
├── docker-compose.yaml     # Docker 配置文件，用于快速启动数据库和 Redis
├── event/                  # 事件相关代码
├── go.mod                  # Go 模块依赖文件
├── go.sum                  # Go 模块依赖校验文件
├── imageData/              # 图片数据目录
├── init.sh                 # 项目初始化脚本
├── library/                # 第三方库封装
├── log/                    # 日志文件存储目录
├── logic/                  # 业务逻辑层
│   ├── domain/             # 领域模型定义
│   ├── repository/         # 仓库接口定义
│   └── service/            # 业务服务实现
├── main.go                 # 项目入口文件
├── resources/              # 资源文件目录
├── wire.go                 # 依赖注入配置文件
└── wire_gen.go             # 依赖注入生成文件
```

## 快速开始
### 1. 克隆项目
```
git clone <your-repo-url> 
your-project-name
cd your-project-name
```
### 2. 初始化项目
运行初始化脚本 `init.sh` ，根据提示输入项目名称和模块名称，脚本会自动替换项目中的占位符并初始化 Go 模块。

```
chmod +x init.sh
./init.sh
```
### 3. 配置环境
根据实际需求修改 config 目录下的配置文件，如 application.dev.yaml 、 application.prod.yaml 和 application.test.yaml ，配置数据库、Redis 等信息。

### 4. 启动服务
使用 Docker 快速启动数据库和 Redis：

```
docker-compose up -d
```
运行项目：

```
go run main.go
```
项目启动后，默认监听 http://localhost:8080 。