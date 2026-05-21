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

func (*DropDownListItemHandler) Create(c *gin.Context) {
	var body struct {
		Type       string `json:"type"`
		Value      string `json:"value"`
		ShortValue string `json:"shortValue"`
	}

	var errs = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(errs) > 0 {
		c.JSON(400, gin.H{
			"token": "aaa",
			"errs":  errs,
		})
		return
	}

	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	dropDownListItemModel.Create(tableComponentTypes.DropDownListItem{
		Id:         dropDownListItemModel.GenerateUniqueId(),
		CompanyId:  account.CompanyId,
		Type:       body.Type,
		Value:      body.Value,
		ShortValue: body.ShortValue,
	})

	c.IndentedJSON(201, gin.H{})
}

func (*DropDownListItemHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	currentDropDownListItem, _ := dropDownListItemModel.GetById(Id)

	if currentDropDownListItem.IsEmpty() || currentDropDownListItem.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Value      string `json:"value"`
		ShortValue string `json:"shortValue"`
	}

	var errs = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(errs) > 0 {
		c.JSON(400, gin.H{
			"token": "aaa",
			"errs":  errs,
		})
		return
	}

	dropDownListItemModel.Update(Id, map[string]interface{}{
		"value":      body.Value,
		"shortValue": body.ShortValue,
	})

	c.JSON(200, gin.H{})
}

func (*DropDownListItemHandler) Delete(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	currentDropDownListItem, _ := dropDownListItemModel.GetById(Id)
	if currentDropDownListItem.IsEmpty() || currentDropDownListItem.CompanyId != account.CompanyId {
		c.JSON(400, gin.H{})
		return
	}

	dropDownListItemModel.Delete(Id)

	c.JSON(200, gin.H{})
}
