package user

import (
	"time"
	"net/http"
	"github.com/rs/zerolog/log"
	"github.com/labstack/echo/v4"
	"github.com/go-pg/pg/v10"
	
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

// Res:
// { ...User }
func GetUser(c echo.Context) error {
	errMsg  := "Erro buscando usuario"
	succMsg := "Usuario buscado com sucesso"
	cc := c.(*misc.CustomContext)
	rawParams := &GetUserRequest{misc.UniqueID(c.Param("id"))}
		
	params, err := rawParams.Sanitize()
	if err != nil {
		log.Error().Interface("params", params).Err(err).Msg(errMsg)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Se não for um id válido, nem busca no banco mas retorna 404
	err = params.Validate()
	if err != nil {
		log.Error().Interface("params", params).Err(err).Msg(errMsg)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	u := new(User)
	err = cc.DB.Model(u).Where("id = ?", params.UserID).Select()
	if err != nil {
		log.Error().Interface("params", params).Err(err).Msg(errMsg)
		return echo.NewHTTPError(http.StatusNotFound)
	}
	
	log.Info().Interface("data", u).Msg(succMsg)
	return c.JSON(http.StatusOK, u)
}

// Req:
// {
// 	"name": string,
// 	"email": string
// }
// Res:
// { ...User }
func CreateUser(c echo.Context) error {
	errMsg  := "Erro ao criar novo usuario"
	succMsg := "Novo usuario criado com sucesso"
	cc := c.(*misc.CustomContext)
	rawParams := new(CreateUserRequest)

	err := c.Bind(rawParams); 
	if err != nil {
		log.Fatal().Err(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	
	params, err := rawParams.Sanitize()
	if err != nil {
		log.Error().Interface("params", params).Err(err).Msg(errMsg)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = params.Validate()
	if err != nil {
		log.Error().Interface("params", params).Err(err).Msg(errMsg)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	u := &User{
		ID: misc.GenerateUniqueID(),
		Name: params.Name,
		Email: params.Email,
		CreatedAt: time.Now(),
		Collection: make([]misc.UniqueID, 0),
		Lent: make([]misc.UniqueID, 0),
		Borrowed: make([]misc.UniqueID, 0),
	}

	_, err = cc.DB.Model(u).Insert()
	if err != nil {
		log.Error().Interface("params", params).Err(err).Msg(errMsg)

		pgErr, ok := err.(pg.Error)
		if ok && pgErr.IntegrityViolation() {
			return echo.NewHTTPError(http.StatusBadRequest, "Email ja cadastrado")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Info().Interface("data", u).Msg(succMsg)
	return c.JSON(http.StatusCreated, u)
}
