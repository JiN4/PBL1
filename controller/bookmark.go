package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PBL1/model"
	"github.com/PBL1/service"
	"github.com/gin-gonic/gin"
)

// GetBookmarkByUserID ...対象ユーザのブックマーク情報を取得する
func GetBookmarkByUserID(c *gin.Context) {
	bookmark := Bookmark{}
	bookmarks := []Bookmark{}
	recipe := model.Recipe{}
	var err error
	var recipeIDs []uint

	userID := c.Param("user_id")

	errCode := AuthCheck(c, userID)
	if errCode != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	recipeIDs, err = service.GetBookmarkRecipeIDsByUserID(userID)
	if err != nil {
		log.Println("レシピデータの取得ができませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for _, recipeID := range recipeIDs {

		recipe = service.GetRecipeByRecipeID(recipeID)

		bookmark.ID = recipe.ID
		bookmark.Name = recipe.Name
		bookmark.Description = recipe.Description
		bookmark.ImageURL = recipe.ImageURL
		bookmark.Price = "¥" + strconv.FormatUint(uint64(recipe.Price), 10)
		bookmark.Point = recipe.Point

		bookmarks = append(bookmarks, bookmark)
	}

	c.JSON(http.StatusOK, bookmarks)
}

func PostBookmarkByUserID(c *gin.Context) {
	userID := c.PostForm("user_id")
	recipeID, _ := strconv.Atoi(c.PostForm("recipe_id"))

	bookmark := model.Bookmark{}
	bookmark.UserID = userID
	bookmark.RecipeID = uint(recipeID)

	_, err := service.CreateBookmark(bookmark)
	if err != nil {
		log.Println("ブックマークデータの追加ができませんでした")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
