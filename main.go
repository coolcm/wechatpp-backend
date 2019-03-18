package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/controller/chat"
	"github.com/sjtucsn/wechatpp-backend/controller/paper"
	"github.com/sjtucsn/wechatpp-backend/controller/solution"
	"github.com/sjtucsn/wechatpp-backend/controller/transaction"
	"github.com/sjtucsn/wechatpp-backend/controller/user"
	"github.com/sjtucsn/wechatpp-backend/controller/webSocket"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
)

func setupRouter() *gin.Engine {

	hub := utils.NewHub()

	r := gin.Default()
	r.GET("/ws", func(context *gin.Context) {
		webSocket.WsHandler(context, hub)
	})

	group := r.Group("/api/v1")
	{
		group.POST("/user/create", user.HandleCreateUser)
		group.POST("/user/logout", user.HandleUserLogout)
		group.GET("/user/query", user.HandleQueryUser)
		group.POST("/user/award", user.HandleAwardUserToken)

		group.POST("/paper/upload", paper.HandleUploadExamPaper)
		group.GET("/paper/query", paper.HandleQueryExamByType)
		group.GET("/paper/detail", paper.HandleQueryExamAndSolutions)
		group.GET("/paper/download", paper.HandleDownloadExamPaper)
		group.POST("/paper/credit/add", paper.HandleAddExamPaperCredit)

		group.POST("/chat/create", chat.HandleCreateChat)
		group.POST("/chat/score", chat.HandleScoreChat)
		group.POST("/chat/end", chat.HandleEndChat)
		group.GET("/chat/query", chat.HandleQueryChat)

		group.POST("/transaction/create", transaction.HandleCreateTransaction)
		group.GET("/transaction/query", transaction.HandleQueryTransactionByHash)

		group.POST("/solution/upload", solution.HandleUploadSolution)
		group.GET("/solution/query", solution.HandleQuerySolution)
		group.GET("/solution/authority", solution.HandleSolutionAuthority)
		group.GET("/solution/download", solution.HandleDownloadSolutions)
		group.POST("/solution/like/add", solution.HandleAddSolutionLikes)
	}

	return r
}

func main() {
	defer func() {
		model.Db.Close()
		println("hahaha goodbye")
	}()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}