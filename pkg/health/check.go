package health

// To create a check without an asynchronous timeout, return 0 for TimeoutSeconds()
// To create a simple check, you can call CreateHealthCheck or CreateHealthCheckWithTimeout
// For more control, create your own implementation of the Checker interface
type Checker interface {
	Name() string
	TimeoutSeconds() int
	Run(c chan CheckResult)
}

type check struct {
	name           string
	timeoutSeconds int
	run            HealthCheckFunc
}

type HealthCheckFunc func(name string, out chan CheckResult)

func CreateHealthCheck(name string, run HealthCheckFunc) Checker {
	return CreateHealthCheckWithTimeout(name, 0, run)
}

func CreateHealthCheckWithTimeout(name string, timeoutSeconds int, run HealthCheckFunc) Checker {
	return &check{
		name:           name,
		timeoutSeconds: timeoutSeconds,
		run:            run,
	}
}

func (c *check) Name() string {
	return c.name
}

func (c *check) TimeoutSeconds() int {
	return c.timeoutSeconds
}

func (c *check) Run(out chan CheckResult) {
	c.run(c.Name(), out)
}
