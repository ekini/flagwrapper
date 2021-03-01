package flagwrapper

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	args := []string{"--kms-key", "xxxxx", "--help", "-x", "--user", "tester", "--noop"}
	wrapperArgs := []*Arg{
		NewParameterArg("--kms-key"),
		NewSimpleArg("-x"),
		NewSimplePassThroughArg("-h", "--help"),
		NewParameterPassThroughArg("--user"),
	}
	parsedWrapperArgs, rest, err := ParseArgs(args, wrapperArgs)
	if err != nil {
		t.Errorf("got unexpected error while parsing: %s", err)
	} else {
		if parsedWrapperArgs[0] != "--kms-key" || parsedWrapperArgs[1] != "xxxxx" {
			t.Errorf("Parsed args: %#v, expected [\"--kms-key\", \"xxxxx\"] in the wrappedArgs", parsedWrapperArgs)
		}
		if parsedWrapperArgs[2] != "--help" || rest[0] != "--help" {
			t.Errorf("Parsed args: %#v, expected [..., \"--help\"] in the wrappedArgs and [\"--help\", ...] in rest", parsedWrapperArgs)
		}
		if parsedWrapperArgs[3] != "-x" {
			t.Errorf("Parsed args: %#v, expected [..., \"-x\"] in the wrappedArgs", parsedWrapperArgs)
		}
		if parsedWrapperArgs[4] != "--user" || parsedWrapperArgs[5] != "tester" {
			t.Errorf("Parsed args: %#v, expected [..., \"--user\", \"tester\"] in the wrappedArgs", parsedWrapperArgs)
		}
		if rest[1] != "--user" || rest[2] != "tester" {
			t.Errorf("Rest args: %#v, expected [..., \"--user\", \"tester\"] in the rest args", rest)
		}
		if rest[3] != "--noop" {
			t.Errorf("Rest args: %#v, expected [..., \"--noop\"] in the rest args", rest)
		}
	}
}
