package service

import (
	"github.com/PBL1/model"
)

// GetRecipeByMenuID ...メニューIDを受け取り、レシピのURLを返す
func GetRecipeByMenuID(menuID uint) (string, error) {
	modelRecipe := model.Recipe{}

	err := db.Where("menu_id = ?", menuID).First(&modelRecipe).Error

	return modelRecipe.URL, err
}

// CreateRecipe ...DBに与えられたデータをinsertする
func CreateRecipe(recipe model.Recipe) (model.Recipe, error) {
	err := db.Create(&recipe).Error
	if err != nil {
		return model.Recipe{}, err
	}
	return recipe, nil
}
