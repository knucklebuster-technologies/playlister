package models

// Modeler interface for models to marshal and unmarshal JSON
type Modeler interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
