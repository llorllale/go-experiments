package nil

// Blog post: https://llorllale.github.io/golang-methods-on-nil-references/

// You can call methods on nil references.
// This is one way of implementing the Null Object pattern in go.
// In the struct's methods, provide alternative behaviour if
// the reference to the receiver is nil.
// One interesting thing is that you can't simply declare
// a variable of a type and assign nil to it - the compiler
// won't allow it ("cannot assign nil to variable to type Person").
// You can't even try casting nil to Person, because then the
// compiler will complain about the "use of untyped nil".

type Person interface {
	Name() string
}

func GetPerson(name string) Person {
	return nil
}

type person struct {
}

// Name returns this person's name
func (p *person) Name() string {
	return "person found"
}

type NotFound struct {
}

func (NotFound) Name() string {
	panic("person not found")
}
