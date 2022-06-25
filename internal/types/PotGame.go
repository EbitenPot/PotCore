// Created by EldersJavas(EldersJavas&gmail.com)

package types

// TODO: add types
import "encoding/json"

func UnmarshalPotGame(data []byte) (PotGame, error) {
	var r PotGame
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PotGame) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

/*
PotGame
PotGame,即Game结构,也就是PotCore里最高的结构,下设Config(GameConfig)和Scene(场景)
*/
type PotGame struct {
	Config *Config       `json:"config,omitempty"`
	Scene  []interface{} `json:"scene,omitempty"`
}
