package shared

import "fmt"

var (
	NotImplemented    = fmt.Errorf("not implemented")
	PathEmpty         = fmt.Errorf("path cannot be empty")
	PathAlreadyExists = fmt.Errorf("path already exists")

	EngineNotFound = fmt.Errorf("engine not found")
)
