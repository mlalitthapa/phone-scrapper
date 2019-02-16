package device

type Spec struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Specs []*Spec

type Device struct {
	Name  string           `json:"name"`
	Specs map[string]Specs `json:"specs"`
}
