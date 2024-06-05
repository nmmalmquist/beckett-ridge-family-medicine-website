package types

import "bytes"

type Modal struct {
	Type             string
	ModalTitle       string
	ModalSubtitle    string
	ModalIcon        *bytes.Buffer
	ModalIconBgColor string
	ClearFormOnClose bool
}
