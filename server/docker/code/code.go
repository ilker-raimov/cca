package code

type Code interface {
	Compile(code string) (bool, error)
	Test(code string, tests []any)
}
