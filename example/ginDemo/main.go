package main

import (
	"github.com/TelenLiu/knife4j_go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	r := gin.Default()
	r.StaticFS("/doc", http.FS(knife4j_go.GetKnife4jVueDistRoot()))
	r.GET("/services.json", func(c *gin.Context) {
		c.String(200, `
    [
    {
        "name": "knife4j-front服务模块",
        "url": "/static/server1.json",
        "swaggerVersion": "2.0",
        "location": "/static/server1.json"
    },
    {
        "name": "第二个模块",
        "url": "/static/server2.json",
        "swaggerVersion": "2.0",
        "location": "/static/server2.json"
    },
    {
        "name": "第3个模块",
        "url": "/static/server3.json",
        "swaggerVersion": "2.0",
        "location": "/static/server3.json"
    },
    {
        "name": "api接口样例",
        "location": "/api-docs.json",
        "swaggerVersion": "2.0",
    }
]
`)
	})

	ccc, _ := os.ReadFile("docs/api-docs.json")
	knife4j_go.SetDiyContent("api-docs.json", ccc)

	r.Run(":9077")
}
