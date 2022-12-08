package textjustification

import (
	"fmt"
	"log"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	// create map that contains the word that we want to store for each line
	m := make(map[int]string)
	line := 0
	wordIndex := 0
	for wordIndex < len(words) {
		lLen := len(m[line])
		nextWordLen := len(words[wordIndex])
		if lLen+nextWordLen <= maxWidth {
			m[line] = m[line] + fmt.Sprintf("%s0", words[wordIndex])
			wordIndex++
		} else {
			m[line] = strings.TrimSuffix(m[line], "0")
			line++
		}
	}
	log.Println(m)
	m[line] = strings.TrimSuffix(m[line], "0")
	return justifyLines(m, maxWidth)
}

func justifyLines(m map[int]string, maxWidth int) []string {
	ret := []string{}
	line := 0
	for line < len(m) {
		words := strings.Split(m[line], "0")
		totalSize := getTotalSize(words)
		spacesRequired := maxWidth - totalSize
		spacesInserted := 0
		if len(words) == 1 {
			for spacesInserted < spacesRequired {
				words[0] += " "
				spacesInserted++
			}
		} else if line == len(m)-1 { // last line
			for spacesInserted < spacesRequired {
				for i := range words {
					if i == len(words)-1 {
						continue
					}
					words[i] += " "
					spacesInserted++
				}
				for spacesInserted < spacesRequired {
					words[len(words)-1] += " "
					spacesInserted++
				}
			}
		} else {
			for spacesInserted < spacesRequired {
				w := spacesInserted % (len(words) - 1)
				words[w] += " "
				spacesInserted++
			}
		}

		ret = append(ret, strings.Join(words, ""))
		line++
	}

	return ret
}

// This    is    an

func getTotalSize(words []string) int {
	sum := 0
	for _, word := range words {
		sum += len(word)
	}
	return sum
}
