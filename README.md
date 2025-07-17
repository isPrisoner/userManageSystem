# 用户管理系统 (User Management System)

## 项目介绍

这是一个基于Go语言开发的用户管理系统，提供了完整的用户管理功能，包括用户注册、登录、查询、编辑、删除等操作。系统采用了MVC架构设计，前端使用HTML+TailwindCSS构建美观的用户界面，后端使用Go语言提供RESTful API服务。系统还包含数据统计和可视化功能，可以查看用户访问趋势和用户状态分布等信息。

## 技术栈

- **后端**: Go 1.23
- **Web框架**: Gorilla Mux
- **数据库**: MySQL 8.0
- **前端**: HTML5, TailwindCSS, JavaScript, Chart.js
- **其他**: Feather Icons

## 功能特性

- 用户认证与授权
  - 用户注册
  - 用户登录/登出
  - 密码找回
  - 会话管理
- 用户管理
  - 用户列表查看（分页）
  - 用户信息编辑
  - 用户删除
  - 按状态和用户名筛选
- 数据统计与可视化
  - 用户访问趋势图表
  - 用户状态分布统计
  - 月度注册/注销用户统计
- 系统日志
  - 操作日志记录
  - 错误日志记录

## 安装说明

### 前提条件

- Go 1.23 或更高版本
- MySQL 8.0 或更高版本

### 安装步骤

1. 克隆代码仓库

```bash
# 使用HTTPS
git clone https://github.com/isPrisoner/userManageSystem.git

# 或使用SSH
git clone git@github.com:isPrisoner/userManageSystem.git
```

2. 进入项目目录

```bash
cd userManageSystem
```

3. 安装依赖

```bash
go mod tidy
```

4. 配置数据库

创建MySQL数据库并导入以下表结构：

```sql
USE user_admin;

CREATE TABLE IF NOT EXISTS user (
  id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role ENUM('管理员','普通用户') NOT NULL DEFAULT '普通用户',
  email VARCHAR(255) NULL,
  image VARCHAR(255) DEFAULT 'default.jpg',
  status TINYINT(1) NOT NULL DEFAULT 1,
  session_id VARCHAR(255),
  last_login DATETIME,
  register_time DATETIME DEFAULT CURRENT_TIMESTAMP,
  delete_time DATETIME,
  delete_status TINYINT(1) DEFAULT 0,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS index_visits (
  id INT AUTO_INCREMENT PRIMARY KEY,
  visit_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

5. 修改数据库连接配置

编辑 `utils/conn.go` 文件，修改数据库连接信息：

```go
func InitDB() *sql.DB {
    // 修改为你的数据库连接信息
    db, err := sql.Open("mysql", "用户名:密码@tcp(主机地址:端口)/user_admin")
    if err != nil {
        log.Println(err)
        return nil
    }
    return db
}
```

6. 运行项目

```bash
go run main.go
```

7. 访问系统

打开浏览器，访问 `http://localhost:8090`

## 目录结构

```
userManageSystem/
├── handlers/       # 处理HTTP请求的控制器
├── images/         # 用户头像和系统图片
├── logs/           # 系统日志文件
├── middlewares/    # HTTP中间件
├── models/         # 数据模型和数据库操作
├── service/        # 业务逻辑层
├── utils/          # 工具函数和配置
├── view/           # HTML模板文件
├── go.mod          # Go模块定义
├── go.sum          # Go模块校验和
├── main.go         # 应用入口
└── README.md       # 项目说明文档
```

## 使用说明

### 登录系统

- 访问 `http://localhost:8090` 进入登录页面
- 输入用户名和密码登录
- 如忘记密码，可通过"忘记密码"功能找回

### 用户管理

- 在系统首页可查看数据概览和访问趋势
- 点击"用户管理"查看用户列表
- 可对用户进行编辑、删除等操作
- 支持按状态和用户名筛选用户

### 数据统计

- 系统首页展示用户注册、登录、注销等统计数据
- 可查看不同时间段的访问趋势图表

## 注意事项

- 默认监听端口为8090，可在main.go中修改
- 系统日志存储在logs目录下
- 用户头像存储在images目录下，默认头像为default.jpg

## 许可证

本项目采用MIT许可证，详情请查看LICENSE文件。

