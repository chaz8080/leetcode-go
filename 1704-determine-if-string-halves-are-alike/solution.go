package determineIfStringHalvesAreAlike

import (
	"fmt"
)

func halvesAreAlike(s string) bool {
	m := make(map[string]bool)
	m["a"] = true
	m["e"] = true
	m["i"] = true
	m["o"] = true
	m["u"] = true
	m["A"] = true
	m["E"] = true
	m["I"] = true
	m["O"] = true
	m["U"] = true

	var av, bv int
	bs := []byte(s)

	a, b := bs[0:len(bs)/2], bs[(len(bs)/2):]

	fmt.Println(string(a), string(b))
	for i := 0; i < len(bs)/2; i++ {
		if m[string(a[i])] {
			av++
		}
		if m[string(b[i])] {
			bv++
		}
	}

	return av == bv
}
