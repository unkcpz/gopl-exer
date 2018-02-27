package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

// randomPalindrome returns a palindrome whose length and
// contents are derived from the random number generator rng
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(runes)
}

//
func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23) + 2
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		rn := rng.Intn(len(letters))
		rnAux := rng.Intn(len(letters))
		runes[i] = rn
		runes[n-i-1] = runes[i] + 1
	}
	return string(runes)
}

func TestIsPalindrome(t *testing.T) {
	// Initialize a pseudo-random number generator
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf(`IsPalindrome(%q) = false`, p)
		}
	}
}

func TestNonIsPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		np := randomNonPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf(`IsPalindrome(%q) = true`, np)
		}
	}
}

var letters []rune

func init() {
	//just stick to ASCII
	for r := rune(0x21); r < 0x7e; r++ {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		}
	}
}
