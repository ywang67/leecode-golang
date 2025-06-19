package depencyinvensionandinterfaceabstraction

import (
	"fmt"
	"testing"
)

type Mock_Repo struct{}

func (mRepo *Mock_Repo) CreateUser(userId string) bool {
	fmt.Println("mocking respo testing")
	return true
}

func TestMockRepo(t *testing.T) {
	service := &UserService{
		repo: &Mock_Repo{},
	}

	result := service.repo.CreateUser("testingUser123")

	t.Run("mock repo testing: ", func(t *testing.T) {
		if result {
			t.Log("mocing tesitng passed")
		} else {
			t.Fatal("mocing tesitng failed")
		}
	})

}
