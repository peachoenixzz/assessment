package expense

type Service struct {
	Repo Repo
}

// Repo Expense ServiceUseCase
type Repo interface {
	InsertExpense(req Request) (int, error)
	GetExpense(id string) (Response, error)
}

func NewService(repo Repo) ServiceUseCase {
	return &Service{
		Repo: repo,
	}
}

func (s Service) AddExpense(req Request) (int, error) {
	return s.Repo.InsertExpense(req)
}

func (s Service) ViewExpense(id string) (Response, error) {
	return s.Repo.GetExpense(id)
}
