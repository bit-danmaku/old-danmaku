package main

import (
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"strconv"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/gin-gonic/gin"
)

var (
	serviceName = "old-danmaku"
	version     = "latest"
)

func main() {
	dbc := InitDB()
	log.Info(dbc)

	httpSrv := httpServer.NewServer(
		server.Name(serviceName),
		server.Address(":8080"),
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// register router
	demo := newDemo(dbc)
	demo.InitRouter(router)

	hd := httpSrv.NewHandler(router)
	if err := httpSrv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	// Create service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Server(httpSrv),
		micro.Registry(registry.NewRegistry()),
	)
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//demoRouter
type demoRouter struct {
	dbConnector *DBConnector
}

type danmaku struct {
	Author string  `json:"author" binding:"required"`
	Time   float64 `json:"time" binding:"required"`
	Text   string  `json:"text" binding:"required"`
	Color  uint32  `json:"color"`
	Type   uint8   `json:"type"`
}

// [float64, uint8, uint32, string, string]
type danmakuResp = [5]interface{}

func newDemo(dbc *DBConnector) *demoRouter {
	return &demoRouter{
		dbConnector: dbc,
	}
}

func (a *demoRouter) InitRouter(router *gin.Engine) {
	router.POST("/channel/:id", a.PostDanmaku)
	router.GET("/channel/:id", a.GetDanmakuList)
}

func (a *demoRouter) PostDanmaku(c *gin.Context) {
	channelID, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(501, gin.H{"code": 1, "msg": "Failed When Parse Channnel ID."})
		return
	}
	strconv.Atoi(c.Param("id"))
	var dmk danmaku

	if err := c.ShouldBindJSON(&dmk); err == nil {
		log.Infof("get body: %+v", dmk)
		d := fromDanmakuPost(channelID, dmk)
		log.Warn(d)
		a.dbConnector.AddDanmaku(d)
	} else {
		log.Error(err)
		c.JSON(501, gin.H{"code": 2, "msg": "Failed When Get Post Data."})
	}

	c.JSON(200, gin.H{"code": 0, "data": dmk})
}

func (a *demoRouter) GetDanmakuList(c *gin.Context) {
	//channelID := c.Param("id")
	// TODO: 
	// your works are:
	// 1. query data from database
	// 2. process data to type `[]danmakuResp`
	// 3. return

	data := danmakuResp{1, 2, 3, "author", "hello world"}

	c.JSON(200, gin.H{"msg": data})
}
