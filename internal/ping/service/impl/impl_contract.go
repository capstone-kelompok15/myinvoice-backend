package impl

type service struct {
}

type Service struct {
}

func NewPingService(params Service) *service {
	return &service{}
}
