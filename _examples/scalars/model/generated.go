// Code generated by github.com/operandinc/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/operandinc/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location"`
}
