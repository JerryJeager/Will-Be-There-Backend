package http

import (
	"net/http"

	"github.com/JerryJeager/will-be-there-backend/service/invitees"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InviteeController struct {
	serv invitees.InviteeSv
}

func NewInviteeController(serv invitees.InviteeSv) *InviteeController {
	return &InviteeController{serv: serv}
}

func (o *InviteeController) CreateInvitee(ctx *gin.Context) {
	var invitee invitees.Invitee
	if err := ctx.ShouldBindJSON(&invitee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := o.serv.CreateInvitee(ctx, &invitee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *InviteeController) CreateInviteeByEmail(ctx *gin.Context) {
	var invitee invitees.InviteeByEmail
	if err := ctx.ShouldBindJSON(&invitee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := o.serv.CreateInviteeByEmail(ctx, &invitee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *InviteeController) GetInvitees(ctx *gin.Context) {
	var pp EventIDPathParam
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid uuid format"})
		return
	}

	invitees, err := o.serv.GetInvitees(ctx, uuid.MustParse(pp.EventID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *invitees)
}

func (o *InviteeController) GetInviteeByID(ctx *gin.Context) {
	var pp InviteeIDPathParam
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid uuid format"})
		return
	}

	invitee, err := o.serv.GetInviteeByID(ctx, uuid.MustParse(pp.InviteeID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *invitee)
}

func (o *InviteeController) UpdateInviteeStatus(ctx *gin.Context) {
	var pp InviteeIDPathParam
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "guest id is of invalid uuid format"})
		return
	}

	var status invitees.NewStatus

	if err := ctx.ShouldBindJSON(&status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "status be 'approved', 'pending', 'attending' or 'rejected'"})
		return
	}

	err := o.serv.UpdateInviteeStatus(ctx, uuid.MustParse(pp.InviteeID), &status)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (o *InviteeController) UpdateInvitee(ctx *gin.Context) {
	var pp InviteeIDPathParam
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "guest id is of invalid uuid format"})
		return
	}

	var invitee invitees.Invitee
	if err := ctx.ShouldBindJSON(&invitee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := o.serv.UpdateInvitee(ctx, uuid.MustParse(pp.InviteeID), &invitee)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)

}

func (o *InviteeController) DeleteInvitee(ctx *gin.Context) {
	var pp InviteeIDPathParam
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "guest id is of invalid uuid format"})
		return
	}

	err := o.serv.DeleteInvitee(ctx, uuid.MustParse(pp.InviteeID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
