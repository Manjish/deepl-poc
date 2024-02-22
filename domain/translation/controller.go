package translation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) HandleRoot(c *gin.Context) {
	message := ctrl.service.GetMessage()
	c.JSON(http.StatusOK, gin.H{"message": message.Text})
}

func (ctrl *Controller) GetTransalation(c *gin.Context) {
	body := &Model{Text: "Hello"}
	// text := c.ShouldBindJSON(&body)
	translated, err := ctrl.service.GetTranslation(body.Text)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"translated": translated})
}
