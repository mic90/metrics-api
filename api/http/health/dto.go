package health

type Status string

const (
	OK    Status = "OK"
	ERROR Status = "ERROR"
)

type StatusResponse struct {
	Status Status `json:"status"`
}
