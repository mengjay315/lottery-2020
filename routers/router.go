package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mengjay315/lottery/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sign_in", api.SignIn) // 签到

	r.POST("/vote_id", api.SetVoteID) // 获取投票信号
	r.POST("/vote", api.Vote) // 投票

	// 查询
	r.GET("/persons", api.GetPersons) // 所有参会人员
	r.GET("/vote_res", api.GetVoteRes) // 投票结果
	r.GET("/signin_res", api.GetSignInRes) // 获取签到结果 // 改为websocket返回前端数据

	return r
}
