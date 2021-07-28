package entities

type Multichannel struct {
	appID     string
	secretKey string
}

func NewMultichannel(appID string, secretKey string) *Multichannel {
	return &Multichannel{appID, secretKey}
}

func (m *Multichannel) GetAppID() string {
	return m.appID
}

func (m *Multichannel) GetSecretKey() string {
	return m.secretKey
}
