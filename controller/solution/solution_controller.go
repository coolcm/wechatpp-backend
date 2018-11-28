package solution

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/model"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func HandleUploadSolution(c *gin.Context) {
	examHash := c.Query("exam_hash")
	solveId := c.Query("solve_id")
	pic, _ := c.FormFile("solution_image")
	title := pic.Filename
	solution := model.CreateSolution(model.Db, examHash, solveId, title)

	if err := c.SaveUploadedFile(pic, path.Join("public", "solutions", solution.Hash, title)); err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "solution": solution})
	}
}

func HandleQuerySolution(c *gin.Context) {
	examHash := c.Query("exam_hash")
	solutions := model.QuerySolutionsByExamHash(model.Db, examHash)
	if num := len(solutions); num != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "size": num, "solutions": solutions})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "size": 0})
	}
}

func HandleDownloadSolutions(c *gin.Context) {
	hash := c.Query("hash")
	userId := c.Query("user_id")
	solution := model.QuerySolutionsByHash(model.Db, hash)
	if solution.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "solution does not exist"})
	} else {
		var flag bool
		for _, v := range strings.Split(solution.AccessIds, ":") {
			if v == userId {
				flag = true
			}
		}
		if flag {
			imagePath := path.Join("public", "solutions", solution.Hash, solution.Title)
			if reader, err := os.Open(imagePath); err != nil {
				fmt.Println(err)
				c.JSON(http.StatusNotFound, gin.H{"status": "fail", "info": "exam paper does not exist"})
			} else {
				io.Copy(c.Writer, reader)
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "fail", "info": "no access to solutions"})
		}
	}
}
