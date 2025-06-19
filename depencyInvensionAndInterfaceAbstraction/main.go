package depencyinvensionandinterfaceabstraction

import (
	"database/sql"
	"fmt"
)

type Custom_Repo interface {
	CreateUser(userId string) bool
}

type UserService struct {
	repo Custom_Repo
}

func (servcie *UserService) RegisterUser(userId string) bool {
	return servcie.repo.CreateUser(userId)
}

type Postgres_Repo struct {
	DbConn *sql.DB
}

func (rds *Postgres_Repo) CreateUser(userId string) bool {
	fmt.Println("000tester db process: ", rds.DbConn, userId)
	return true
}

func main() {
	rdsRepo := &Postgres_Repo{}
	service := &UserService{
		repo: rdsRepo,
	}

	testUserId := "123"
	service.RegisterUser(testUserId)
}
