package two

import "encoding/json"

type A struct {
	b    *B
	json json.RawMessage
}

type B struct {
	a A
}
