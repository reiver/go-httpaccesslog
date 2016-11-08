package httpaccesslog


import (
	"math/rand"
	"time"
)


var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	defaultTraceIDCharacters = []byte{
			'.',

			'0',
			'1',
			'2',
			'3',
			'4',
			'5',
			'6',
			'7',
			'8',
			'9',

			'A',
			'B',
			'C',
			'D',
			'E',
			'F',
			'G',
			'H',
			'I',
			'J',
			'K',
			'L',
			'M',
			'N',
			'O',
			'P',
			'Q',
			'R',
			'S',
			'T',
			'U',
			'V',
			'W',
			'X',
			'Y',
			'Z',

			'_',

			'a',
			'b',
			'c',
			'd',
			'e',
			'f',
			'g',
			'h',
			'i',
			'j',
			'k',
			'l',
			'm',
			'n',
			'o',
			'p',
			'q',
			'r',
			's',
			't',
			'u',
			'v',
			'w',
			'x',
			'y',
			'z',
	}
)


type Trace struct {
	BeginTime time.Time
	EndTime   time.Time
	ID        [126]byte
}


func generateTraceID(p []byte) {
	length := len(defaultTraceIDCharacters)

	for i,_ := range p {
		p[i] = defaultTraceIDCharacters[randomness.Intn(length)]
	}
}
