// cmd flag to print SHA#*$ or SHA512, default SHA256
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var n = flag.String("n", "sha256", "[sha256(default) | sha384 | sha512]")

func main() {
	flag.Parse()
	for _, s := range flag.Args() {
		fmt.Printf("%s:\t%x\n", s, sha(s, *n))
	}
}

func sha(s string, sn string) []byte {
	switch sn {
	case "sha256":
		sum := sha256.Sum256([]byte(s))
		return sum[:]
	case "sha384":
		sum := sha512.Sum384([]byte(s))
		return sum[:]
	case "sha512":
		sum := sha512.Sum512([]byte(s))
		return sum[:]
	}
	return nil
}
