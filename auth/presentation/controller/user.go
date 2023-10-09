package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/architecture-template/echo-ddd/auth/service"
	"github.com/architecture-template/echo-ddd/auth/presentation/output"
	"github.com/architecture-template/echo-ddd/auth/presentation/response"
	"github.com/architecture-template/echo-ddd/auth/presentation/parameter"
)

type UserController interface {
	RegisterUser() echo.HandlerFunc
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
    return &userController{
        userService: userService,
    }
}

// @tags        User
// @Summary     ユーザー登録
// @Accept      json
// @Produce     json
// @Router      /user/register_user [post]
// @Param       body body parameter.RegisterUser true "ユーザー登録"
// @Success     200  {object} response.Success{items=output.User}
// @Failure     500  {array}  output.Error
func (u *userController) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := &parameter.RegisterUser{}
		c.Bind(param)

		// 登録済のメールアドレスを検索
		check, _ := u.userService.FindByEmail(param.Email)
		if check != nil {
			response := response.ErrorWith("register_user", 400, "email already exists")

			return c.JSON(400, response)
		}

		result, err := u.userService.RegisterUser(param)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("register_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToUser(result)
		response := response.SuccessWith("register_user", 200, out)

		return c.JSON(200, response)
	}
}
