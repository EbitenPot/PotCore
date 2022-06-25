// Created by EldersJavas(EldersJavas&gmail.com)

package types

type Config struct {
	Ver    *int64       `json:"ver,omitempty"`
	Ebiver *string      `json:"ebiver,omitempty"`
	Uid    *string      `json:"uid,omitempty"`
	Name   *string      `json:"name,omitempty"`
	Extra  *interface{} `json:"extra,omitempty"`
}
