package object

type ClientInfo struct {
	OsInfo  osInfo  `json:"os"`
	NetInfo netInfo `json:"net"`
	DbInfo  dbInfo  `json:"db"`
	ErpInfo erpInfo `json:"erp"`
}

type osInfo struct {
	Type string `json:"type"`
	//Name     string `json:"name"`
	//Ver      string `json:"ver"`
	HostName string `json:"hostname"`
}

type netInfo struct {
	InternetIp string `json:"internet"`
	IntranetIp string `json:"intranet"`
}

type dbInfo struct {
	Ver string `json:"ver"`
}

type erpInfo struct {
	Ver string `json:"ver"`
}
