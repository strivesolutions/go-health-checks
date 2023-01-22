package health

type Status string

const (
	Status_Ok        Status = "ok"
	Status_Unhealthy Status = "unhealthy"
)
