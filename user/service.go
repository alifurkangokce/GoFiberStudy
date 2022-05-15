package user

type Service interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo: repo}
}

var _ Service = service{}

// Create implements Service
func (s service) Create(model Model) (uint, error) {
	return s.repo.Create(model)
}

// Get implements Service
func (s service) Get(id uint) (*Model, error) {
	return s.repo.Get(id)
}
