package user

import (
	"log"
	"net/http"

	"userapi-ddd-go/pkg/domain/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us user.UserService
}

type URI struct {
	Id string `uri:"id"`
}

func NewUserController(s user.UserService) *UserController {
	return &UserController{s}
}

// @Summary add a new user to the user list
// @ID create-user
// @Produce json
// @Param data body user.User true "user data"
// @Success 200 {object} user.User
// @Failure 400 {object} object
// @Router /user [post]
func (uc UserController) AddUser(c *gin.Context) {
	body := user.User{}
	if err := c.BindJSON(&body); err != nil {
		log.Fatal(err)
	}
	id, err := uc.us.AddUser(body)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	body.Id = id
	c.JSON(http.StatusAccepted, &body)
}
