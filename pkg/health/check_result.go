package health

type CheckResult struct {
	CheckName string   `json:"name"`
	Status    Status   `json:"status"`
	Error     string   `json:"error,omitempty"`
	Extra     []string `json:"extra,omitempty"`
}

func Ok(checkName string, extra []string) CheckResult {
	return CheckResult{
		CheckName: checkName,
		Status:    Status_Ok,
		Extra:     extra,
	}
}

func Unhealthy(checkName, errorDetails string, extra []string) CheckResult {
	return CheckResult{
		CheckName: checkName,
		Status:    Status_Unhealthy,
		Error:     errorDetails,
		Extra:     extra,
	}
}
