# shortrul

[zero教程](https://gocn.vip/topics/10884)

## api
- 使用 goctl 生成 API Gateway 代码

```
cd api
goctl api go -api shorturl.api -dir .
```

- 目录信息
```
├── api
│   ├── etc
│   │   └── shorturl-api.yaml         // 配置文件
│   ├── internal
│   │   ├── config
│   │   │   └── config.go             // 定义配置
│   │   ├── handler
│   │   │   ├── expandhandler.go      // 实现expandHandler
│   │   │   ├── routes.go             // 定义路由处理
│   │   │   └── shortenhandler.go     // 实现shortenHandler
│   │   ├── logic
│   │   │   ├── expandlogic.go        // 实现ExpandLogic
│   │   │   └── shortenlogic.go       // 实现ShortenLogic
│   │   ├── svc
│   │   │   └── servicecontext.go     // 定义ServiceContext
│   │   └── types
│   │       └── types.go              // 定义请求、返回结构体
│   ├── shorturl.api
│   └── shorturl.go                   // main入口定义
├── go.mod
└── go.sum
```

## rpc


goctl rpc template -o transform.proto
