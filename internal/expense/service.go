package expense

type Service interface {
	InsertExpense(stdNme string) error
}

type ServiceExpense struct {
	Repo Repo
}

// Repo Expense Service
type Repo interface {
}

func (s ServiceExpense) InsertExpense(stdNme string) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo Repo) Service {
	return &ServiceExpense{
		Repo: repo,
	}
}
