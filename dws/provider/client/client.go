package client

type DWSClient struct {
	Config          DWSProviderConfiguration
	transactionNote string
}

type DWSProviderConfiguration struct {
	AccountName  string
	AccountKey   string
	AccessRegion string
	ApiKey       string
	SessionToken string
}

func (dc *DWSProviderConfiguration) FromSlice(values []string) {
	if len(values) < 5 {
		return
	}
	dc.AccountName = values[0]
	dc.AccountKey = values[1]
	dc.AccessRegion = values[2]
	dc.ApiKey = values[3]
	dc.SessionToken = values[4]
}

func (c *DWSClient) SetGlobalTransactionNote(note string) {
	c.transactionNote = note
}

func NewClient(configuration DWSProviderConfiguration) *DWSClient {
	return &DWSClient{Config: configuration}
}
