// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    instanceSegmentationMarkup, err := UnmarshalInstanceSegmentationMarkup(bytes)
//    bytes, err = instanceSegmentationMarkup.Marshal()

package instance_segmentation

import "encoding/json"

func UnmarshalInstanceSegmentationMarkup(data []byte) (InstanceSegmentationMarkup, error) {
	var r InstanceSegmentationMarkup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstanceSegmentationMarkup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InstanceSegmentationMarkup struct {
	// Instance segmentation annotations             
	Annotations                         []Annotation `json:"annotations"`
	Height                              int64        `json:"height"`
	Width                               int64        `json:"width"`
}

type Annotation struct {
	BottomRightX float64 `json:"bottom_right_x"`
	BottomRightY float64 `json:"bottom_right_y"`
	LabelID      string  `json:"label_id"`
	Mask         Mask    `json:"mask"`
	TopLeftX     float64 `json:"top_left_x"`
	TopLeftY     float64 `json:"top_left_y"`
}

// A binary mask with REE (Run-End Encoding) compressed data
type Mask struct {
	Data   string `json:"data"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}
