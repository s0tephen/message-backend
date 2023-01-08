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

func ShowMessage(ctx *gin.Context) {
	db := mysql.DB.GetDb()

	var message []Message
	var total int64
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	//判断会否需要分页
	if pageNum == 0 {
		pageNum = 1
	}
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	//页-1*每页大小
	offset := (pageNum - 1) * pageSize

	db.Model(message).Count(&total).Limit(pageSize).Offset(offset).Order("create_at desc").Find(&message)
	if len(message) == 0 {
		ctx.JSON(200, gin.H{
			"msg":  "没有数据",
			"code": -1,
			"data": gin.H{},
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg":  "查询成功",
			"code": 1,
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
		"msg":  "success",
		"code": 1,
	})
}

func SendMessageAndUpdate(ctx *gin.Context) {
	db := mysql.DB.GetDb()

	var message Message
	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": -1,
			"msg":  "传入的json格式错误",
		})
	}
	if message.ID == 0 {
		message.CreateAt = time.Now()
		db.Create(&message)
		ctx.JSON(200, gin.H{
			"msg":  "留言成功",
			"code": 1,
		})
	} else {
		db.Model(&message).Where("id = ?", message.ID).Updates(map[string]interface{}{"name": message.Name, "content": message.Content})
		ctx.JSON(200, gin.H{
			"msg":  "修改成功",
			"code": 1,
		})
	}
}
