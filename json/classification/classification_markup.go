// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    classificationMarkup, err := UnmarshalClassificationMarkup(bytes)
//    bytes, err = classificationMarkup.Marshal()

package classification

import "encoding/json"

func UnmarshalClassificationMarkup(data []byte) (ClassificationMarkup, error) {
	var r ClassificationMarkup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClassificationMarkup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ClassificationMarkup struct {
	Annotations []Annotation `json:"annotations"`
	Height      int64        `json:"height"`
	Width       int64        `json:"width"`
}

type Annotation struct {
	ID      string  `json:"id"`
	LabelID string  `json:"label_id"`
	Value   float64 `json:"value"`
}
