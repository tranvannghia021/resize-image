package src

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type payloadResize struct {
	image string `json:"image,required"`
}

func ResizeHandler(ctx context.Context, r *app.RequestContext) {

	image, ok := r.GetPostForm("image")
	if !ok {
		responseError(r, "image is required")
		return
	}
	// service image
	service := NewResize(image)

	// check re-size duplicate

	isDuplicate, body, err := service.IsReSizeAgain()

	if isDuplicate {
		responseError(r, "[property] :"+err.Error())
		return
	}
	// re-size
	result, err := service.ReSize(body)

	if err != nil {
		responseError(r, "[re-size] :"+err.Error())
		return
	}
	r.JSON(consts.StatusOK, utils.H{
		"status":  true,
		"code":    consts.StatusOK,
		"message": "",
		"data": map[string]string{
			"image": result,
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
