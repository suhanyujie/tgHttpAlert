package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tgHttpAlert/internal/service"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/tg/alert1", handlerForTgAlert1)
	router.POST("/tg/alert1", handlerForTgAlert1ForPost)
	router.Run(":9101") // 监听并在 0.0.0.0:8080 上启动服务
}

// 接受一些参数，发送信息到告警群 todo
func handlerForTgAlert1(c *gin.Context) {
	resMap := gin.H{
		"code": -1,
	}
	env := c.DefaultQuery("env", "")
	if len(env) == 0 {
		resMap["msg"] = "param env is null"
		c.JSON(200, resMap)
		return
	}
	serverFlag := c.DefaultQuery("serverFlag", "")
	if len(serverFlag) == 0 {
		resMap["msg"] = "param serverFlag is null"
		c.JSON(200, resMap)
		return
	}
	app := c.DefaultQuery("app", "")
	if len(app) == 0 {
		resMap["msg"] = "param app is null"
		c.JSON(200, resMap)
		return
	}
	msg := c.DefaultQuery("msg", "")
	if len(msg) == 0 {
		resMap["msg"] = "param msg is null"
		c.JSON(200, resMap)
		return
	}
	err := service.NewAlertMsg(Convert2Struct(env), Convert2Struct(serverFlag), Convert2Struct(app), Convert2Struct(msg)).
		Send()
	if err != nil {
		fmt.Printf("%+W \n", err)
	}
	resMap["code"] = 1
	c.JSON(200, resMap)
}

type AlertReq struct {
	Env        string `json:"env" binding:"required"`
	ServerFlag string `json:"serverFlag" binding:"required"`
	App        string `json:"app" binding:"required"`
	Msg        string `json:"msg" binding:"required"`
}

func handlerForTgAlert1ForPost(c *gin.Context) {
	resMap := gin.H{
		"code": -1,
	}
	alertReq := AlertReq{}
	if errA := c.ShouldBindJSON(&alertReq); errA != nil {
		resMap["msg"] = "param alertReq is invalid | " + errA.Error()
		c.JSON(http.StatusOK, resMap)
		return
	}
	err := service.NewAlertMsg(alertReq.Env, alertReq.ServerFlag, alertReq.App, alertReq.Msg).
		Send()
	if err != nil {
		fmt.Printf("%+W \n", err)
		resMap["msg"] = err.Error()
	} else {
		resMap["msg"] = "ok"
		resMap["code"] = 0
	}
	c.JSON(200, resMap)
}

func Convert2Struct[T any](v T) string {
	return fmt.Sprintf("%v", v)
}
