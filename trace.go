package httpaccesslog


import (
	"time"
)


type Trace struct {
	BeginTime time.Time
	EndTime   time.Time
}
