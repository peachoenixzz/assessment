package expense

import "fmt"

type Service interface {
	AddExpense(req Request) (int, error)
}

type ServiceExpense struct {
	Repo Repo
}

// Repo Expense Service
type Repo interface {
	InsertExpense(req Request) (int, error)
}

func NewService(repo Repo) Service {
	return &ServiceExpense{
		Repo: repo,
	}
}

func (s ServiceExpense) AddExpense(req Request) (int, error) {
	fmt.Println("registering student: ", s)
	id, err := s.Repo.InsertExpense(req)
	if err != nil {
		return 0, err
	}
	fmt.Println("inserted with id: ", req)
	return id, nil
}
