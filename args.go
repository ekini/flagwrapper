package flagwrapper

import (
	"fmt"
	"strings"
)

type Arg struct {
	names        []string // names of the arg, can be either short or long and can have multiple
	passThrough  bool     // if it has a parameter, so the next arg needs to be consumed
	hasParameter bool     // if we want to passthrough this arg to the program that we are wrapping
}

func (a *Arg) HasParameter() bool { return a.hasParameter }
func (a *Arg) Names() []string    { return a.names }
func (a *Arg) PassThrough() bool  { return a.passThrough }

func NewSimpleArg(names ...string) *Arg {
	if len(names) == 0 {
		panic("need to pass at least one option to the NewParameterArg")
	}
	return &Arg{names: names, hasParameter: false, passThrough: false}
}

func NewSimplePassThroughArg(names ...string) *Arg {
	if len(names) == 0 {
		panic("need to pass at least one option to the NewSimplePassThroughArg")
	}
	return &Arg{names: names, hasParameter: false, passThrough: true}
}

func NewParameterArg(names ...string) *Arg {
	if len(names) == 0 {
		panic("need to pass at least one option to the NewParameterArg")
	}
	return &Arg{names: names, hasParameter: true, passThrough: false}
}

func NewParameterPassThroughArg(names ...string) *Arg {
	if len(names) == 0 {
		panic("need to pass at least one option to the NewParameterPassThroughArg")
	}
	return &Arg{names: names, hasParameter: true, passThrough: true}
}

func ParseArgs(args []string, wrapperArgs []*Arg) (parsedWrapperArgs []string, restArgs []string, err error) {
	// make a map of args for faster lookup
	wrapperArgsMap := make(map[string]*Arg)
	for _, arg := range wrapperArgs {
		arg := arg
		for _, optName := range arg.Names() {
			wrapperArgsMap[optName] = arg
		}
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if a, ok := wrapperArgsMap[arg]; ok {
			parsedWrapperArgs = append(parsedWrapperArgs, arg)
			if a.PassThrough() {
				restArgs = append(restArgs, arg)
			}
			if a.HasParameter() {
				if i+1 == len(args) || strings.HasPrefix(args[i+1], "-") {
					err = fmt.Errorf("argument %s requires a value", arg)
					return
				} else {
					parsedWrapperArgs = append(parsedWrapperArgs, args[i+1])
					if a.PassThrough() {
						restArgs = append(restArgs, args[i+1])
					}
					i++
				}
			}
		} else {
			restArgs = append(restArgs, arg)
		}
	}
	return
}
