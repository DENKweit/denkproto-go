// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    objectDetectionPrediction, err := UnmarshalObjectDetectionPrediction(bytes)
//    bytes, err = objectDetectionPrediction.Marshal()

package object_detection_prediction

import "encoding/json"

func UnmarshalObjectDetectionPrediction(data []byte) (ObjectDetectionPrediction, error) {
	var r ObjectDetectionPrediction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ObjectDetectionPrediction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ObjectDetectionPrediction struct {
	Height      int64        `json:"height"`
	Predictions []Prediction `json:"predictions"`
	Width       int64        `json:"width"`
}

type Prediction struct {
	Angle           *float64 `json:"angle,omitempty"`
	BottomRightX    float64  `json:"bottom_right_x"`
	BottomRightY    float64  `json:"bottom_right_y"`
	FullOrientation *bool    `json:"full_orientation,omitempty"`
	LabelID         string   `json:"label_id"`
	Probability     float64  `json:"probability"`
	TopLeftX        float64  `json:"top_left_x"`
	TopLeftY        float64  `json:"top_left_y"`
}
