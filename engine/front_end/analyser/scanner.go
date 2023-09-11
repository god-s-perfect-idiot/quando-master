package analyser

import (
	"bufio"
	"strings"
)

type Scanner struct {
	scanner    *bufio.Scanner
	line       string
	sourceCode string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		scanner:    bufio.NewScanner(strings.NewReader(source)),
		sourceCode: source,
	}
}

func (s *Scanner) setLine(line string) {
	s.line = line
}

func (s *Scanner) getLineLength() int {
	return len(s.line)
}

func (s *Scanner) getLine() string {
	return s.line
}

func (s *Scanner) readLine() bool {
	line := s.scanner.Text()
	line = strings.TrimSpace(line)
	s.setLine(line)
	return line != ""
}

func (s *Scanner) split() []string {
	var tokens []string
	currentToken := ""
	inQuotes := false

	input := s.getLine()
	for i := 0; i < len(input); i++ {
		char := input[i]

		if char == '"' {
			inQuotes = !inQuotes
			currentToken += string(char)
		} else if inQuotes {
			currentToken += string(char)
		} else {
			if strings.ContainsRune(" =", rune(char)) {
				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
			} else if i < len(input)-1 && char == ',' && input[i+1] == ' ' {
				if currentToken != "" {
					tokens = append(tokens, currentToken)
					currentToken = ""
				}
				i++ // Skip the space after the comma
			} else {
				currentToken += string(char)
			}
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

// base class overrides

func (s *Scanner) Scan() bool {
	return s.scanner.Scan()
}

func (s *Scanner) Text() string {
	return s.scanner.Text()
}

func (s *Scanner) Bytes() []byte {
	return s.scanner.Bytes()
}

func (s *Scanner) Err() error {
	return s.scanner.Err()
}
