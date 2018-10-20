package interfaces

// Assume that interface 'Person' was defined in some other package/file
type Person interface {
	Name() string
	Email() string
}

// While reading this code, how would I know that this 'person' struct
// implements the 'Person' interface?
// The answer seems to be to use the 'guru' tool. So now, instead of
// being able to tell with my own eyes, I have to rely on an external
// tool.
type person struct {
}

func (p *person) Name() string {
	return "Bob"
}

func (p *person) Email() string {
	return "bob@gmail.com"
}
