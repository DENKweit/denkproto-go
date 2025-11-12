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
	ClassLabels     []ClassLabelElement `json:"class_labels"`
	Flavor          string              `json:"flavor"`
	NetworkTypename string              `json:"network_typename"`
}

type ClassLabelElement struct {
	ID  string `json:"id"`
	Idx int64  `json:"idx"`
}

type Source struct {
	ClassificationMarkup  *ClassificationMarkup  `json:"classification_markup,omitempty"`
	ObjectDetectionMarkup *ObjectDetectionMarkup `json:"object_detection_markup,omitempty"`
	OcrMarkup             *OCRMarkup             `json:"ocr_markup,omitempty"`
	SegmentationMarkup    *SegmentationMarkup    `json:"segmentation_markup,omitempty"`
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

type SegmentationMarkup struct {
	Annotations         []SegmentationMarkupAnnotation `json:"annotations,omitempty"`
	AverageObjectWidths []float64                      `json:"average_object_widths,omitempty"`
	Height              *int64                         `json:"height,omitempty"`
	Thumbnail           *Thumbnail                     `json:"thumbnail,omitempty"`
	Width               *int64                         `json:"width,omitempty"`
	SegmentationMaps    []SegmentationMap              `json:"segmentation_maps,omitempty"`
}

type SegmentationMarkupAnnotation struct {
	AnnotationType      AnnotationType       `json:"annotation_type"`
	AverageWidth        float64              `json:"average_width"`
	CircleAnnotation    *CircleAnnotation    `json:"circle_annotation,omitempty"`
	ID                  string               `json:"id"`
	LabelID             string               `json:"label_id"`
	MagicwandAnnotation *MagicwandAnnotation `json:"magicwand_annotation,omitempty"`
	PenAnnotation       *PenAnnotation       `json:"pen_annotation,omitempty"`
	PixelAnnotation     *PixelAnnotation     `json:"pixel_annotation,omitempty"`
	PolygonAnnotation   *Polygon             `json:"polygon_annotation,omitempty"`
	RectangleAnnotation *RectangleAnnotation `json:"rectangle_annotation,omitempty"`
	SausageAnnotation   *SausageAnnotation   `json:"sausage_annotation,omitempty"`
}

type CircleAnnotation struct {
	CenterX float64 `json:"center_x"`
	CenterY float64 `json:"center_y"`
	Radius  float64 `json:"radius"`
}

type MagicwandAnnotation struct {
	BottomRightX float64        `json:"bottom_right_x"`
	BottomRightY float64        `json:"bottom_right_y"`
	CenterX      float64        `json:"center_x"`
	CenterY      float64        `json:"center_y"`
	Data         *string        `json:"data,omitempty"`
	DataURL      *string        `json:"dataURL,omitempty"`
	Points       []PointElement `json:"points"`
	Threshold    int64          `json:"threshold"`
	TopLeftX     float64        `json:"top_left_x"`
	TopLeftY     float64        `json:"top_left_y"`
}

type PenAnnotation struct {
	BottomRightX float64        `json:"bottom_right_x"`
	BottomRightY float64        `json:"bottom_right_y"`
	Data         *string        `json:"data,omitempty"`
	DataURL      *string        `json:"dataURL,omitempty"`
	Points       []PointElement `json:"points"`
	Thickness    float64        `json:"thickness"`
	TopLeftX     float64        `json:"top_left_x"`
	TopLeftY     float64        `json:"top_left_y"`
}

type PixelAnnotation struct {
	BlobID       *string `json:"blob_id,omitempty"`
	BottomRightX float64 `json:"bottom_right_x"`
	BottomRightY float64 `json:"bottom_right_y"`
	Data         *string `json:"data,omitempty"`
	TopLeftX     float64 `json:"top_left_x"`
	TopLeftY     float64 `json:"top_left_y"`
}

type RectangleAnnotation struct {
	BottomRightX float64 `json:"bottom_right_x"`
	BottomRightY float64 `json:"bottom_right_y"`
	TopLeftX     float64 `json:"top_left_x"`
	TopLeftY     float64 `json:"top_left_y"`
}

type SausageAnnotation struct {
	BottomRightX float64        `json:"bottom_right_x"`
	BottomRightY float64        `json:"bottom_right_y"`
	Data         *string        `json:"data,omitempty"`
	DataURL      *string        `json:"dataURL,omitempty"`
	Points       []PointElement `json:"points"`
	Radius       float64        `json:"radius"`
	TopLeftX     float64        `json:"top_left_x"`
	TopLeftY     float64        `json:"top_left_y"`
}

type SegmentationMap struct {
	Blob       Blob                      `json:"blob"`
	ClassLabel SegmentationMapClassLabel `json:"class_label"`
}

type Blob struct {
	ID             string `json:"id"`
	OwnedByGroupID string `json:"owned_by_group_id"`
}

type SegmentationMapClassLabel struct {
	ID  string `json:"id"`
	Idx int64  `json:"idx"`
}

// A binary mask with REE (Run-End Encoding) compressed data
type Thumbnail struct {
	Data   string `json:"data"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}

type AnnotationType string

const (
	Ignore   AnnotationType = "IGNORE"
	Negative AnnotationType = "NEGATIVE"
	Positive AnnotationType = "POSITIVE"
	Roi      AnnotationType = "ROI"
)
