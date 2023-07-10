package entity

const (
	VerdictOK    = Verdict("OK")
	VerdictEMPTY = Verdict("")
	VerdictWA    = Verdict("WA")
	VerdictWAIT  = Verdict("WAIT")
)

type (
	Verdict string

	VerdictDto struct {
		Verdict Verdict `json:"verdict"`
	}
)
