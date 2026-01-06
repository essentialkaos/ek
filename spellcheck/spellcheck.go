// Package spellcheck provides spellcheck based on Damerau–Levenshtein distance algorithm
package spellcheck

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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

// suggestItem is struct for storing suggestion item
type suggestItem struct {
	term  string
	score int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Train trains spellcheck model with given words
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

// Distance calculates Damerau–Levenshtein distance between two strings
func Distance(source, target string) int {
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
				h[i+1][j+1] = min(h[i][j], min(h[i+1][j], h[i][j+1])) + 1
			}

			h[i+1][j+1] = min(h[i+1][j+1], h[i1][j1]+(i-i1-1)+1+(j-j1-1))
		}

		sd[rune(source[i-1])] = i
	}

	return h[sl+1][tl+1]
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Correct returns corrected word for given word
func (m *Model) Correct(word string) string {
	if m == nil || len(m.terms) == 0 {
		return word
	}

	result := suggestItem{score: 99999999999}

	for _, si := range getSuggestSlice(m.terms, word) {
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

// Suggest returns suggestions for given word
func (m *Model) Suggest(word string, max int) []string {
	if m == nil || len(m.terms) == 0 {
		return []string{word}
	}

	if max == 1 {
		return []string{m.Correct(word)}
	}

	sis := getSuggestSlice(m.terms, word)

	sort.Slice(sis, func(i, j int) bool {
		return sis[i].score < sis[j].score
	})

	var result []string

	for i := range mathutil.Between(max, 1, len(sis)) {
		result = append(result, sis[i].term)
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getSuggestSlice returns slice of suggestItem for given terms and word
func getSuggestSlice(terms []string, word string) []suggestItem {
	var result []suggestItem

	for _, t := range terms {
		result = append(result, suggestItem{t, Distance(strings.ToLower(t), strings.ToLower(word))})
	}

	return result
}
