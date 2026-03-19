package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	dbusermod "github.com/web-rabis/db/user"
	dbuserman "github.com/web-rabis/db/user/manager/user"
)

type Manager struct {
	userMan dbuserman.IManager
}

func NewUserManager(userMan dbuserman.IManager, authMan *auth.Manager) *Manager {
	return &Manager{
		userMan: userMan,
	}
}

func (m *Manager) SignIn(ctx context.Context, username, password string) (*dbusermod.User, error) {
	user, err := m.userMan.ByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	passwordHash := md5Hash(password)
	if strings.ToLower(passwordHash) != strings.ToLower(user.Password) || user.State != "ACTIVE" {
		return nil, errors.New("wrong username or password")
	}
	return user, nil
}
func (m *Manager) ById(ctx context.Context, id int) (*dbusermod.User, error) {
	return m.userMan.ById(ctx, id)
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))      // md5.Sum returns [16]byte
	return hex.EncodeToString(hash[:]) // Convert to hex string
}
