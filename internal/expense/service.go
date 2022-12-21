package expense

type Service struct {
	Repo Repo
}

// Repo Expense ServiceUseCase
type Repo interface {
	InsertExpense(req Request) (int, error)
}

func NewService(repo Repo) ServiceUseCase {
	return &Service{
		Repo: repo,
	}
}

func (s Service) AddExpense(req Request) (int, error) {
	return s.Repo.InsertExpense(req)
}
