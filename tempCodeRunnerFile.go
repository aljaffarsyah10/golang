Before sync.Mutex
type counter struct {
    val int
}

func (c *counter) Add(int) {
    c.val++
}

func (c *counter) Value() (int) {
    return c.val
}
