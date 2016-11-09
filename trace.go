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
	ID        [32]byte
}


func (trace *Trace) initialize() {
	trace.BeginTime = time.Now()

	traceID := trace.ID[:]

	unixTime := uint64(trace.BeginTime.Unix())

	// 1111_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[0]  = defaultTraceIDCharacters[ (0xf000000000000000 & unixTime) >> 58]

	// 0000_1111_1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[1]  = defaultTraceIDCharacters[ (0x0fc0000000000000 & unixTime) >> 52]

	// 0000_0000_0011_1111_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[2]  = defaultTraceIDCharacters[ (0x003f000000000000 & unixTime) >> 46]

	// 0000_0000_0000_0000_1111_1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[3]  = defaultTraceIDCharacters[ (0x0000fc0000000000 & unixTime) >> 40]

	// 0000_0000_0000_0000_0000_0011_1111_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[4]  = defaultTraceIDCharacters[ (0x000003f000000000 & unixTime) >> 36]

	// 0000_0000_0000_0000_0000_0000_0000_1111_1100_0000_0000_0000_0000_0000_0000_0000
	traceID[5]  = defaultTraceIDCharacters[ (0x0000000fc0000000 & unixTime) >> 30]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0011_1111_0000_0000_0000_0000_0000_0000
	traceID[6]  = defaultTraceIDCharacters[ (0x000000003f000000 & unixTime) >> 24]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_1100_0000_0000_0000_0000
	traceID[7]  = defaultTraceIDCharacters[ (0x0000000000fc0000 & unixTime) >> 18]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0011_1111_0000_0000_0000
	traceID[8]  = defaultTraceIDCharacters[ (0x000000000003f000 & unixTime) >> 12]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_1100_0000
	traceID[9]  = defaultTraceIDCharacters[ (0x0000000000000fc0 & unixTime) >> 6]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0011_1111
	traceID[10] = defaultTraceIDCharacters[ (0x000000000000003f & unixTime)]

	const offset = 10 // <---- This needs to match the last index manually assigned to in the code before it.

	length := len(defaultTraceIDCharacters)
	for i,_ := range traceID[offset:] {
		traceID[offset+i] = defaultTraceIDCharacters[randomness.Intn(length)]
	}
}
