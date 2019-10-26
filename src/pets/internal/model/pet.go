package model

import "encoding/json"

type Pet struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Pets []Pet

func (p Pets) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
