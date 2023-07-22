package updatepass

type Request struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
