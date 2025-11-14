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
	ClassificationMarkup           *ClassificationMarkup         `json:"classification_markup,omitempty"`
	ObjectDetectionMarkup          *ObjectDetectionMarkup        `json:"object_detection_markup,omitempty"`
	OcrMarkup                      *OCRMarkup                    `json:"ocr_markup,omitempty"`
	SegmentationInstances          *SegmentationInstances        `json:"segmentation_instances,omitempty"`
	ClassificationPrediction       *ClassificationPrediction     `json:"classification_prediction,omitempty"`
	InstanceSegmentationPrediction *SegmentationMarkupPrediction `json:"instance_segmentation_prediction,omitempty"`
	ObjectDetectionPrediction      *ObjectDetectionPrediction    `json:"object_detection_prediction,omitempty"`
	OcrPrediction                  *OCRPrediction                `json:"ocr_prediction,omitempty"`
}

type ClassificationMarkup struct {
	Annotations []ClassificationMarkupAnnotation `json:"annotations"`
	Height      int64                            `json:"height"`
	Width       int64                            `json:"width"`
}

type ClassificationMarkupAnnotation struct {
	ID      string  `json:"id"`
	LabelID string  `json:"label_id"`
	Value   float64 `json:"value"`
}

type ClassificationPrediction struct {
	Predictions []ClassificationPredictionPrediction `json:"predictions"`
}

type ClassificationPredictionPrediction struct {
	InterpretationMap *InterpretationMap `json:"interpretation_map,omitempty"`
	LabelID           string             `json:"label_id"`
	Probability       float64            `json:"probability"`
}

type InterpretationMap struct {
	Data    string `json:"data"`
	GroupID string `json:"group_id"`
}

type SegmentationMarkupPrediction struct {
	Height                                         int64                                  `json:"height"`
	// Thresholded segmentation maps for each class                                       
	Objects                                        []InstanceSegmentationPredictionObject `json:"objects"`
	Width                                          int64                                  `json:"width"`
}

type InstanceSegmentationPredictionObject struct {
	Data        string  `json:"data"`
	LabelID     string  `json:"label_id"`
	Probability float64 `json:"probability"`
	X           int64   `json:"x"`
	Y           int64   `json:"y"`
}

type ObjectDetectionMarkup struct {
	Annotations []ObjectDetectionMarkupAnnotation `json:"annotations"`
	Height      int64                             `json:"height"`
	Width       int64                             `json:"width"`
}

type ObjectDetectionMarkupAnnotation struct {
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

type ObjectDetectionPrediction struct {
	Height      int64                                 `json:"height"`
	Predictions []ObjectDetectionPredictionPrediction `json:"predictions"`
	Width       int64                                 `json:"width"`
}

type ObjectDetectionPredictionPrediction struct {
	Angle           *float64 `json:"angle,omitempty"`
	BottomRightX    float64  `json:"bottom_right_x"`
	BottomRightY    float64  `json:"bottom_right_y"`
	FullOrientation *bool    `json:"full_orientation,omitempty"`
	LabelID         string   `json:"label_id"`
	Probability     float64  `json:"probability"`
	TopLeftX        float64  `json:"top_left_x"`
	TopLeftY        float64  `json:"top_left_y"`
}

type OCRMarkup struct {
	Annotations         []OcrMarkupAnnotation `json:"annotations"`
	AverageObjectWidths []float64             `json:"average_object_widths"`
	Height              int64                 `json:"height"`
	Width               int64                 `json:"width"`
}

type OcrMarkupAnnotation struct {
	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`
	ID          string       `json:"id"`
	LabelID     string       `json:"label_id"`
	Polygon     *Polygon     `json:"polygon,omitempty"`
	Text        string       `json:"text"`
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

type OCRPrediction struct {
	Height      int64                     `json:"height"`
	Predictions []OcrPredictionPrediction `json:"predictions"`
	Width       int64                     `json:"width"`
}

type OcrPredictionPrediction struct {
	BoundingBox          *BoundingBox          `json:"bounding_box,omitempty"`
	CharacterPredictions []CharacterPrediction `json:"character_predictions"`
	LabelID              string                `json:"label_id"`
	Polygon              *Polygon              `json:"polygon,omitempty"`
	Text                 string                `json:"text"`
}

type CharacterPrediction struct {
	Character   string  `json:"character"`
	Probability float64 `json:"probability"`
}

type SegmentationInstances struct {
	// height of the image for which the segmentation was generated                              
	Height                                                         int64                         `json:"height"`
	// Instance segmentation objects                                                             
	Objects                                                        []SegmentationInstancesObject `json:"objects"`
	// width of the image for which the segmentation was generated                               
	Width                                                          int64                         `json:"width"`
}

type SegmentationInstancesObject struct {
	// references the annotation in the segmentation markup that was used to generate this       
	// binary mask                                                                               
	AnnotationID                                                                          string `json:"annotation_id"`
	Data                                                                                  string `json:"data"`
	LabelID                                                                               string `json:"label_id"`
	// x offset of the object                                                                    
	X                                                                                     int64  `json:"x"`
	// y offset of the object                                                                    
	Y                                                                                     int64  `json:"y"`
}

type AnnotationType string

const (
	Ignore   AnnotationType = "IGNORE"
	Negative AnnotationType = "NEGATIVE"
	Positive AnnotationType = "POSITIVE"
	Roi      AnnotationType = "ROI"
)
