// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ocrMarkup, err := UnmarshalOcrMarkup(bytes)
//    bytes, err = ocrMarkup.Marshal()

package ocr

import "encoding/json"

func UnmarshalOcrMarkup(data []byte) (OcrMarkup, error) {
	var r OcrMarkup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OcrMarkup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OcrMarkup struct {
	Annotations []Annotation `json:"annotations"`
	Height      int64        `json:"height"`
	Width       int64        `json:"width"`
}

type Annotation struct {
	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`
	ID          string       `json:"id"`
	LabelID     string       `json:"label_id"`
	Polygon     *Polygon     `json:"polygon,omitempty"`
	Text        string       `json:"text"`
}

type BoundingBox struct {
	BottomRightX float64 `json:"bottom_right_x"`
	BottomRightY float64 `json:"bottom_right_y"`
	TopLeftX     float64 `json:"top_left_x"`
	TopLeftY     float64 `json:"top_left_y"`
}

type Polygon struct {
	Points []Point `json:"points"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
