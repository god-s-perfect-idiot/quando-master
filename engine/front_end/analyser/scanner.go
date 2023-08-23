package analyser

import (
	"bufio"
	"strings"
)

type Scanner struct {
	scanner *bufio.Scanner
	line    string
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		scanner: bufio.NewScanner(strings.NewReader(source)),
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
	delimiters := []string{" ", ", ", "="}
	parts := make([]string, 0)
	startIdx := 0
	for i := 0; i < len(s.line); i++ {
		for _, delimiter := range delimiters {
			delimiterLen := len(delimiter)
			if i+delimiterLen <= len(s.line) && s.line[i:i+delimiterLen] == delimiter {
				if i > startIdx {
					parts = append(parts, s.line[startIdx:i])
				}
				i += delimiterLen - 1
				startIdx = i + 1
				break
			}
		}
	}
	if startIdx < len(s.line) {
		parts = append(parts, s.line[startIdx:])
	}
	return parts
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
