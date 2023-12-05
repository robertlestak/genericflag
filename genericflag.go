package genericflag

import (
	"strings"
)

// FlagSet is a generic flag parser
// whereas the standard library flag package requires flags to be defined,
// this package will parse any flag that is provided and return the remainder
// of the args that were not parsed into flags
type FlagSet struct {
	// Name is the name of the flag set
	Name string
	// Expected is a slice of flags that are expected to be parsed
	Expected []string
	// Flags is a map of flags to their values
	Flags map[string][]string
	// args is a slice of args that were not parsed into flags
	args []string
}

// NewFlagSet returns a new FlagSet with the provided name
func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		Name:  name,
		Flags: make(map[string][]string),
	}
}

func quoteValWithSpaces(val string) string {
	if strings.Contains(val, " ") && !strings.HasPrefix(val, "\"") && !strings.HasPrefix(val, "'") {
		val = "\"" + val + "\""
	}
	return val
}

// Parse loops through all of the args and parses them into flags
// if an Expected slice is provided, then only flags in that slice will be parsed
// and the remainder will be added to the args slice
func (f *FlagSet) Parse(args []string) error {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		toParse := true
		if arg == "" {
			continue
		}
		if arg[0] != '-' {
			toParse = false
		}
		flagName := arg[1:]
		flagVal := ""
		if flagName == "" {
			f.args = append(f.args, arg)
			continue
		}
		if flagName[0] == '-' {
			flagName = flagName[1:]
		}
		if strings.Contains(flagName, "=") {
			split := strings.Split(flagName, "=")
			flagName = split[0]
			flagVal = strings.Join(split[1:], "=")
		} else {
			if i+1 < len(args) {
				if args[i+1][0] != '-' {
					flagVal = args[i+1]
					i++
				}
			}
		}
		if f.Expected != nil && len(f.Expected) > 0 {
			toParse = false
			for _, expected := range f.Expected {
				if flagName == expected {
					toParse = true
				}
			}
		} else {
			toParse = true
		}
		if !toParse {
			f.args = append(f.args, arg)
			if flagVal != "" && !strings.Contains(arg, flagVal) {
				flagVal = quoteValWithSpaces(flagVal)
				f.args = append(f.args, flagVal)
			}
			continue
		} else {
			flagVal = quoteValWithSpaces(flagVal)
		}
		f.Flags[flagName] = append(f.Flags[flagName], flagVal)
	}
	return nil
}

// Args returns the args slice that was not parsed into flags
func (f *FlagSet) Args() []string {
	return f.args
}
