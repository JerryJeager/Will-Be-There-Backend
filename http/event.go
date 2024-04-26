package http

import (
	"net/http"

	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/JerryJeager/will-be-there-backend/service/event"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid format"})
		return
	}

	id, err := o.serv.CreateEvent(ctx, &Event)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *EventController) CreateEventType(ctx *gin.Context) {
	var eventType service.EventType

	if err := ctx.ShouldBindJSON(&eventType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid format"})
		return
	}

	var pp EventIDPathParam

	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid format"})
		return
	}

	id, err := o.serv.CreateEventType(ctx, uuid.MustParse(pp.EventID), &eventType)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *EventController) UpdateEventType(ctx *gin.Context) {
	var EventPP EventIDPathParam

	if err := ctx.ShouldBindUri(&EventPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event id is of invalid format"})
		return
	}

	var eventTypeID EventTypeIDPathParam

	if err := ctx.ShouldBindUri(&eventTypeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event type id is of invalid format"})
		return
	}

	var eventType service.EventType

	if err := ctx.ShouldBindJSON(&eventType); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := o.serv.UpdateEventType(ctx, uuid.MustParse(EventPP.EventID), uuid.MustParse(eventTypeID.EventTypeID), &eventType)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (o *EventController) DeleteEventType(ctx *gin.Context) {
	var EventPP EventIDPathParam

	if err := ctx.ShouldBindUri(&EventPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event id is of invalid format"})
		return
	}

	var eventTypeID EventTypeIDPathParam

	if err := ctx.ShouldBindUri(&eventTypeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event type id is of invalid format"})
		return
	}

	err := o.serv.DeleteEventType(ctx, uuid.MustParse(EventPP.EventID), uuid.MustParse(eventTypeID.EventTypeID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
