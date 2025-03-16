package service

type PingService struct {
}

func NewPingService() *PingService {
	return &PingService{}
}

type PingResponse struct {
	Ping string `json:"ping"`
}

func (s *PingService) Ping() (PingResponse, error) {
	return PingResponse{"pong"}, nil
}
