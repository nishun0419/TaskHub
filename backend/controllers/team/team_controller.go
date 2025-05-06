package team

import (
	domain "backend/domain/team"
	usecase "backend/usecase/team"
	utils "backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	usecase *usecase.TeamUsecase
}

func NewTeamController(u *usecase.TeamUsecase) *TeamController {
	return &TeamController{usecase: u}
}

func (c *TeamController) CreateTeam(ctx *gin.Context) {
	var input domain.CreateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	customerID, ok := ctx.Get("customer_id")
	if !ok {
		utils.ErrorResponse(ctx, "Customer ID not found")
		return
	}
	err := c.usecase.CreateTeam(input, customerID.(int))
	if err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "チームが作成されました", nil)
}

func (c *TeamController) GetTeam(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(ctx, "Invalid team ID")
		return
	}
	team, err := c.usecase.GetTeam(id)
	if err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Team retrieved successfully", team)
}

func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	var input domain.UpdateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	if err := c.usecase.UpdateTeam(input); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Team updated successfully", nil)
}

func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(ctx, "Invalid team ID")
		return
	}
	if err := c.usecase.DeleteTeam(id); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "Team deleted successfully", nil)
}

func (c *TeamController) GetTeamsByCustomerID(ctx *gin.Context) {
	customerID, ok := ctx.Get("customer_id")
	if !ok {
		utils.ErrorResponse(ctx, "Customer ID not found")
		return
	}
	teams, err := c.usecase.GetTeamsByCustomerID(customerID.(int))
	if err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teams,
	})
}

func (c *TeamController) GenerateInviteToken(ctx *gin.Context) {
	var input domain.InviteTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}
	teamID := ctx.Param("team_id")
	teamIDInt, err := strconv.Atoi(teamID)
	if err != nil {
		utils.ErrorResponse(ctx, "Invalid team ID")
		return
	}
	input.TeamID = teamIDInt
	token, err := c.usecase.GenerateInviteToken(input)
	if err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}
	utils.SuccessResponse(ctx, "Invite token generated successfully", token)
}

func (c *TeamController) JoinTeam(ctx *gin.Context) {
	var input domain.JoinTeamInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}
	customerID, ok := ctx.Get("customer_id")
	if !ok {
		utils.ErrorResponse(ctx, "Customer ID not found")
		return
	}
	teamID, err := c.usecase.JoinTeam(customerID.(int), input)
	if err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}
	utils.SuccessResponse(ctx, "Team joined successfully", teamID)
}
