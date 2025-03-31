// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    segmentationMarkup, err := UnmarshalSegmentationMarkup(bytes)
//    bytes, err = segmentationMarkup.Marshal()

package segmentation

import "encoding/json"

func UnmarshalSegmentationMarkup(data []byte) (SegmentationMarkup, error) {
	var r SegmentationMarkup
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SegmentationMarkup) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SegmentationMarkup struct {
	Annotations         []Annotation      `json:"annotations"`
	AverageObjectWidths []float64         `json:"average_object_widths"`
	Height              int64             `json:"height"`
	SegmentationMaps    []SegmentationMap `json:"segmentation_maps"`
	Width               int64             `json:"width"`
}

type Annotation struct {
	AnnotationType      AnnotationType       `json:"annotation_type"`
	AverageWidth        float64              `json:"average_width"`
	CircleAnnotation    *CircleAnnotation    `json:"circle_annotation,omitempty"`
	ID                  string               `json:"id"`
	LabelID             string               `json:"label_id"`
	MagicwandAnnotation *MagicwandAnnotation `json:"magicwand_annotation,omitempty"`
	PenAnnotation       *PenAnnotation       `json:"pen_annotation,omitempty"`
	PixelAnnotation     *PixelAnnotation     `json:"pixel_annotation,omitempty"`
	PolygonAnnotation   *PolygonAnnotation   `json:"polygon_annotation,omitempty"`
	RectangleAnnotation *RectangleAnnotation `json:"rectangle_annotation,omitempty"`
	SausageAnnotation   *SausageAnnotation   `json:"sausage_annotation,omitempty"`
}

type CircleAnnotation struct {
	CenterX float64 `json:"center_x"`
	CenterY float64 `json:"center_y"`
	Radius  float64 `json:"radius"`
}

type MagicwandAnnotation struct {
	BottomRightX float64                    `json:"bottom_right_x"`
	BottomRightY float64                    `json:"bottom_right_y"`
	CenterX      float64                    `json:"center_x"`
	CenterY      float64                    `json:"center_y"`
	DataURL      string                     `json:"dataURL"`
	Points       []MagicwandAnnotationPoint `json:"points"`
	Threshold    int64                      `json:"threshold"`
	TopLeftX     float64                    `json:"top_left_x"`
	TopLeftY     float64                    `json:"top_left_y"`
}

type MagicwandAnnotationPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PenAnnotation struct {
	BottomRightX float64              `json:"bottom_right_x"`
	BottomRightY float64              `json:"bottom_right_y"`
	DataURL      string               `json:"dataURL"`
	Points       []PenAnnotationPoint `json:"points"`
	Thickness    float64              `json:"thickness"`
	TopLeftX     float64              `json:"top_left_x"`
	TopLeftY     float64              `json:"top_left_y"`
}

type PenAnnotationPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PixelAnnotation struct {
	BlobID       string  `json:"blob_id"`
	BottomRightX float64 `json:"bottom_right_x"`
	BottomRightY float64 `json:"bottom_right_y"`
	TopLeftX     float64 `json:"top_left_x"`
	TopLeftY     float64 `json:"top_left_y"`
}

// A polygon defined by one or more rings, allowing for holes and nested structures.
type PolygonAnnotation struct {
	// Array of polygon rings. The hierarchy field within each ring determines nesting and                           
	// fill/hole status.                                                                                             
	Rings                                                                                 []SegmentationMarkupSchema `json:"rings"`
}

// A single closed loop (ring) of a polygon, defining either an outer boundary or a hole.
type SegmentationMarkupSchema struct {
	// Nesting level: 0=outer, 1=hole in level 0, 2=poly in level 1 hole, etc. Even levels are            
	// filled areas, odd levels are holes.                                                                
	Hierarchy                                                                                 int64       `json:"hierarchy"`
	// Vertices of the ring.                                                                              
	Points                                                                                    []RingPoint `json:"points"`
}

type RingPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type RectangleAnnotation struct {
	BottomRightX float64 `json:"bottom_right_x"`
	BottomRightY float64 `json:"bottom_right_y"`
	TopLeftX     float64 `json:"top_left_x"`
	TopLeftY     float64 `json:"top_left_y"`
}

type SausageAnnotation struct {
	BottomRightX float64                  `json:"bottom_right_x"`
	BottomRightY float64                  `json:"bottom_right_y"`
	DataURL      string                   `json:"dataURL"`
	Points       []SausageAnnotationPoint `json:"points"`
	Radius       float64                  `json:"radius"`
	TopLeftX     float64                  `json:"top_left_x"`
	TopLeftY     float64                  `json:"top_left_y"`
}

type SausageAnnotationPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type SegmentationMap struct {
	BlobID    string `json:"blob_id"`
	LabelID   string `json:"label_id"`
	Thumbnail string `json:"thumbnail"`
}

type AnnotationType string

const (
	Ignore   AnnotationType = "IGNORE"
	Negative AnnotationType = "NEGATIVE"
	Positive AnnotationType = "POSITIVE"
	Roi      AnnotationType = "ROI"
)
