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
	if len(words) == 0 {
		return structures.NewLineBreak(l.cursor.getL())
	}
	if len(words) == 1 {
		return structures.NewCallbackterminator(l.cursor.getL())
	} else {
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
				type_ = Tokens[NUMBER]
			} else if IsStringParameter(value) {
				type_ = Tokens[STRING]
			} else if IsVal(value) {
				type_ = Tokens[VAL]
			} else if IsCallbackParameter(value) {
				type_ = Tokens[CALLBACK]
				hasCallback = true
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
	}
}

func (l *Analyser) Scan() *structures.Essence {
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
	scriptEssence := structures.NewExecutionContext(invocationTable, callGraph)
	return scriptEssence
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
