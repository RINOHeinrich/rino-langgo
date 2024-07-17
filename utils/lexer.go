package utils

type TokenType int

const (
	WORD TokenType = iota
	NUMBER
	OPERATOR
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
		} else if char == '+' || char == '!' || char == '-' || char == '*' || char == '/' {
			token.SetValue(string(char))
			token.SetTokenType(OPERATOR)
		} else {
			token.AppendValue(string(char))
			token.SetTokenType(WORD)
			if char >= '0' && char <= '9' || char == '.' {
				token.SetTokenType(NUMBER)
			}
			// if the token slice is empty, append the token
			if len(i.mTokenSlice) == 0 {
				(i.mTokenSlice) = append(i.mTokenSlice, token)
				continue
			}
			// if the previous token is a word or a number, append the token
			if i.mTokenSlice[index-1].mTokenType != WORD && i.mTokenSlice[index-1].mTokenType != NUMBER {
				(i.mTokenSlice) = append(i.mTokenSlice, token)
				//if token contains a word then set token type to word
				continue
			} else {
				i.mTokenSlice[index-1] = token
				//i.mTokenSlice[index-1].mTokenType = token.GetTokenType()
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
