package types

import "bytes"

type CardFact struct {
	Icon        *bytes.Buffer
	Title       string
	Description string
}
