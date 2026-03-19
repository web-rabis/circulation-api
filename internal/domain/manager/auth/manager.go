package auth

import (
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

type Manager struct {
	jwtKey          []byte
	TokenTTL        time.Duration
	RefreshTokenTTL time.Duration

	isTesting bool
}

func NewAuthManager(jwtKey []byte, tokenTTL, refreshTokenTTL time.Duration) *Manager {
	return &Manager{
		jwtKey:          jwtKey,
		TokenTTL:        tokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
	}
}

func (m *Manager) Testing() {
	m.isTesting = true
}

func (m *Manager) JWTKey() []byte {
	return m.jwtKey
}

func (m *Manager) JWTAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", m.jwtKey, nil)
}

func (m Manager) NewRefreshToken(userId string) (string, error) {

	claims := map[string]interface{}{
		"id":              userId,
		"is_refresh":      true,
		jwt.ExpirationKey: time.Now().Add(m.RefreshTokenTTL).Unix(),
	}

	_, tokenString, err := m.JWTAuth().Encode(claims)

	return tokenString, err
}

func (m Manager) NewAccessToken(userId int64) (string, error) {
	claims := map[string]interface{}{
		"id":              userId,
		jwt.ExpirationKey: time.Now().Add(m.TokenTTL).Unix(),
	}

	_, tokenString, err := m.JWTAuth().Encode(claims)

	return tokenString, err

}
