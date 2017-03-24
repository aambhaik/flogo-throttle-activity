package throttle

import (
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"sync"
	"time"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-throttle")

const (
	ivEndPoint       = "endPoint"
	ivLimitPerMinute = "limitPerMinute"
	ivDisable        = "disable"

	ovThrottled = "throttled"
)

// ThrottleActivity is a Throttle Activity implementation
type ThrottleActivity struct {
	sync.Mutex
	metadata *activity.Metadata
	counters map[string]int
	timers   map[string]int64
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&ThrottleActivity{metadata: md, counters: make(map[string]int), timers: make(map[string]int64)})
}

// Metadata implements activity.Activity.Metadata
func (a *ThrottleActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ThrottleActivity) Eval(context activity.Context) (bool, error) {

	isDisabled := false
	if context.GetInput(ivDisable) != nil {
		isDisabled = context.GetInput(ivDisable).(bool)
	}
	endPoint := context.GetInput(ivEndPoint).(string)
	limitPerMinute := context.GetInput(ivLimitPerMinute).(int)

	if isDisabled {
		// no throttling the endpoint..
		log.Debugf("No Throttling for the endpoint '%s'", endPoint)

		context.SetOutput(ovThrottled, false)

	} else {
		throttle := a.evaluateThrottleCondition(endPoint, limitPerMinute)
		log.Debugf("Throttle for endpoint [%s] : %v", endPoint, throttle)

		context.SetOutput(ovThrottled, throttle)
	}

	return true, nil
}

func (a *ThrottleActivity) evaluateThrottleCondition(endPoint string, limitPerMinute int) bool {
	a.Lock()
	defer a.Unlock()

	throttle := false
	count := 1

	if counter, exists := a.counters[endPoint]; exists {
		firstInvocationTime := a.timers[endPoint]
		currentTime := time.Now().Unix()

		if counter < limitPerMinute && currentTime-firstInvocationTime < 60 {
			//less than a minute has elapsed since first invocation and the counter is still not reached the limit
			count = counter + 1
		} else if counter == limitPerMinute && currentTime-firstInvocationTime < 60 {
			//less than a minute has elapsed since first invocation and the counter has reached the limit. Throttle in effect now!
			throttle = true
			return throttle
		} else if currentTime-firstInvocationTime >= 60 {
			//a minute has elapsed since first invocation. reset the throttle counters
			a.resetThrottles(endPoint)
			return throttle
		}
	}

	a.counters[endPoint] = count
	if count == 1 {
		//first count
		a.timers[endPoint] = time.Now().Unix()
	}

	return throttle
}

func (a *ThrottleActivity) resetThrottles(endPoint string) {
	a.Lock()
	defer a.Unlock()

	delete(a.counters, endPoint)
	delete(a.timers, endPoint)
}
