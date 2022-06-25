// Created by EldersJavas(EldersJavas&gmail.com)

package types

import "encoding/json"

func UnmarshalPotObject(data []byte) (PotObject, error) {
	var r PotObject
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PotObject) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PotObject struct {
	Objconf *Objconf     `json:"objconf,omitempty"`
	Content *interface{} `json:"content,omitempty"`
}

type Objconf struct {
	Ver      *int64       `json:"ver,omitempty"`
	Ebiver   *string      `json:"ebiver,omitempty"`
	Uid      *string      `json:"uid,omitempty"`
	Name     *string      `json:"name,omitempty"`
	Extra    *interface{} `json:"extra,omitempty"`
	Typename *string      `json:"typename,omitempty"`
}
