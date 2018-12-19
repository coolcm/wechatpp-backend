package solution

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"github.com/sjtucsn/wechatpp-backend/utils"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// 处理上传试卷解答的请求
func HandleUploadSolution(c *gin.Context) {
	examHash := c.PostForm("exam_hash")
	solveId := c.PostForm("wechat_id")
	description := c.PostForm("description")
	if !utils.VerifyParams(c, map[string]string{"exam_hash": examHash, "wechat_id": solveId, "description": description}) {
		return
	}

	// 获取post的解答图片body
	if pic, err := c.FormFile("solution_image"); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "msg": "couldn't get post solution image"})
	} else {
		title := pic.Filename
		solution := model.CreateSolution(model.Db, examHash, solveId, description, title)

		if err := c.SaveUploadedFile(pic, path.Join("public", "solutions", solution.Hash, title)); err != nil {
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success", "solution": solution})
		}
	}
}

// 处理根据试卷哈希对相应解答列表的请求
func HandleQuerySolution(c *gin.Context) {
	examHash := c.Query("exam_hash")
	if !utils.VerifyParams(c, map[string]string{"exam_hash": examHash}) {
		return
	}

	solutions := model.QuerySolutionsByExamHash(model.Db, examHash)
	if num := len(solutions); num != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "size": num, "solutions": solutions})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "size": 0})
	}
}

// 处理用户是否有权限查看解答的请求
func HandleSolutionAuthority(c *gin.Context) {
	hash := c.Query("hash")
	userId := c.Query("wechat_id")
	if !utils.VerifyParams(c, map[string]string{"hash": hash, "wechat_id": userId}) {
		return
	}

	solution := model.QuerySolutionsByHash(model.Db, hash)
	if solution.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "solution does not exist"})
	} else {
		var flag bool //判断用户是否有权限查看该解答
		for _, v := range strings.Split(solution.AccessIds, ":") {
			if v == userId {
				flag = true
			}
		}
		if flag {
			c.JSON(http.StatusOK, gin.H{"status": "success", "info": "you have access to solutions"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "no access to solutions"})
		}
	}
}

// 处理下载试卷解答的请求
func HandleDownloadSolutions(c *gin.Context) {
	hash := c.Query("hash")
	userId := c.Query("wechat_id")
	if !utils.VerifyParams(c, map[string]string{"hash": hash, "wechat_id": userId}) {
		return
	}

	solution := model.QuerySolutionsByHash(model.Db, hash)
	if solution.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "solution does not exist"})
	} else {
		imagePath := path.Join("public", "solutions", solution.Hash, solution.Title)
		if reader, err := os.Open(imagePath); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"status": "fail", "info": "exam paper does not exist"})
		} else {
			io.Copy(c.Writer, reader)
		}
	}
}
