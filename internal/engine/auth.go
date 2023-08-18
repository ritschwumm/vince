package engine

import (
	"net"

	"github.com/dgraph-io/badger/v4"
	"github.com/dolthub/vitess/go/mysql"
	querypb "github.com/dolthub/vitess/go/vt/proto/query"
	"github.com/vinceanalytics/vince/internal/tokens"
)

type Auth struct {
	DB *badger.DB
}

var _ mysql.AuthServer = (*Auth)(nil)

func (a Auth) Salt() ([]byte, error) {
	return mysql.NewSalt()
}

func (a Auth) ValidateHash(salt []byte, user string, authResponse []byte, remoteAddr net.Addr) (mysql.Getter, error) {
	return &mysql.StaticUserData{}, mysql.NewSQLError(mysql.ERAccessDeniedError, mysql.SSAccessDeniedError, "Access denied for user '%v'", user)
}

func (a Auth) AuthMethod(user, addr string) (string, error) {
	return mysql.MysqlClearPassword, nil
}

func (a Auth) Negotiate(c *mysql.Conn, user string, remoteAddr net.Addr) (mysql.Getter, error) {
	// Finish the negotiation.
	password, err := mysql.AuthServerNegotiateClearOrDialog(c, mysql.MysqlClearPassword)
	if err != nil {
		return nil, err
	}
	if !tokens.Valid(a.DB, password) {
		return &mysql.StaticUserData{}, mysql.NewSQLError(mysql.ERAccessDeniedError, mysql.SSAccessDeniedError, "Access denied for user '%v'", user)
	}
	return StaticUserData(user), nil
}

type StaticUserData string

// Get returns the wrapped username and groups
func (sud StaticUserData) Get() *querypb.VTGateCallerID {
	return &querypb.VTGateCallerID{Username: string(sud)}
}
