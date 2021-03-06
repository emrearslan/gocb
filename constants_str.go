package gocb

func serviceTypeToString(service ServiceType) string {
	switch service {
	case ServiceTypeManagement:
		return "mgmt"
	case ServiceTypeKeyValue:
		return "kv"
	case ServiceTypeViews:
		return "views"
	case ServiceTypeQuery:
		return "query"
	case ServiceTypeSearch:
		return "search"
	case ServiceTypeAnalytics:
		return "analytics"
	}
	return ""
}

func serviceTypeFromString(service string) ServiceType {
	switch service {
	case "mgmt":
		return ServiceTypeManagement
	case "kv":
		return ServiceTypeKeyValue
	case "views":
		return ServiceTypeViews
	case "query":
		return ServiceTypeQuery
	case "search":
		return ServiceTypeSearch
	case "analytics":
		return ServiceTypeAnalytics
	}
	return ServiceType(0)
}

func clusterStateToString(state ClusterState) string {
	switch state {
	case ClusterStateOnline:
		return "online"
	case ClusterStateDegraded:
		return "degraded"
	case ClusterStateOffline:
		return "offline"
	}
	return ""
}

func endpointStateToString(state EndpointState) string {
	switch state {
	case EndpointStateDisconnected:
		return "disconnected"
	case EndpointStateConnecting:
		return "connecting"
	case EndpointStateConnected:
		return "connected"
	case EndpointStateDisconnecting:
		return "disconnecting"
	}
	return ""
}

func pingStateToString(state PingState) string {
	switch state {
	case PingStateOk:
		return "ok"
	case PingStateTimeout:
		return "timeout"
	case PingStateError:
		return "error"
	}
	return ""
}
