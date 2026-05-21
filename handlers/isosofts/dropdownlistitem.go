package isosoftsHandlers

import (
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DropDownListItemHandler struct {
}

func (*DropDownListItemHandler) DuplicateDefaults(c *gin.Context) {
	companyId := c.Query("companyId")

	if companyId == "" {
		c.IndentedJSON(400, gin.H{})
	}

	fmt.Println("i am gay")
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	dropDownListItemModel.DuplicateDefaults(companyId)
	c.IndentedJSON(201, gin.H{})
}
