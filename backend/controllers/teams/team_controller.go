package teams

import (
	domain "backend/domain/teams"
	usecase "backend/usecase/teams"
	utils "backend/utils"
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
		utils.ErrorResponse(err.Error())
		return
	}

	if err := c.usecase.CreateTeam(input); err != nil {
		utils.ErrorResponse(err.Error())
		return
	}

	utils.SuccessResponse("Team created successfully", nil)
}

func (c *TeamController) GetTeam(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse("Invalid team ID")
		return
	}
	team, err := c.usecase.GetTeam(id)
	if err != nil {
		utils.ErrorResponse(err.Error())
		return
	}

	utils.SuccessResponse("Team retrieved successfully", team)
}

func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	var input domain.UpdateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(err.Error())
		return
	}

	if err := c.usecase.UpdateTeam(input); err != nil {
		utils.ErrorResponse(err.Error())
		return
	}

	utils.SuccessResponse("Team updated successfully", nil)
}

func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse("Invalid team ID")
		return
	}
	if err := c.usecase.DeleteTeam(id); err != nil {
		utils.ErrorResponse(err.Error())
		return
	}

	utils.SuccessResponse("Team deleted successfully", nil)
}
