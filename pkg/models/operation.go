package models

type Operation struct {
	ID   int32
	Name string
}

var (
	OperationUnknown = Operation{ID: 0, Name: "unknown"}
	OperationEquals  = Operation{ID: 1, Name: "EQUAL"}
	OperationGreater = Operation{ID: 2, Name: "GREATER"}
	OperationLess    = Operation{ID: 3, Name: "LESS"}
)
