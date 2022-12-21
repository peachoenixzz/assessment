package expense

import "fmt"

type Service interface {
	AddExpense(stdNme string) error
}

type ServiceExpense struct {
	Repo Repo
}

// Repo Expense Service
type Repo interface {
	InsertExpense(stdNme string) (string, error)
}

func NewService(repo Repo) Service {
	return &ServiceExpense{
		Repo: repo,
	}
}

func (s ServiceExpense) AddExpense(stdNme string) error {
	fmt.Println("registering student: ", s)
	_, err := s.Repo.InsertExpense(stdNme)
	if err != nil {
		return err
	}
	fmt.Println("inserted with id: ", stdNme)
	return nil
}
