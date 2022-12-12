package seller

type sSeller struct {
}

func init() {
	//service.RegisterSeller(NewSeller())
}

func NewSeller() *sSeller {
	return &sSeller{}
}

func (s *sSeller) Name(id int) int {
	return id
}
