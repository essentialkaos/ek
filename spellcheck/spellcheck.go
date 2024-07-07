// Package spellcheck provides spellcheck based on Damerau–Levenshtein distance algorithm
package spellcheck

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sort"
	"strings"

	"github.com/essentialkaos/ek/v13/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Model is spellcheck model struct
type Model struct {
	Threshold int // Score threshold (default: 2)

	terms []string
}

// ////////////////////////////////////////////////////////////////////////////////// //

type suggestItem struct {
	term  string
	score int
}

type suggestItems []*suggestItem

func (s suggestItems) Len() int {
	return len(s)
}

func (s suggestItems) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s suggestItems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Train trains words by given string slice
func Train(words []string) *Model {
	model := &Model{Threshold: 2}

	if len(words) == 0 {
		return model
	}

	sm := make(map[string]bool)

	for _, w := range words {
		sm[w] = true
	}

	for cw := range sm {
		model.terms = append(model.terms, cw)
	}

	return model
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Correct corrects given value
func (m *Model) Correct(word string) string {
	if m == nil || len(m.terms) == 0 {
		return word
	}

	var result *suggestItem

	for _, si := range getSuggestSlice(m.terms, word) {
		if result == nil {
			result = si
			continue
		}

		if si.score < result.score {
			result = si
			continue
		}
	}

	if result.score > mathutil.Between(m.Threshold, 1, 1000) {
		return word
	}

	return result.term
}

// Suggest suggests words for given word or word part
func (m *Model) Suggest(word string, max int) []string {
	if m == nil || len(m.terms) == 0 {
		return []string{word}
	}

	if max == 1 {
		return []string{m.Correct(word)}
	}

	sis := getSuggestSlice(m.terms, word)

	sort.Sort(sis)

	var result []string

	for i := 0; i < mathutil.Between(max, 1, len(sis)); i++ {
		result = append(result, sis[i].term)
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// I don't have an idea how we could separate this method
// codebeat:disable[LOC,ABC,CYCLO]

// Damerau–Levenshtein distance algorithm and code
func getDLDistance(source, target string) int {
	sl, tl := len(source), len(target)

	if sl == 0 {
		if tl == 0 {
			return 0
		}
		return tl
	} else if tl == 0 {
		return sl
	}

	h := make([][]int, sl+2)

	for i := range h {
		h[i] = make([]int, tl+2)
	}

	ll := sl + tl

	h[0][0] = ll

	for i := 0; i <= sl; i++ {
		h[i+1][0] = ll
		h[i+1][1] = i
	}

	for j := 0; j <= tl; j++ {
		h[0][j+1] = ll
		h[1][j+1] = j
	}

	sd := make(map[rune]int)

	for _, rn := range source + target {
		sd[rn] = 0
	}

	for i := 1; i <= sl; i++ {
		d := 0

		for j := 1; j <= tl; j++ {
			i1 := sd[rune(target[j-1])]
			j1 := d

			if source[i-1] == target[j-1] {
				h[i+1][j+1] = h[i][j]
				d = j
			} else {
				h[i+1][j+1] = mathutil.Min(h[i][j], mathutil.Min(h[i+1][j], h[i][j+1])) + 1
			}

			h[i+1][j+1] = mathutil.Min(h[i+1][j+1], h[i1][j1]+(i-i1-1)+1+(j-j1-1))
		}

		sd[rune(source[i-1])] = i
	}

	return h[sl+1][tl+1]
}

// codebeat:enable[LOC,ABC,CYCLO]

func getSuggestSlice(terms []string, word string) suggestItems {
	var result suggestItems

	for _, t := range terms {
		result = append(result, &suggestItem{t, getDLDistance(strings.ToLower(t), strings.ToLower(word))})
	}

	return result
}
