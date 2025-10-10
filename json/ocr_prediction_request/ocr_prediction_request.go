// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ocrPredictionRequest, err := UnmarshalOcrPredictionRequest(bytes)
//    bytes, err = ocrPredictionRequest.Marshal()

package ocr_prediction_request

import "encoding/json"

func UnmarshalOcrPredictionRequest(data []byte) (OcrPredictionRequest, error) {
	var r OcrPredictionRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OcrPredictionRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OcrPredictionRequest struct {
	CreatedByUserID    string            `json:"created_by_user_id"`
	ID                 string            `json:"id"`
	Image              Image             `json:"image"`
	NetworkExperiment  NetworkExperiment `json:"network_experiment"`
	Objects            []Object          `json:"objects"`
	OwnedByGroupID     string            `json:"owned_by_group_id"`
	PredictionPriority int64             `json:"prediction_priority"`
}

type Image struct {
	BlobID         string `json:"blob_id"`
	Height         int64  `json:"height"`
	OwnedByGroupID string `json:"owned_by_group_id"`
	Width          int64  `json:"width"`
}

type NetworkExperiment struct {
	ClassLabels              []ClassLabel              `json:"class_labels"`
	Config                   Config                    `json:"config"`
	Flavor                   string                    `json:"flavor"`
	NetworkTypename          string                    `json:"network_typename"`
	OcrCharacterRestrictions []OcrCharacterRestriction `json:"ocr_character_restrictions"`
	Snapshot                 Snapshot                  `json:"snapshot"`
}

type ClassLabel struct {
	ID  string `json:"id"`
	Idx int64  `json:"idx"`
}

type Config struct {
	Metadata             map[string]interface{} `json:"metadata"`
	UsesValidationTiling bool                   `json:"uses_validation_tiling"`
}

type OcrCharacterRestriction struct {
	AllowedCharacters             string                        `json:"allowed_characters"`
	Index                         int64                         `json:"index"`
	NumberOfCharacters            int64                         `json:"number_of_characters"`
	OcrCharacterRestrictionPreset OcrCharacterRestrictionPreset `json:"ocr_character_restriction_preset"`
}

type OcrCharacterRestrictionPreset struct {
	Characters string `json:"characters"`
	Value      string `json:"value"`
}

type Snapshot struct {
	Onnx    Onnx    `json:"onnx"`
	Pytorch Pytorch `json:"pytorch"`
}

type Onnx struct {
	BlobID         string `json:"blob_id"`
	OwnedByGroupID string `json:"owned_by_group_id"`
}

type Pytorch struct {
	BlobID         string `json:"blob_id"`
	OwnedByGroupID string `json:"owned_by_group_id"`
}

type Object struct {
	AnnotationType AnnotationType `json:"annotation_type"`
	AverageWidth   float64        `json:"average_width"`
	BoundingBox    *BoundingBox   `json:"bounding_box,omitempty"`
	ID             string         `json:"id"`
	LabelID        string         `json:"label_id"`
	Polygon        *Polygon       `json:"polygon,omitempty"`
}

type BoundingBox struct {
	Angle           *float64 `json:"angle,omitempty"`
	BottomRightX    float64  `json:"bottom_right_x"`
	BottomRightY    float64  `json:"bottom_right_y"`
	FullOrientation *bool    `json:"full_orientation,omitempty"`
	TopLeftX        float64  `json:"top_left_x"`
	TopLeftY        float64  `json:"top_left_y"`
}

// A polygon defined by one or more rings, allowing for holes and nested structures.
type Polygon struct {
	// Array of polygon rings. The hierarchy field within each ring determines nesting and                             
	// fill/hole status.                                                                                               
	Rings                                                                                 []OcrPredictionRequestSchema `json:"rings"`
}

// A single closed loop (ring) of a polygon, defining either an outer boundary or a hole.
type OcrPredictionRequestSchema struct {
	// Nesting level: 0=outer, 1=hole in level 0, 2=poly in level 1 hole, etc. Even levels are        
	// filled areas, odd levels are holes.                                                            
	Hierarchy                                                                                 int64   `json:"hierarchy"`
	// Vertices of the ring.                                                                          
	Points                                                                                    []Point `json:"points"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type AnnotationType string

const (
	Ignore   AnnotationType = "IGNORE"
	Negative AnnotationType = "NEGATIVE"
	Positive AnnotationType = "POSITIVE"
	Roi      AnnotationType = "ROI"
)
