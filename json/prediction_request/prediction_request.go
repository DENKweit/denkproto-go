// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    predictionRequest, err := UnmarshalPredictionRequest(bytes)
//    bytes, err = predictionRequest.Marshal()

package prediction_request

import "encoding/json"

func UnmarshalPredictionRequest(data []byte) (PredictionRequest, error) {
	var r PredictionRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PredictionRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PredictionRequest struct {
	CreatedByUserID                                                  string            `json:"created_by_user_id"`
	HasuraURL                                                        string            `json:"hasura_url"`
	ID                                                               string            `json:"id"`
	Image                                                            Image             `json:"image"`
	NetworkExperiment                                                NetworkExperiment `json:"network_experiment"`
	// Only present for OCR prediction requests                                        
	Objects                                                          []ObjectElement   `json:"objects,omitempty"`
	OwnedByGroupID                                                   string            `json:"owned_by_group_id"`
	PredictionPriority                                               int64             `json:"prediction_priority"`
	// Only present for standard prediction requests                                   
	RequestClassificationInterpretation                              *bool             `json:"request_classification_interpretation,omitempty"`
	// Discriminator field to identify the type of prediction request                  
	RequestType                                                      RequestType       `json:"request_type"`
}

type Image struct {
	BlobID         string `json:"blob_id"`
	FileID         string `json:"file_id"`
	Height         int64  `json:"height"`
	OwnedByGroupID string `json:"owned_by_group_id"`
	Width          int64  `json:"width"`
}

type NetworkExperiment struct {
	ClassLabels                                []ClassLabelElement              `json:"class_labels"`
	Config                                     Config                           `json:"config"`
	Flavor                                     string                           `json:"flavor"`
	ID                                         string                           `json:"id"`
	NetworkTypename                            string                           `json:"network_typename"`
	// Only present for OCR prediction requests                                 
	OcrCharacterRestrictions                   []OcrCharacterRestrictionElement `json:"ocr_character_restrictions,omitempty"`
	Snapshot                                   Snapshot                         `json:"snapshot"`
}

type ClassLabelElement struct {
	ID  string `json:"id"`
	Idx int64  `json:"idx"`
}

type Config struct {
	Metadata             map[string]interface{} `json:"metadata"`
	UsesValidationTiling bool                   `json:"uses_validation_tiling"`
}

type OcrCharacterRestrictionElement struct {
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
	ID      string `json:"id"`
	Onnx    Onnx   `json:"onnx"`
	Pytorch Onnx   `json:"pytorch"`
}

type Onnx struct {
	BlobID         string `json:"blob_id"`
	OwnedByGroupID string `json:"owned_by_group_id"`
}

type ObjectElement struct {
	AnnotationType AnnotationType `json:"annotation_type"`
	AverageWidth   float64        `json:"average_width"`
	BoundingBox    *BoundingBox   `json:"bounding_box,omitempty"`
	ID             string         `json:"id"`
	LabelID        string         `json:"label_id"`
	Polygon        *Polygon       `json:"polygon,omitempty"`
}

// A bounding box with optional rotation information
type BoundingBox struct {
	// Optional rotation angle                
	Angle                            *float64 `json:"angle,omitempty"`
	BottomRightX                     float64  `json:"bottom_right_x"`
	BottomRightY                     float64  `json:"bottom_right_y"`
	// Optional full orientation flag         
	FullOrientation                  *bool    `json:"full_orientation,omitempty"`
	TopLeftX                         float64  `json:"top_left_x"`
	TopLeftY                         float64  `json:"top_left_y"`
}

// A polygon defined by one or more rings, allowing for holes and nested structures.
type Polygon struct {
	// Array of polygon rings. The hierarchy field within each ring determines nesting and                 
	// fill/hole status.                                                                                   
	Rings                                                                                 []GeometrySchema `json:"rings"`
}

// A single closed loop (ring) of a polygon, defining either an outer boundary or a hole.
type GeometrySchema struct {
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

// Discriminator field to identify the type of prediction request
type RequestType string

const (
	Ocr      RequestType = "ocr"
	Standard RequestType = "standard"
)
