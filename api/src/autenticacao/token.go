package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	//criamos o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	//assinamos o token
	return token.SignedString([]byte(config.SecretKey))
}
