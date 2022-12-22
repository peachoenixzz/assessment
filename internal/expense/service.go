package expense

type Service struct {
	Repo Repo
}

// Repo Expense ServiceUseCase
type Repo interface {
	InsertExpense(req Request) (int, error)
	GetExpenseByID(id string) (Response, error)
	GetExpense() (Response, error)
	UpdateExpenseByID(req Request, id string) (Response, error)
}

func NewService(repo Repo) ServiceUseCase {
	return &Service{
		Repo: repo,
	}
}

func (s Service) AddExpense(req Request) (int, error) {
	return s.Repo.InsertExpense(req)
}

func (s Service) ViewExpense() (Response, error) {
	return s.Repo.GetExpense()
}

func (s Service) ViewExpenseByID(id string) (Response, error) {
	return s.Repo.GetExpense(id)
}
