package maps

import (
	"html/template"

	"go.formulabun.club/srb2kart/addons"
	srb2kartstrings "go.formulabun.club/srb2kart/strings"
)

var templateFuncs = template.FuncMap{
	"removeColorCodes": srb2kartstrings.RemoveColorCodes,
	"isRace": func(file string) bool {
		return addons.GetAddonType(file)&addons.RaceFlag != 0
	},
}
