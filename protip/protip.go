// Package protip provides methods for displaying usage tips
package protip

import (
	"math/rand"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil/panel"
	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Tip contains basic tip content
type Tip struct {
	Title   string // Tip title (required)
	Message string // Tip message (required)

	ColorTag string  // Custom tip color (optional)
	Weight   float64 // Custom tip weight (optional)
}

// Tips contains tips data
type Tips struct {
	data      []*Tip
	selectors []int
	maxWeight int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Probability is showing probability coefficient (0 ← less | more → 1)
var Probability = 0.25

// ColorTag is default panel color tag
var ColorTag = "{#75}"

// Options contains default panel options
var Options = panel.Options{panel.TOP_LINE, panel.BOTTOM_LINE}

// ////////////////////////////////////////////////////////////////////////////////// //

var collection *Tips

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds one or more tips to collection
func Add(tips ...*Tip) {
	if collection == nil {
		collection = &Tips{}
	}

	for _, tip := range tips {
		if tip == nil || tip.Title == "" ||
			tip.Message == "" || !fmtc.IsTag(tip.ColorTag) {
			continue
		}

		if tip.Weight == 0 {
			tip.Weight = 0.5
		}

		collection.data = append(collection.data, tip)
	}
}

// Show shows random tip if required
func Show(force bool) bool {
	if collection == nil || len(collection.data) == 0 {
		return false
	}

	if rand.Float64() > Probability && !force {
		return false
	}

	if len(collection.selectors) != len(collection.data) {
		collection.selectors = make([]int, len(collection.data))
		collection.maxWeight = 0

		for i, tip := range collection.data {
			collection.maxWeight += int(tip.Weight * 100)
			collection.selectors[i] = collection.maxWeight
		}
	}

	rnd := rand.Intn(collection.maxWeight) + 1
	index := searchInts(collection.selectors, rnd)
	tip := collection.data[index]

	color := strutil.Q(tip.ColorTag, ColorTag)
	color = strutil.B(fmtc.IsTag(color), color, "{#75}")

	panel.Panel("❏ PROTIP", color, tip.Title, tip.Message, Options...)

	return true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// searchInts improved sort.SearchInts version
//
// Original: https://github.com/mroth/weightedrand
func searchInts(a []int, x int) int {
	var i int

	j := len(a)

	for i < j {
		h := int(uint(i+j) >> 1)

		if a[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}

	return i
}
