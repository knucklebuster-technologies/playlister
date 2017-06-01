package models

import "io"

// Modeler interface for models to marshal and decode JSON
type Modeler interface {
	Encode(io.Writer) error
	Decode(io.Reader) error
}
