package proxy

// usProxyData contains data generated from https://www.us-proxy.com/
type usProxyDataRow struct {
	IP          string
	Port        string
	Code        string
	Country     string
	Anonymity   string
	Google      string
	Https       string
	LastChecked string
}

type usProxyData []usProxyDataRow

// Marshall maps [][]string csv data into usProxyData struct
func (d *usProxyData) Marshal(s [][]string) error {
	workingD := []usProxyDataRow{}
	for _, row := range s {
		mappedRow := usProxyDataRow{}
		mappedRow.IP = row[0]
		mappedRow.Port = row[1]
		mappedRow.Code = row[2]
		mappedRow.Country = row[3]
		mappedRow.Anonymity = row[4]
		mappedRow.Google = row[5]
		mappedRow.Https = row[6]
		mappedRow.LastChecked = row[7]
		workingD = append(workingD, mappedRow)
	}
	d = (*usProxyData)(&workingD)
	return nil
}
