package health

import (
	"fmt"
	"time"
)

func withTimeout(check Checker, out chan CheckResult) {
	defer close(out)

	r := make(chan CheckResult)
	go check.Run(r)

	select {
	case <-time.After(time.Duration(check.TimeoutSeconds()) * time.Second):
		out <- Unhealthy(check.Name(), fmt.Sprintf("did not respond after %d seconds", check.TimeoutSeconds()), nil)
	case checkResult := <-r:
		out <- checkResult
	}

}

func RunChecks() Health {
	result := CreateResponse(config.ServiceName)
	checkResults := make([]chan CheckResult, len(config.Checks))

	for i, check := range config.Checks {
		checkResults[i] = make(chan CheckResult)
		if check.TimeoutSeconds() > 0 {
			go withTimeout(check, checkResults[i])
		} else {
			go check.Run(checkResults[i])
		}
	}

	for i := 0; i < len(config.Checks); i++ {
		r := <-checkResults[i]
		result.AddResult(r)
		checkResults[i] = nil
	}

	return result
}
