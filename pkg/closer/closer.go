package closer

import (
	"fmt"
	"strings"
	"sync"
)

type Closer struct {
	funcs []func() error
	mu    sync.Mutex
}

func New() Closer {
	return Closer{}
}

func (c *Closer) Add(f func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(len(c.funcs))

	var msgs []string
	for _, f := range c.funcs {
		if err := f(); err != nil {
			msgs = append(msgs, fmt.Sprintf("[!] %v", err))
		}
	}

	return fmt.Errorf(
		"shutdown finished with error(s): \n%s",
		strings.Join(msgs, "\n"),
	)
}
