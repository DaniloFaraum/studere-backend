package controllers

import (
    "net/http"

    "github.com/DaniloFaraum/studere-backend/models"
    "github.com/DaniloFaraum/studere-backend/requests"
    "github.com/DaniloFaraum/studere-backend/utils"
    "github.com/gin-gonic/gin"
)

func CreateImageController(ctx *gin.Context) {
    request := requests.CreateImageRequest{}

    ctx.BindJSON(&request)

    if err := request.Validate(); err != nil {
        utils.HandleControllerError(ctx, http.StatusBadRequest, "image validation error", err)
        return
    }

    image := models.Image{
		EncryptedImage: request.EncryptedImage,
    }

    if err := db.Create(&image).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not create image on db", err)
        return
    }

    utils.SendSuccess(ctx, "create-image", request)
}

func ListImagesController(ctx *gin.Context) {
    var images []models.Image
    var imageResponses []models.ImageResponse

    if err := db.Find(&images).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find images", err)
        return
    }

    for _, image := range images {
        imageResponses = append(imageResponses, models.ImageResponse{
			ID: image.ID,
			EncryptedImage: image.EncryptedImage,
        })
    }

    utils.SendSuccess(ctx, "list-images", imageResponses)
}

func ShowImageController(ctx *gin.Context) {
    id := ctx.Query("id")
    if id == "" {
        utils.SendError(ctx, http.StatusBadRequest, utils.ErrParamIsrequired("id", "queryParam").Error())
        return
    }

    image := models.Image{}

    if err := db.First(&image, id).Error; err != nil {
        utils.HandleControllerError(ctx, http.StatusInternalServerError, "could not find image", err)
        return
    }

    utils.SendSuccess(ctx, "show-image", image)
}
