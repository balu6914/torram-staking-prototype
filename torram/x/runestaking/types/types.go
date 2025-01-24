package types

type Stake struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
	Asset   string `json:"asset"` // BTC or Runes
}
