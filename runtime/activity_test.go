package throttle

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/TIBCOSoftware/flogo-lib/flow/test"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("tibco-throttle")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestThrottleEnabled(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &ThrottleActivity{metadata: md, counters: make(map[string]int), timers: make(map[string]int64)}

	tc := test.NewTestActivityContext(md)

	//setup attrs
	tc.SetInput(ivEndPoint, "http://localhost:9980/now")
	tc.SetInput(ivLimitPerMinute, 1)

	act.Eval(tc)
	act.Eval(tc)

	value := tc.GetOutput(ovThrottled).(bool)

	if !value {
		t.Fail()
	}
}

func TestThrottleDisbled(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &ThrottleActivity{metadata: md, counters: make(map[string]int), timers: make(map[string]int64)}

	tc := test.NewTestActivityContext(md)

	//setup attrs
	tc.SetInput(ivEndPoint, "http://localhost:9980/now")
	tc.SetInput(ivLimitPerMinute, 1)
	tc.SetInput(ivDisable, true)

	act.Eval(tc)
	act.Eval(tc)
	act.Eval(tc)

	value := tc.GetOutput(ovThrottled).(bool)

	if value {
		t.Fail()
	}
}
