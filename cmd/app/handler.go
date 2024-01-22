package app

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/tranvannghia021/resize"
	"log"
	"resize-image/configs"
)

type payloadResize struct {
	Image string `json:"image,required"`
}

func ResizeHandler(ctx context.Context, r *app.RequestContext) {
	var imageS payloadResize
	err := r.BindAndValidate(&imageS)
	if err != nil {
		responseError(r, "[validate] :"+err.Error())
		return
	}
	// service image
	service := resize.NewResize(imageS.Image)

	// check re-size duplicate

	isDuplicate, body, err := service.IsReSizeAgain()

	if isDuplicate {
		log.Println(err.Error())
		responseError(r, "[property] :"+err.Error())
		return
	}
	// re-size
	result, err := service.ReSize(body)

	if err != nil {
		log.Println(err.Error())
		responseError(r, "[re-size] :"+err.Error())
		return
	}
	r.JSON(consts.StatusOK, utils.H{
		"status":  true,
		"code":    consts.StatusOK,
		"message": "",
		"data": map[string]string{
			"image": fmt.Sprintf("%s:%d/assets/%s", configs.GetAppUrl(), configs.GetPort(), result),
		},
	})
}

func responseError(r *app.RequestContext, message string) {
	r.JSON(consts.StatusUnprocessableEntity, utils.H{
		"status":  false,
		"code":    consts.StatusUnprocessableEntity,
		"message": message,
		"data":    nil,
	})
}
