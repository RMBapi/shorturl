package routes

import (
	"net/http"

	"example.com/shorturl/models"
	"github.com/gin-gonic/gin"
)

func HandleRequest(context *gin.Context) {
	if context.Request.Method == http.MethodPost {
		var u models.URL

		type requestBody struct {
			LongURL *string `json:"longurl"`
		}

		var req requestBody
		err := context.ShouldBindJSON(&req)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse request"})
			return
		}

		u.LongURL = *req.LongURL

		err = u.UrlManager()

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create URL"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Short URL Created", "event": u.Shorturl})

	} else if context.Request.Method == http.MethodGet {

		type requestBody struct {
			ShortURL *string `json:"shorturl"`
		}

		var req requestBody
		err := context.ShouldBindJSON(&req)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse request"})
			return
		}

		long_key, err := models.Urlretrive(*req.ShortURL)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Can't Find the Long key"})
			return
		}

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create URL"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Actual URL", "event": long_key})

	}
}
