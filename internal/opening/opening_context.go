package opening

import (
	"openingjobs/pkg/config"

	"github.com/gin-gonic/gin"
)

func InitializeOpeningHandlerConext(routerGroup *gin.RouterGroup) error {
	db := config.GetDB()
	qb := config.GetQueryBuilder()

	storage := newOpeningStorage(db, qb)
	useCase := newOpeningUseCase(storage)
	handle := newOpeningHandler(useCase)

	handle.InitializeHandlers(routerGroup)

	return nil
}
