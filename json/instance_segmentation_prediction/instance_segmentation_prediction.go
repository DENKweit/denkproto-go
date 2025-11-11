// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    instanceSegmentationPrediction, err := UnmarshalInstanceSegmentationPrediction(bytes)
//    bytes, err = instanceSegmentationPrediction.Marshal()

package instance_segmentation_prediction

import "encoding/json"

func UnmarshalInstanceSegmentationPrediction(data []byte) (InstanceSegmentationPrediction, error) {
	var r InstanceSegmentationPrediction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstanceSegmentationPrediction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InstanceSegmentationPrediction struct {
	Height                                         int64    `json:"height"`
	// Thresholded segmentation maps for each class         
	Objects                                        []Object `json:"objects"`
	Width                                          int64    `json:"width"`
}

type Object struct {
	Data        string  `json:"data"`
	Height      int64   `json:"height"`
	LabelID     string  `json:"label_id"`
	Probability float64 `json:"probability"`
	Width       int64   `json:"width"`
	X           int64   `json:"x"`
	Y           int64   `json:"y"`
}
