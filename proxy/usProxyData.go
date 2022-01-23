package proxy

// usProxyData contains data generated from https://www.us-proxy.com/
type usProxy struct {
	IP          string
	Port        string
	Code        string
	Country     string
	Anonymity   string
	Google      string
	Https       string
	LastChecked string
}

type usProxyData []usProxy

// unmarshal maps [][]string csv data into usProxyData struct
func unmarshal(s [][]string, i *[]usProxy) error {
	proxies := []usProxy{}
	for _, row := range s {
		mappedRow := usProxy{}
		mappedRow.IP = row[0]
		mappedRow.Port = row[1]
		mappedRow.Code = row[2]
		mappedRow.Country = row[3]
		mappedRow.Anonymity = row[4]
		mappedRow.Google = row[5]
		mappedRow.Https = row[6]
		mappedRow.LastChecked = row[7]
		proxies = append(proxies, mappedRow)
	}
	*i = proxies
	return nil
}
