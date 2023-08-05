package updatepass

type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
