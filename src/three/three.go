package three

import "encoding/json"

type A struct {
	b         *B
	something json.RawMessage
}

type B struct {
	a A
}
