// Created by EldersJavas(EldersJavas&gmail.com)

package types

import "encoding/json"

func UnmarshalPotScene(data []byte) (PotScene, error) {
	var r PotScene
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PotScene) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PotScene struct {
	Config *Config       `json:"config,omitempty"`
	Object []interface{} `json:"object,omitempty"`
}
