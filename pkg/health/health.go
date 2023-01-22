package health

type Health struct {
	ServiceName string                 `json:"serviceName"`
	Checks      map[string]CheckResult `json:"checks"`
	Unhealthy   bool                   `json:"-"`
}

func CreateResponse(serviceName string) Health {
	return Health{
		ServiceName: serviceName,
		Checks:      map[string]CheckResult{},
	}
}

func (s *Health) AddResult(result CheckResult) {
	if s.Checks == nil {
		s.Checks = map[string]CheckResult{}
	}

	s.Checks[result.CheckName] = result

	if result.Status == Status_Unhealthy {
		s.Unhealthy = true
	}
}
