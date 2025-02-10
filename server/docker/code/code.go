package code

import (
	"os"

	"github.com/ilker-raimov/cca/common/setup"
)

type Code interface {
	Compile(code string) (bool, string, error)
	Test(code string, tests []any)
	Cleanup(id string)
}

var instance Code

func Init() {
	arg_count := len(os.Args)

	if arg_count != 2 {
		panic("Invalid argument count.")
	}

	arg := os.Args[1]
	language := setup.From(arg)

	switch language {
	case setup.JAVA_8, setup.JAVA_11, setup.JAVA_17, setup.JAVA_21:
		version, err := language.GetVersion()

		if err != nil {
			panic(err)
		}

		instance = NewJava(version)
	default:
		panic("Unkown language.")
	}
}

func GetInstance() Code {
	return instance
}
