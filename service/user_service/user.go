package user_service

import (
	"MessageBoard/mysql"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Message struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

func SendMessage(ctx *gin.Context) {
	db := mysql.DB.GetDb()
	message := new(Message)
	err := ctx.ShouldBind(message)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
	}
	message.CreateAt = time.Now()
	db.Create(message)
	ctx.JSON(200, gin.H{
		"code": message,
		"msg":  "success",
	})

}

func ShowMessage(ctx *gin.Context) {
	db := mysql.DB.GetDb()

	var message []Message
	var total int64
	//String转int
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	//判断会否需要分页
	offsetVal := (pageNum + 1) * pageSize
	if pageNum == 0 && pageSize == 0 {
		offsetVal = -1
	}
	db.Model(message).Count(&total).Limit(pageSize).Offset(offsetVal).Order("create_at desc").Find(&message)
	if len(message) == 0 {
		ctx.JSON(200, gin.H{
			"msg":  "没有数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":  "查询成功",
			"code": 200,
			"data": gin.H{
				"list":     message,
				"total":    total,
				"pageNum":  pageNum,
				"pageSize": pageSize,
			},
		})
	}
}

func DetectMessage(ctx *gin.Context) {
	db := mysql.DB.GetDb()
	message := new(Message)
	id, _ := strconv.Atoi(ctx.Query("id"))
	db.Find(&message)
	db.Delete(&Message{}, id)
	ctx.JSON(200, gin.H{
		"msg":  "删除成功",
		"code": 200,
	})
}
