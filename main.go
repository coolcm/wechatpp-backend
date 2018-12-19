package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/controller/chat"
	"github.com/sjtucsn/wechatpp-backend/controller/paper"
	"github.com/sjtucsn/wechatpp-backend/controller/solution"
	"github.com/sjtucsn/wechatpp-backend/controller/transaction"
	"github.com/sjtucsn/wechatpp-backend/controller/user"
	"github.com/sjtucsn/wechatpp-backend/model"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	group := r.Group("/api/v1")
	{
		group.POST("/user/create", user.HandleCreateUser)
		group.POST("/user/logout", user.HandleUserLogout)
		group.GET("/user/query", user.HandleQueryUser)

		group.POST("/paper/upload", paper.HandleUploadExamPaper)
		group.GET("/paper/query", paper.HandleQueryExamByType)
		group.GET("/paper/detail", paper.HandleQueryExamAndSolutions)
		group.GET("/paper/download", paper.HandleDownloadExamPaper)

		group.POST("/chat/create", chat.HandleCreateChat)
		group.POST("/chat/score", chat.HandleScoreChat)
		group.POST("/chat/end", chat.HandleEndChat)

		group.POST("/transaction/create", transaction.HandleCreateTransaction)
		group.GET("/transaction/query", transaction.HandleQueryTransactionByHash)

		group.POST("/solution/upload", solution.HandleUploadSolution)
		group.GET("/solution/query", solution.HandleQuerySolution)
		group.GET("/solution/authority", solution.HandleSolutionAuthority)
		group.GET("/solution/download", solution.HandleDownloadSolutions)
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