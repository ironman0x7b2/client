package types

type Node struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Deposit Coin   `json:"deposit"`

	IP   string `json:"ip"`
	Port string `json:"port"`

	Type          string    `json:"type"`
	Version       string    `json:"version"`
	Moniker       string    `json:"moniker"`
	PricesPerGB   []Coin    `json:"prices_per_gb"`
	InternetSpeed Bandwidth `json:"internet_speed"`
	Encryption    string    `json:"encryption"`

	Status string `json:"status" bson:"status"`
}

type Bandwidth struct {
	Upload   int64 `json:"upload"`
	Download int64 `json:"download"`
}

type Nodes struct {
	Nodes    []Node
	Resolver string
}
