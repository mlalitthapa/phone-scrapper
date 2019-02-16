package search

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app"
	"github.com/mlalitthapa/phone-scrapper/app/shared"
	"github.com/mlalitthapa/phone-scrapper/utils"
	"net/http"
)

func Register(r *gin.RouterGroup) {
	r.GET("/search", Results)
}

func Results(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		app.ErrorResponse(c, http.StatusUnprocessableEntity, errors.New("invalid query"))
		return
	}

	doc, err := app.Scrape(fmt.Sprintf(utils.SearchUrl, query))
	if err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	devices := shared.GetDeviceList(doc)
	app.SuccessResponse(c, devices)
}
