package device

type Spec struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Specs []*Spec

type Image struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}

type Device struct {
	Name   string           `json:"name"`
	Specs  map[string]Specs `json:"specs"`
	Images []*Image         `json:"images"`
}
