package types

type InputString = string
type InputChar = byte

func ToInputString(ch InputChar) InputString {
	return InputString(ch)
}
