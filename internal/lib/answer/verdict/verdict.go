package verdict

const (
	OK       = Type("OK")
	EMPTY    = Type("")
	WA       = Type("WA")
	WAIT     = Type("WAIT")
	REVIEWED = Type("REVIEWED")
)

type (
	Type string

	Verdict struct {
		Verdict Type `json:"verdict" enums:"OK,WA,WAIT,"`
	}
)
