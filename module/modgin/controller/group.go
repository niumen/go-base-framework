package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-frame/go-base-framework/db"
	"github.com/go-frame/go-base-framework/g"
	"github.com/go-frame/go-base-framework/model"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func CreateGroup(ctx *gin.Context) {
	var data model.Group
	if err := ctx.Bind(&data); err != nil {
		fmt.Println("group create err", err.Error())
		ctx.JSON(http.StatusBadRequest, &g.Result{
			Code:    g.BadRequest,
			Message: g.FailedMsg,
		})
		return
	}
	logrus.Debug("data is ", data)
	data.Created = time.Now()
	res := db.FengConnect.Create(&data)
	if res.Error != nil {
		logrus.Debug("create err is ", res.Error)
	}
	ctx.JSON(http.StatusOK, &g.Result{
		Code:    g.Success,
		Message: g.SuccessMsg,
		Result:  data,
	})
}

func SearchGroup(ctx *gin.Context) {
	var data []*model.Group
	res := db.FengConnect.Find(&data)
	if res.Error != nil {
		logrus.Debug("search err is ", res.Error)
	}
	ctx.JSON(http.StatusOK, &g.Result{
		Code:    g.Success,
		Message: g.SuccessMsg,
		Result:  data,
	})
}