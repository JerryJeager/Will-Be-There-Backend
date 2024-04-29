package http

import (
	"mime/multipart"
	"net/http"

	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/JerryJeager/will-be-there-backend/service/event"
	"github.com/JerryJeager/will-be-there-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventController struct {
	serv event.EventSv
}

func NewEventController(serv event.EventSv) *EventController {
	return &EventController{serv: serv}
}

func (o *EventController) GetEvent(ctx *gin.Context) {
	var pp EventIDPathParam

	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid format"})
		return
	}

	Event, err := o.serv.GetEvent(ctx, uuid.MustParse(pp.EventID))
	if err != nil {
		ctx.Status(http.StatusNotFound)
	}

	ctx.JSON(http.StatusOK, *Event)
}

func (o *EventController) GetEvents(ctx *gin.Context) {
	var pp UserIDPathParm
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid uuid format"})
		return
	}

	Events, err := o.serv.GetEvents(ctx, uuid.MustParse(pp.UserID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *Events)
}

func (o *EventController) CreateEvent(ctx *gin.Context) {
	var Event service.Event

	if err := ctx.ShouldBindJSON(&Event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid format", "error": err.Error()})
		return
	}

	id, err := o.serv.CreateEvent(ctx, &Event)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *EventController) UpdateImageurl(ctx *gin.Context) {
	var pp EventIDPathParam

	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event id is of invalid format"})
		return
	}

	filename, ok := ctx.Get("filePath")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
	}

	file, ok := ctx.Get("file")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	imageUrl, err := utils.UploadToCloudinary(file.(multipart.File), filename.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = o.serv.UpdateImageUrl(ctx, uuid.MustParse(pp.EventID), imageUrl)

	if err != nil{
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, imageUrl)

}

func (o *EventController) DeleteEvent(ctx *gin.Context) {
	var pp EventIDPathParam

	if err := ctx.ShouldBindUri(&pp); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event id is of invalid format"})
		return
	}

	err := o.serv.DeleteEvent(ctx, uuid.MustParse(pp.EventID))

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}