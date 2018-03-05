// clockall get timezone's url from args and echo the time
package main

import (
	"fmt"
	"os"
	"strings"
)

type tz string

type Clock struct {
  host string
}

func main() {
	args := os.Args[:]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: clockall TIMEZONE=host:port")
		os.Exit(1)
	}
	tzClock := make(map[tz]Clock)
	for _, arg := range args[1:] {
		s := strings.Split(arg, "=")
		tzClock[s[0]] = lock{url: }
	}
}
