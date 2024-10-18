package controller

import (
	"gin-demo/model"
	"gin-demo/pkg/response"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitStudentRouter(r *gin.RouterGroup) {
	// 路由组
	v1 := r.Group("student")
	{
		v1.POST("/", AddStudent)
		v1.DELETE("/:id", DeleteStudent)
		v1.PUT("/", UpdateStudent)
		v1.GET("/:id", GetStudent)
	}
}

func AddStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.AddStudent(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.DeleteStudent(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func UpdateStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.UpdateStudent(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	student, err := service.GetStudent(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(student, c)
}
