package tableComponentHandlers

import (
	"algebra-isosofts-api/middlewares"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type DropDownListItemHandler struct {
}

func (*DropDownListItemHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	dropDownListItems, err := dropDownListItemModel.GetAll(map[string]interface{}{
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, tableComponentTypes.GroupDropDownListItems(dropDownListItems))
}
