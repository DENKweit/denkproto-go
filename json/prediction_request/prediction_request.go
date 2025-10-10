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
	CreatedByUserID                     string            `json:"created_by_user_id"`
	ID                                  string            `json:"id"`
	Image                               Image             `json:"image"`
	NetworkExperiment                   NetworkExperiment `json:"network_experiment"`
	OwnedByGroupID                      string            `json:"owned_by_group_id"`
	PredictionPriority                  int64             `json:"prediction_priority"`
	RequestClassificationInterpretation bool              `json:"request_classification_interpretation"`
}

type Image struct {
	BlobID         string `json:"blob_id"`
	Height         int64  `json:"height"`
	OwnedByGroupID string `json:"owned_by_group_id"`
	Width          int64  `json:"width"`
}

type NetworkExperiment struct {
	ClassLabels     []ClassLabel `json:"class_labels"`
	Config          Config       `json:"config"`
	Flavor          string       `json:"flavor"`
	NetworkTypename string       `json:"network_typename"`
	Snapshot        Snapshot     `json:"snapshot"`
}

type ClassLabel struct {
	ID  string `json:"id"`
	Idx int64  `json:"idx"`
}

type Config struct {
	Metadata             map[string]interface{} `json:"metadata"`
	UsesValidationTiling bool                   `json:"uses_validation_tiling"`
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
