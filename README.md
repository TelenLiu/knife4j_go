# knife4j_go

采用的  https://github.com/xiaoymin/knife4j 项目的vue编译



由于用java发现，knife4j很好用，有其他网友也做了一个较早版本的发现有些bug，此版本封装最新knife4j的【4.5.0】 的ui。



具体使用请看example



命令行编译

```sh
swag init
```

swag 版本是1.16.3



```sh
go get github.com/TelenLiu/knife4j_go
```



main.go中添加

```go
//把swag 生成的 docs.go 或 swagger.json 写入
knife4j_go.SetDiyContent("doc.json", []byte(docs.SwaggerInfo.ReadDoc()))

//挂载knife4j主题ui到 你要目录，如 /doc
r.StaticFS("/doc", http.FS(knife4j_go.GetKnife4jVueDistRoot()))

//给前端配置，左上角的可切换的微服务模，默认自己 【/services.json】固定不变，json数组，由你决定
r.GET("/services.json", func(c *gin.Context) {
		c.String(200, `[
		    {
				"name": "AiTask服务",
				"url": "/doc.json",
				"swaggerVersion": "2.0",
				"location": "/doc.json",
			}
		]`)
	})
```



如此跑起来就额可以了。



