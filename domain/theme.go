package domain

import (
	"../ddd"
)

type Theme struct {
	ddd.Entity
	ThemeName string
}
