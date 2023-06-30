package entity

const (
	VerdictOK    = Verdict("OK")
	VerdictEMPTY = Verdict("")
	VerdictWA    = Verdict("WA")
)

type (
	Verdict string

	VerdictDto struct {
		Verdict Verdict `json:"verdict"`
	}
)
