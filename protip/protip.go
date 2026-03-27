// Package protip provides methods for displaying usage tips
package protip

import (
	"math/rand"
	"os"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/fmtutil/panel"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Tip contains the content and display configuration for a single usage tip
type Tip struct {
	Title   string // Tip title (required)
	Message string // Tip message (required)

	ColorTag string  // Custom tip color (optional)
	Weight   float64 // Custom tip weight (optional)
}

// Tips holds the registered tip collection and its precomputed weighted selector state
type Tips struct {
	data      []*Tip
	selectors []int
	maxWeight int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Probability is the chance (0.0–1.0) that [Show] will display a tip when force
// is false
var Probability = 0.25

// ColorTag is the default fmtc color tag applied to the panel border when a tip
// has none set
var ColorTag = "{#75}"

// Options holds the default panel rendering options applied to every displayed tip
var Options = panel.Options{panel.TOP_LINE, panel.BOTTOM_LINE}

// ////////////////////////////////////////////////////////////////////////////////// //

var disabled = os.Getenv("PROTIP") == "0"

var collection *Tips

// ////////////////////////////////////////////////////////////////////////////////// //

// Add appends one or more tips to the global collection, silently skipping nil entries,
// tips with empty Title or Message, or tips with an invalid ColorTag
func Add(tips ...*Tip) {
	if disabled {
		return
	}

	if collection == nil {
		collection = &Tips{}
	}

	for _, tip := range tips {
		if !isValidTip(tip) {
			continue
		}

		if tip.Weight == 0 {
			tip.Weight = 0.5
		}

		collection.data = append(collection.data, tip)
	}
}

// Show displays a randomly selected tip from the global collection based on weighted
// probability; if force is true the probability check is bypassed. Returns true if
// a tip was shown
func Show(force bool) bool {
	if collection == nil || len(collection.data) == 0 || disabled {
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

// isValidTip validats tip
func isValidTip(tip *Tip) bool {
	if tip == nil || tip.Title == "" || tip.Message == "" {
		return false
	}

	if tip.ColorTag != "" && !fmtc.IsTag(tip.ColorTag) {
		return false
	}

	return true
}

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
