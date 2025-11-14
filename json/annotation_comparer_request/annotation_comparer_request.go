// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    annotationComparerRequest, err := UnmarshalAnnotationComparerRequest(bytes)
//    bytes, err = annotationComparerRequest.Marshal()

package annotation_comparer_request

import "encoding/json"

func UnmarshalAnnotationComparerRequest(data []byte) (AnnotationComparerRequest, error) {
	var r AnnotationComparerRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AnnotationComparerRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AnnotationComparerRequest struct {
	CreatedByUserID   string            `json:"created_by_user_id"`
	HasuraURL         string            `json:"hasura_url"`
	ID                string            `json:"id"`
	Image             Image             `json:"image"`
	NetworkExperiment NetworkExperiment `json:"network_experiment"`
	OwnedByGroupID    string            `json:"owned_by_group_id"`
	Source            Source            `json:"source"`
	Target            Source            `json:"target"`
	User1ID           *string           `json:"user1_id"`
	User2ID           *string           `json:"user2_id"`
}

type Image struct {
	BlobID         string `json:"blob_id"`
	Height         int64  `json:"height"`
	OwnedByGroupID string `json:"owned_by_group_id"`
	Width          int64  `json:"width"`
}

type NetworkExperiment struct {
	ClassLabels     []ClassLabel `json:"class_labels"`
	Flavor          string       `json:"flavor"`
	NetworkTypename string       `json:"network_typename"`
}

type ClassLabel struct {
	ID  string `json:"id"`
	Idx int64  `json:"idx"`
}

type Source struct {
	Data ClassificationMarkup `json:"data"`
	Type Type                 `json:"type"`
}

type ClassificationMarkup struct {
	Annotations                                                    []Annotation `json:"annotations,omitempty"`
	// height of the image for which the segmentation was generated             
	Height                                                         *int64       `json:"height,omitempty"`
	// width of the image for which the segmentation was generated              
	Width                                                          *int64       `json:"width,omitempty"`
	// Instance segmentation objects                                            
	//                                                                          
	// Thresholded segmentation maps for each class                             
	Objects                                                        []Object     `json:"objects,omitempty"`
	AverageObjectWidths                                            []float64    `json:"average_object_widths,omitempty"`
	Predictions                                                    []Prediction `json:"predictions,omitempty"`
}

type Annotation struct {
	ID              string          `json:"id"`
	LabelID         string          `json:"label_id"`
	Value           *float64        `json:"value,omitempty"`
	Angle           *float64        `json:"angle,omitempty"`
	AnnotationType  *AnnotationType `json:"annotation_type,omitempty"`
	AverageWidth    *float64        `json:"average_width,omitempty"`
	BottomRightX    *float64        `json:"bottom_right_x,omitempty"`
	BottomRightY    *float64        `json:"bottom_right_y,omitempty"`
	FullOrientation *bool           `json:"full_orientation,omitempty"`
	TopLeftX        *float64        `json:"top_left_x,omitempty"`
	TopLeftY        *float64        `json:"top_left_y,omitempty"`
	BoundingBox     *BoundingBox    `json:"bounding_box,omitempty"`
	Polygon         *Polygon        `json:"polygon,omitempty"`
	Text            *string         `json:"text,omitempty"`
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
	Rings                                                                                 []RingElement `json:"rings"`
}

// A single closed loop (ring) of a polygon, defining either an outer boundary or a hole.
type RingElement struct {
	// Nesting level: 0=outer, 1=hole in level 0, 2=poly in level 1 hole, etc. Even levels are               
	// filled areas, odd levels are holes.                                                                   
	Hierarchy                                                                                 int64          `json:"hierarchy"`
	Points                                                                                    []PointElement `json:"points"`
}

// An array of points
type PointElement struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Object struct {
	// references the annotation in the segmentation markup that was used to generate this         
	// binary mask                                                                                 
	AnnotationID                                                                          *string  `json:"annotation_id,omitempty"`
	Data                                                                                  string   `json:"data"`
	LabelID                                                                               string   `json:"label_id"`
	// x offset of the object                                                                      
	X                                                                                     int64    `json:"x"`
	// y offset of the object                                                                      
	Y                                                                                     int64    `json:"y"`
	Probability                                                                           *float64 `json:"probability,omitempty"`
}

type Prediction struct {
	InterpretationMap    *InterpretationMap    `json:"interpretation_map,omitempty"`
	LabelID              string                `json:"label_id"`
	Probability          *float64              `json:"probability,omitempty"`
	Angle                *float64              `json:"angle,omitempty"`
	BottomRightX         *float64              `json:"bottom_right_x,omitempty"`
	BottomRightY         *float64              `json:"bottom_right_y,omitempty"`
	FullOrientation      *bool                 `json:"full_orientation,omitempty"`
	TopLeftX             *float64              `json:"top_left_x,omitempty"`
	TopLeftY             *float64              `json:"top_left_y,omitempty"`
	BoundingBox          *BoundingBox          `json:"bounding_box,omitempty"`
	CharacterPredictions []CharacterPrediction `json:"character_predictions,omitempty"`
	Polygon              *Polygon              `json:"polygon,omitempty"`
	Text                 *string               `json:"text,omitempty"`
}

type CharacterPrediction struct {
	Character   string  `json:"character"`
	Probability float64 `json:"probability"`
}

type InterpretationMap struct {
	BlobID  string `json:"blob_id"`
	GroupID string `json:"group_id"`
}

type AnnotationType string

const (
	Ignore   AnnotationType = "IGNORE"
	Negative AnnotationType = "NEGATIVE"
	Positive AnnotationType = "POSITIVE"
	Roi      AnnotationType = "ROI"
)

type Type string

const (
	Classification       Type = "classification"
	InstanceSegmentation Type = "instance_segmentation"
	ObjectDetection      Type = "object_detection"
	Ocr                  Type = "ocr"
	Segmentation         Type = "segmentation"
)
