package network

type AppNetworkInfo struct {
	Info      AppInfo
	NetworkStats NetworkInfo
}

type AppInfo struct {
	ProcessID int
	AppName   string
}

type NetworkInfo struct {
	ReceivedBytes float64
	SentBytes     float64
}
