package opening

import (
	"net/http"
	"openingjobs/pkg/response"

	"github.com/gin-gonic/gin"
)

type OpeningHandler struct {
	usecase *OpeningUseCase
}

func newOpeningHandler(usecase *OpeningUseCase) *OpeningHandler {
	return &OpeningHandler{
		usecase: usecase,
	}
}

func (h *OpeningHandler) createOpeningHandler(ctx *gin.Context) {
	openinRequest := CreateOpeningRequest{}
	if err := ctx.BindJSON(&openinRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error binding json")
		return
	}

	opening, err := h.usecase.createOpening(&openinRequest)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error creating opening")
		return
	}

	response.Success(ctx, http.StatusCreated, opening)
}

func (h *OpeningHandler) showOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	opening, err := h.usecase.showOpening(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error show opening")
		return
	}

	if opening == nil {
		response.Error(
			ctx,
			http.StatusNotFound,
			response.NotFoundError(),
			"error show opening",
		)
		return
	}

	response.Success(ctx, http.StatusOK, opening)
}

func (h *OpeningHandler) listOpeningHandler(ctx *gin.Context) {
	openings, err := h.usecase.listOpenings()
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error listing openings")
		return
	}
	response.Success(ctx, http.StatusOK, openings)
}

func (h *OpeningHandler) updateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	openinRequest := CreateOpeningRequest{}
	if err := ctx.BindJSON(&openinRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error binding json")
		return
	}

	opening, err := h.usecase.updateOpening(id, &openinRequest)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error creating opening")
		return
	}

	response.Success(ctx, http.StatusCreated, opening)
}

func (h *OpeningHandler) deleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	err := h.usecase.deleteOpening(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, "error deleting opening")
		return
	}

	response.Success(ctx, http.StatusOK, "opening deleted successfully")
}

func (h *OpeningHandler) InitializeHandlers(routerGroup *gin.RouterGroup) {

	//Initialize routes
	routerGroup.POST("/opening", h.createOpeningHandler)
	routerGroup.GET("/opening", h.showOpeningHandler)
	routerGroup.GET("/openings", h.listOpeningHandler)
	routerGroup.PUT("/opening", h.updateOpeningHandler)
	routerGroup.DELETE("/opening", h.deleteOpeningHandler)
}
