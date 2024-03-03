package types

type InputString = []rune
type InputChar = rune

func ToInputString(ch InputChar) InputString {
	return InputString{ch}
}

/*
모든 단일 2~4 바이트 문자에 대해 완전 탐색을 했을 때, 0~127의 값의 바이트를 가지는 문자는 존재하지 않았습니다
즉, 구분자가 모두 ASCII 문자라면, 128 이상의 바이트를 모두 구분자가 아닌 것으로 취급하면 rune을 사용하지 않아도 됩니다.
*/
