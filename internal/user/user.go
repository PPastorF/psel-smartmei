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
	cc := c.(*misc.CustomContext)
	params := &GetUserRequest{misc.UniqueID(c.Param("id"))}
		
	u := new(User)
	err := cc.DB.Model(u).Where("id = ?", params.ID).Select()
	if err != nil {
		log.Error().Err(err).Msg("Erro buscando usuario")
		return echo.NewHTTPError(http.StatusNotFound)
	}
	
	log.Info().Interface("data", u).Msg("Usuario buscado com sucesso")

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
	cc := c.(*misc.CustomContext)

	rawParams := new(CreateUserRequest)
	err := c.Bind(rawParams); 
	if err != nil {
		log.Fatal().Err(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	
	params, err := rawParams.SanitizeAndValidate()
	if err != nil {
		log.Error().Err(err).Msg("Erro ao criar novo usuario")
		return echo.NewHTTPError(http.StatusBadRequest, "Email ou nome invalido(s)")
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
		log.Error().Err(err).Msg("Erro ao criar novo usuario")

		// Caso o email não seja unico
		pgErr, ok := err.(pg.Error)
		if ok && pgErr.IntegrityViolation() {
			return echo.NewHTTPError(http.StatusBadRequest, "Email ja cadastrado")
		}
		
		return err
	}

	log.Info().Interface("data", u).Msg("Novo usuario criado com sucesso")

	return c.JSON(http.StatusCreated, u)
}
