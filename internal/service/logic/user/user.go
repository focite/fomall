package user

type sUser struct {
}

func init() {
	//service.RegisterAdmin(NewUser())
}

func NewUser() *sUser {
	return &sUser{}
}

func (s *sUser) Name(id int) int {
	return id
}
