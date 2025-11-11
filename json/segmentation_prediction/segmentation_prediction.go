// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    segmentationPrediction, err := UnmarshalSegmentationPrediction(bytes)
//    bytes, err = segmentationPrediction.Marshal()

package segmentation_prediction

import "encoding/json"

func UnmarshalSegmentationPrediction(data []byte) (SegmentationPrediction, error) {
	var r SegmentationPrediction
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SegmentationPrediction) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SegmentationPrediction struct {
	Height                                         int64             `json:"height"`
	// Thresholded segmentation maps for each class                  
	SegmentationMaps                               []SegmentationMap `json:"segmentation_maps"`
	Width                                          int64             `json:"width"`
}

type SegmentationMap struct {
	LabelID string `json:"label_id"`
	Mask    Mask   `json:"mask"`
}

// A binary mask with REE (Run-End Encoding) compressed data
type Mask struct {
	Data   string `json:"data"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}
