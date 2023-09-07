package analyser

import (
	structures "quando/engine/structures"
)

const (
	// Token types
	// Terminal tokens
	EOL = iota
	EOF

	// Redundant tokens
	EMPTY
	WHITESPACE

	// Operations
	ASSIGN

	// Literals
	STRING
	NUMBER
	FLOAT
	BOOLEAN
	CALLBACK
	VAL

	// Delimiters
	COMMA
	OPEN_BRACE
	CLOSE_BRACE
)

type Token struct {
	tokenType int
	value     string
}

var Tokens = []string{
	EOL:         "EOL",
	EOF:         "EOF",
	EMPTY:       "",
	WHITESPACE:  " ",
	ASSIGN:      "=",
	STRING:      "STRING",
	NUMBER:      "NUMBER",
	BOOLEAN:     "BOOLEAN",
	CALLBACK:    "CALLBACK",
	FLOAT:       "FLOAT",
	VAL:         "VAL",
	COMMA:       ",",
	OPEN_BRACE:  "{",
	CLOSE_BRACE: "}",
}

type Analyser struct {
	cursor  Cursor
	scanner Scanner
}

func NewAnalyser(sourceCode string) *Analyser {
	return &Analyser{
		cursor:  *NewCursor(),
		scanner: *NewScanner(sourceCode),
	}
}

func (l *Analyser) ConsumeCharacter() byte {
	line := l.scanner.getLine()
	char := line[l.cursor.getC()]
	l.cursor.advanceColumn()
	return char
}

func (l *Analyser) tokenize() *structures.Definition {
	words := l.scanner.split()
	switch FindLineType(words) {
	case "lineBreak":
		return structures.NewLineBreak(l.cursor.getL())
	case "callbackTerminator":
		return structures.NewCallbackterminator(l.cursor.getL())
	case "conditionalCallback":
		return structures.NewConditionalCallback(l.cursor.getL())
	case "invocation":
		methodIdentifier := words[0]
		var parameters []structures.Parameter
		hasCallback := false
		for i := 1; i < len(words); i += 2 {
			var type_ string
			key := words[i]
			value := words[i+1]
			if IsBooleanParameter(value) {
				type_ = Tokens[BOOLEAN]
			} else if IsIntegerParameter(value) {
				type_ = Tokens[NUMBER]
			} else if IsFloatParameter(value) {
				type_ = Tokens[FLOAT]
			} else if IsStringParameter(value) {
				type_ = Tokens[STRING]
			} else if IsVal(value) {
				type_ = Tokens[VAL]
			} else if IsCallbackParameter(value) {
				hasCallback = true
				continue
			}
			parameter := structures.Parameter{
				Identifier: key,
				Value:      value,
				Type:       type_,
			}
			parameters = append(parameters, parameter)
		}
		callSignature := structures.NewCallSignature(methodIdentifier, parameters, hasCallback)
		return structures.NewInvocation(*callSignature, l.cursor.getL())
	default:
		return nil
	}
}

func (l *Analyser) Scan() *structures.Executable {
	definitions := make([]structures.Definition, 0)
	for l.scanner.Scan() {
		if err := l.scanner.Err(); err != nil {
			panic(err)
		}
		valid := l.scanner.readLine()
		if valid {
			definition := l.tokenize()
			definitions = append(definitions, *definition)
			l.cursor.advanceLine()
		}
	}
	callGraph := structures.ConstructCallGraph(definitions)
	invocationTable := structures.GenerateInvocationTable(definitions)
	hashID := structures.GetHash(l.scanner.sourceCode)
	scriptExecutable := structures.NewExecutionContext(hashID, invocationTable, callGraph)
	return scriptExecutable
}

func Run() {
	const testQuery = `
	quando.title text='CLICKED', append=false
	quando.addButton text='click', up_down='down', callback = {
		quando.title text='CLICKED', append=false, callback = {
			quando.title text='CLICKED', append=false		
		}
		quando.title text='CLICKED', append=false, callback = {
			quando.title text='CLICKED', append=false, callback = {
				quando.title text='CLICKED', append=false		
			}
		}
	}
	quando.title text='CLICKED', append=false
	quando.title text='CLICKED', append=false
	`
	analyser := NewAnalyser(testQuery)
	analyser.Scan()
}
