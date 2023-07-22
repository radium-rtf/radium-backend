package verdict

const (
	OK    = Type("OK")
	EMPTY = Type("")
	WA    = Type("WA")
	WAIT  = Type("WAIT")
)

type (
	Type string

	Verdict struct {
		Verdict Type `json:"verdict"`
	}
)
