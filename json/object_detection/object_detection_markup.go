// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    objectDetectionMarkup, err := UnmarshalObjectDetectionMarkup(bytes)
//    bytes, err = objectDetectionMarkup.Marshal()

package object_detection

import "encoding/json"

func UnmarshalObjectDetectionMarkup(data []byte) (ObjectDetectionMarkup, error) {
	var r ObjectDetectionMarkup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ObjectDetectionMarkup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ObjectDetectionMarkup struct {
	Annotations []Annotation `json:"annotations"`
	Height      int64        `json:"height"`
	Width       int64        `json:"width"`
}

type Annotation struct {
	Angle           *float64       `json:"angle,omitempty"`
	AnnotationType  AnnotationType `json:"annotation_type"`
	AverageWidth    float64        `json:"average_width"`
	BottomRightX    float64        `json:"bottom_right_x"`
	BottomRightY    float64        `json:"bottom_right_y"`
	FullOrientation *bool          `json:"full_orientation,omitempty"`
	ID              string         `json:"id"`
	LabelID         string         `json:"label_id"`
	TopLeftX        float64        `json:"top_left_x"`
	TopLeftY        float64        `json:"top_left_y"`
}

type AnnotationType string

const (
	Ignore   AnnotationType = "IGNORE"
	Negative AnnotationType = "NEGATIVE"
	Positive AnnotationType = "POSITIVE"
	Roi      AnnotationType = "ROI"
)
