package services

type PinetService struct{}

func (p *PinetService) Status() string {
	return "up"
}
