package repo

import (
	"context"
	"database/sql"
	"github.com/lynndotconfig/gkgo2/domain/user/entity"
	"github.com/lynndotconfig/gkgo2/internal/errcode"
	"github.com/lynndotconfig/gkgo2/pkg/gkgo2err"
)


type UserInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, *gkgo2err.Error)
}

type UserImpl struct {}

func NewUser() *UserImpl {
	return &UserImpl{}
}

func(i *UserImpl) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, *gkgo2err.Error) {
	db, err := sql.Open("mysql", "root@secret@tcp(127.0.0.1:3306)gkgo_db")
	defer db.Close()
	
	if err != nil {
		return nil, gkgo2err.New(errcode.SystemDBError, err, "open db fail, err=%s")
	}
	user := &entity.UserEntity{}
	err = db.QueryRow("SELECT * FROM user_tab WHERE email = ?", email).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, gkgo2err.New(errcode.UserEmailNotFound, email, "cannot find user by %s")
		}
		return nil,  gkgo2err.New(errcode.SystemDBError, err, "query user by email, err=%s")
	}
	
	return user, nil
}