# "随心记"Golang编程论坛
基于Gin和Vue搭建的前后端分离的论坛项目。

包含常见论坛所需的注册登录、评论等功能， 同时含有多个频道，
可以在这里随时记录你想说的话、发表你对于编程学习的看法。

上线网址：http://47.108.72.107:8084

参考代码：

[Q1mi/bluebell: bluebell (github.com)](https://github.com/Q1mi/bluebell)

[mao888/bluebell-plus: 基于vue+go+gin+mysql+redis的博客论坛web项目 (github.com)](https://github.com/mao888/bluebell-plus)

## 技能清单
1. 雪花算法生成分布式ID
2. Gin框架
3. Zap日志库
4. Viper配置管理
5. Swagger生成接口文档
6. JWT认证
7. 令牌桶限流
8. Go语言操作MySQL **(sqlx)**
9. Go语言操作Redis **(go-redis)**
10. Gihub热榜
11. Docker部署
12. Vue框架
13. ElementUI
14. axios 

## 项目目录结构
### 后端结构树
```bash
├─ bin                    可执行文件目录
├─ conf                   配置文件目录
├─ controller             控制器代码目录
├─ dao                    数据访问层代码目录
│  ├─ api                 第三方API代码目录
│  ├─ mysql               MySQL 数据访问代码目录
│  └─ redis               Redis 数据访问代码目录
├─ docs                   存放文档目录
├─ log                    存放日志文件目录
├─ logger                 日志相关代码目录
├─ logic                  业务逻辑代码目录
├─ middlewares            中间件代码目录
├─ models                 模型定义代码目录
├─ pkg                    存放通用包目录
│  ├─ jwt                 JWT 相关代码目录
│  └─ snowflake           雪花算法相关代码目录
├─ routers                路由配置代码目录
├─ settings               项目配置代码目录
├─ static                 存放静态文件目录
│  ├─ css                 样式表文件目录
│  ├─ fonts               字体文件目录
│  ├─ img                 图片文件目录
│  └─ js                  JavaScript 文件目录
├─ templates              存放模板文件目录
│  .air.conf              Air自动构建配置文件
│  .gitignore             Git忽略文件配置
│  docker-compose.yml     Docker Compose配置文件
│  Dockerfile             Docker镜像构建文件
│  go.mod                 Go模块依赖配置文件
│  go.sum                 Go模块依赖校验文件
│  init.sql               数据库初始化脚本文件
│  main.exe               主程序可执行文件（Windows）
│  main.go                主程序源代码文件
│  Makefile               Make工程文件
│  version.go             项目版本信息文件
└─ wait-for.sh            脚本文件用于等待服务启动

```
### 前端结构树
```bash
├── bin
│   └── bluebell
├── conf
│   └── config.yaml
├── static
│   ├── css
│   ├── favicon.ico
│   ├── img
│   └── js
└── templates
    └── index.html
```

## 项目预览图

![](https://s2.loli.net/2023/12/14/alqKoUPsAZO2hmM.png)

![img](https://s2.loli.net/2023/12/10/u3qgEZR4mpkt6We.png)

## 项目全套笔记

参考博客：[归档 | 李文周的博客 (liwenzhou.com)](https://www.liwenzhou.com/archives/)

- [《基于雪花算法生成用户ID》](https://www.yuque.com/docs/share/e50bbca1-e019-45e2-b77b-a9ba01fbede3?#) 
- [gin框架中使用validator若干实用技巧](https://www.liwenzhou.com/posts/Go/validator_usages/)
- [《限制账号同一时间只能登录一个设备》](https://www.yuque.com/docs/share/584ddd0f-5158-4cea-8918-a4b6e1d41a07?# )
- [《基于Cookie、Session和基于Token的认证模式介绍》](https://www.yuque.com/docs/share/06a89a55-3e3c-452b-aeb1-acf4d2bac8a5?#)
- [在gin框架中使用JWT认证](https://www.liwenzhou.com/posts/Go/jwt_in_gin/)
- [为Go项目编写Makefile](https://www.liwenzhou.com/posts/Go/makefile/)
- [使用Air实现Go程序实时热重载](https://www.liwenzhou.com/posts/Go/live_reload_with_air/)
- [分页](https://zhidao.baidu.com/question/1573826651037645420.html)
- [JSON实战拾遗之数字精度](https://www.ituring.com.cn/article/506822)
- [你需要知道的那些go语言json技巧](https://www.liwenzhou.com/posts/Go/json_tricks_in_go)
- [帖子投票（点赞）功能设计与实现](https://www.yuque.com/docs/share/d09afe84-90d1-4e04-a73e-95848f073558?#)
- [《基于用户投票的排名算法》](https://www.yuque.com/docs/share/f40f5c41-f327-47d4-88bb-02bcf62515a8?# )
- [使用swagger生成接口文档](https://www.liwenzhou.com/posts/Go/gin_swagger/)
- [HTTP Server常用压测工具介绍](https://www.liwenzhou.com/posts/Go/benchmark_tool/)
- [漏桶和令牌桶限流策略介绍及使用](https://www.liwenzhou.com/posts/Go/ratelimit/)
- [option选项模式](https://www.liwenzhou.com/posts/Go/functional_options_pattern/)
- [Go pprof性能调优](https://www.liwenzhou.com/posts/Go/performance_optimisation/)
- [如何使用docker部署Go Web程序](https://www.liwenzhou.com/posts/Go/how_to_deploy_go_app_using_docker/)
- [部署Go语言程序的N种方法](https://www.liwenzhou.com/posts/Go/deploy_go_app/)
- [《企业代码发布流程及CICD介绍》](https://www.yuque.com/docs/share/e837e5bf-f6a9-4dc8-98e4-4b8ce24808ab?)
