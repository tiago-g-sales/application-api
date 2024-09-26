package controllers

import (
	"github.com/example/application-api/src/domain/usecases"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"net/http"
	"strconv"
)

type DemoController struct {
	DemoGetAll  usecases.IDemoGetAll
	DemoGetByID usecases.IDemoGetByID
}

func NewDemoController() *DemoController {
	return &DemoController{
		DemoGetAll:  usecases.NewDemoGetAll(),
		DemoGetByID: usecases.NewDemoGetByID(),
	}
}

func (c *DemoController) Routes() []restserver.Route {
	return []restserver.Route{
		{
			URI:      "demos",
			Method:   http.MethodGet,
			Function: c.GetAll,
			Prefix:   restserver.PublicApi,
		},
		{
			URI:      "demos/{id}",
			Method:   http.MethodGet,
			Function: c.GetById,
			Prefix:   restserver.PublicApi,
		},
	}
}

// @Summary Get all
// @Tags demo
// @Accept json
// @Produce json
// @Success 200 {array} models.Demo
// @Param X-AccountantId header uint64 true "id da contabilidade" minimum(0)
// @Param X-TenantId header uint64 true "id do dono do negócio" minimum(0)
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Router /private-api/rest/demo [get]
func (c *DemoController) GetAll(ctx restserver.WebContext) {
	result, err := c.DemoGetAll.Execute(ctx.Context())
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}
	ctx.JsonResponse(http.StatusOK, result)
}

// @Summary Get by id
// @Tags demo
// @Accept json
// @Produce json
// @Success 200 {object} models.Demo
// @Param X-AccountantId header uint64 true "id da contabilidade" minimum(0)
// @Param X-TenantId header uint64 true "id do dono do negócio" minimum(0)
// @Param X-UserId header uint64 true "id do usuário" minimum(0)
// @Param id path uint64 true "Demo ID"
// @Router /private-api/rest/demo/{id} [get]
func (c *DemoController) GetById(ctx restserver.WebContext) {
	paramId, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		ctx.ErrorResponse(http.StatusBadRequest, err)
		return
	}

	model, err := c.DemoGetByID.Execute(ctx.Context(), paramId)
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
		return
	}

	ctx.JsonResponse(http.StatusOK, model)
}
