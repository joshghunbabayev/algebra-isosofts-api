package tableComponentHandlers

import (
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type DropDownListItemHandler struct {
}

func (*DropDownListItemHandler) GetAll(c *gin.Context) {
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	dropDownListItems, err := dropDownListItemModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, tableComponentTypes.GroupDropDownListItems(dropDownListItems))
}
