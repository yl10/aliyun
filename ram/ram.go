package ram

type ResourceAcs struct {
	notMatchAll bool
	serviceName string
	region      string
	accountID   string
	relativeID  string
}

func (r *ResourceAcs) Set(serviceName, region, accountID, relativeID string) *ResourceAcs {

	r.serviceName = r.withDefault(serviceName)
	r.region = r.withDefault(region)
	r.accountID = r.withDefault(accountID)
	r.relativeID = r.withDefault(relativeID)

	return r
}

func (r *ResourceAcs) SetServiceName(serviceName string) *ResourceAcs {
	r.serviceName = r.withDefault(serviceName)

	return r
}

func (r *ResourceAcs) SetRegion(region string) *ResourceAcs {
	r.region = r.withDefault(region)

	return r
}

func (r *ResourceAcs) SetAccountID(accountID string) *ResourceAcs {
	r.accountID = r.withDefault(accountID)

	return r
}

func (r *ResourceAcs) SetRelativeID(relativeID string) *ResourceAcs {
	r.relativeID = r.withDefault(relativeID)

	return r
}

func (r *ResourceAcs) SetMatchAll(matchAll bool) *ResourceAcs {
	r.notMatchAll = !matchAll

	return r
}

func (r *ResourceAcs) String() string {
	serviceName := r.withDefault(r.serviceName)
	region := r.withDefault(r.region)
	accountID := r.withDefault(r.accountID)
	relativeID := r.withDefault(r.relativeID)

	return "acs:" + serviceName + ":" + region + ":" + accountID + ":" + relativeID
}

func (r *ResourceAcs) withDefault(val string) string {
	if val != "" {
		return val
	}

	if r.notMatchAll {
		return ""
	}

	return "*"
}
