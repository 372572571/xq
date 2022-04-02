package lang

var Placeholder PlaceholderType

type (
	Any   = interface{}
	Map   = map[string]Any
	Array = []Any

	// seize a seat
	PlaceholderType = struct{}
)
