package depencyinvensionandinterfaceabstraction

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func Run() {
	rdsRepo := &Postgres_Repo{}
	service := &UserService{
		repo: rdsRepo,
	}

	testUserId := "123"
	service.RegisterUser(testUserId)

	r, _ := os.Open("data.csv")
	NewTransactionFromCsv(r)
}

func NewTransactionFromCsv(r io.Reader) {
	reader := csv.NewReader(r)
	_, err := reader.Read()
	if err != nil {
		fmt.Println("reader error: ", err)
		return
	}

	var txns []interface{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("reader error: ", err)
			return
		}

		txns = append(txns, record)

		test, _ := json.Marshal(txns)
		fmt.Println("000tester: ", string(test))
	}
}
