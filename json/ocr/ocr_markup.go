// Code generated from JSON Schema using quicktype. DO NOT EDIT.
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
	Annotations         []Annotation `json:"annotations"`
	AverageObjectWidths []float64    `json:"average_object_widths"`
	Height              int64        `json:"height"`
	Width               int64        `json:"width"`
}

type Annotation struct {
	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`
	ID          string       `json:"id"`
	LabelID     string       `json:"label_id"`
	Polygon     *Polygon     `json:"polygon,omitempty"`
	Text        string       `json:"text"`
}

// A bounding box with optional rotation information
type BoundingBox struct {
	// Optional rotation angle
	Angle        *float64 `json:"angle,omitempty"`
	BottomRightX float64  `json:"bottom_right_x"`
	BottomRightY float64  `json:"bottom_right_y"`
	// Optional full orientation flag
	FullOrientation *bool   `json:"full_orientation,omitempty"`
	TopLeftX        float64 `json:"top_left_x"`
	TopLeftY        float64 `json:"top_left_y"`
}

// A polygon defined by one or more rings, allowing for holes and nested structures.
type Polygon struct {
	// Array of polygon rings. The hierarchy field within each ring determines nesting and
	// fill/hole status.
	Rings []GeometrySchema `json:"rings"`
}

// A single closed loop (ring) of a polygon, defining either an outer boundary or a hole.
type GeometrySchema struct {
	// Nesting level: 0=outer, 1=hole in level 0, 2=poly in level 1 hole, etc. Even levels are
	// filled areas, odd levels are holes.
	Hierarchy int64 `json:"hierarchy"`
	// Vertices of the ring.
	Points []Point `json:"points"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
