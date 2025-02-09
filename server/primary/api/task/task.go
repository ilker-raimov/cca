package task

type Example struct {
	Input  string
	Output string
}

type Task struct {
	Name         string
	Description  string
	Examples     []Example
	Restrictions []string
	Time         int16
}
