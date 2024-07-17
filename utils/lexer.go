package utils

type TokenType int

const (
	WORD TokenType = iota
	NUMBER
	PUNCTUATION
	SPACE
	NEWLINE
	EOF
)

type Token struct {
	mValue     string
	mTokenType TokenType
}

func (t *Token) GetValue() string {
	return t.mValue
}
func (t *Token) GetTokenType() TokenType {
	return t.mTokenType
}
func (t *Token) SetValue(value string) {
	t.mValue = value
}
func (t *Token) SetTokenType(tokenType TokenType) {
	t.mTokenType = tokenType
}
func (t *Token) AppendValue(value string) {
	t.mValue += value
}

type Lexer struct {
	mTokenSlice []Token //
}

func (i *Lexer) Analyse(analyzedString string) {
	token := Token{}

	for _, char := range analyzedString {
		index := len(i.mTokenSlice)
		if char == ' ' {
			token.SetValue(" ")
			token.SetTokenType(SPACE)
		} else if char == '\n' {
			token.SetValue("\n")
			token.SetTokenType(NEWLINE)
		} else if char == ',' || char == '!' || char == '?' {
			token.SetValue(string(char))
			token.SetTokenType(PUNCTUATION)
		} else if char >= '0' && char <= '9' || char == '.' {
			token.AppendValue(string(char))
			token.SetTokenType(NUMBER)
			//	(i.mTokenSlice) = append(i.mTokenSlice, token)
			continue
		} else {
			token.AppendValue(string(char))
			token.SetTokenType(WORD)
			if len(i.mTokenSlice) == 0 {
				(i.mTokenSlice) = append(i.mTokenSlice, token)
				continue
			}
			if i.mTokenSlice[index-1].mTokenType != WORD {
				(i.mTokenSlice) = append(i.mTokenSlice, token)
				continue
			} else {
				i.mTokenSlice[index-1].mValue = token.GetValue()
				continue
			}

		}
		(i.mTokenSlice) = append(i.mTokenSlice, token)
		token.SetValue("")
	}
}
func (i *Lexer) GetTokenSlice() *[]Token {
	return &i.mTokenSlice
}
