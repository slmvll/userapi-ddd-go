package user

import (
	"log"
	"net/http"

	"userapi-ddd-go/pkg/domain/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary get a user record by ID
// @ID get-user-by-id
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} user.User
// @Failure 404 {object} object
// @Router /user/{id} [get]
func (uc UserController) GetUserById(c *gin.Context) {
	uri := URI{}

	// binding to URI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := uuid.Parse(uri.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resUser, err := uc.us.GetUser(id)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusOK, resUser)
}
