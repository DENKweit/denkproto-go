// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    classificationPrediction, err := UnmarshalClassificationPrediction(bytes)
//    bytes, err = classificationPrediction.Marshal()

package classification_prediction

import "encoding/json"

func UnmarshalClassificationPrediction(data []byte) (ClassificationPrediction, error) {
	var r ClassificationPrediction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClassificationPrediction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ClassificationPrediction struct {
	Predictions []Prediction `json:"predictions"`
}

type Prediction struct {
	LabelID     string  `json:"label_id"`
	Probability float64 `json:"probability"`
}
