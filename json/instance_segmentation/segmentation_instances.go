// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    segmentationInstances, err := UnmarshalSegmentationInstances(bytes)
//    bytes, err = segmentationInstances.Marshal()

package instance_segmentation

import "encoding/json"

func UnmarshalSegmentationInstances(data []byte) (SegmentationInstances, error) {
	var r SegmentationInstances
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SegmentationInstances) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SegmentationInstances struct {
	// height of the image for which the segmentation was generated         
	Height                                                         int64    `json:"height"`
	// Instance segmentation objects                                        
	Objects                                                        []Object `json:"objects"`
	// width of the image for which the segmentation was generated          
	Width                                                          int64    `json:"width"`
}

type Object struct {
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
