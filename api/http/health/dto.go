package health

type HealthStatus string

const (
	HealthOk    HealthStatus = "OK"
	HealthError HealthStatus = "ERROR"
)

type HealthCheck struct {
	Status HealthStatus `json:"status"`
}
