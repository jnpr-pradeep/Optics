package models

type OpticsResp struct {
	OpticsResp []ModelTypeDetails `json:"data"`
}

type ModelTypeDetails struct {
	ModelType string         `json:"modelType"`
	Optics    []OpticDetails `json:"optics"`
}

type OpticDetails struct {
	ModelNum        string `json:"modelNum"`
	Desc            string `json:"desc"`
	PartNumber      string `json:"partNumber"`
	ConnectorType   string `json:"connectorType"`
	ProductType     string `json:"productType"`
	CableType       string `json:"cableType"`
	Speed           string `json:"speed"`
	Distance        string `json:"distance"`
	ModelType       string `json:"modelType"`
	BreakoutCapable string `json:"breakoutCapable"`
	TransceiverType string `json:"transceiverType"`
	EOLFlag         string `json:"eolFlag"`
	Standard        string `json:"standard"`
}

type OpticsToDevicesResp struct {
	ModelType string               `json:"modelType`
	Models    []DeviceModelDetails `json:"models"`
}

type DeviceModelDetails struct {
	ModelNum           string       `json:"modelNum"`
	SupportedPlatforms PlatformsMap `json:"supportedPlatforms"`
}

type PlatformsMap struct {
	Security  []ProductInfo `json:"Security"`
	Routing   []ProductInfo `json:"Routing"`
	Switching []ProductInfo `json:"Switching"`
}

type ProductInfo struct {
	ProdName string `json:"prodName"`
}
