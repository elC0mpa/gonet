package network

type AppNetworkInfo struct {
	AppName      string
	NetworkStats NetworkInfo
}

type NetworkInfo struct {
	ReceivedBytes float64
	SentBytes     float64
}
