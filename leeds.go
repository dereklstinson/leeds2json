package leeds2json

import (
	"encoding/json"
)

//Leeds is the stuct that contains the annos for each image.  Should be individial people in the file
type Leeds struct {
	FileName string  `json:"file_name,omitempty"`
	Joints   []Joint `json:"joints,omitempty"`
}

//Joint is a joint
type Joint struct {
	Name  string  `json:"name,omitempty"`
	Label float32 `json:"label,omitempty"`
	Y     float32 `json:"y"`
	X     float32 `json:"x"`
	Vis   float32 `json:"vis"`
}

//GetLeedsExtended returns array of Leeds for Leeds Extended 10,000 images
func GetLeedsExtended() ([]Leeds, error) {

	LSlice := make([]Leeds, 0)
	data := []byte(thisishuge2)
	err := json.Unmarshal(data, &LSlice)
	if err != nil {
		return nil, err
	}
	return LSlice, nil
}
