package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	. "optics/pkg/models"
	"strings"
)

type OpticsObject struct {
	OpticDetails
	Devices []string `json:"devices"`
}

type OpticsProcessor struct {
	speed     string
	sku       string
	cableType string
	distance  string

	possibleOptics []string
	selectedOptics []string
}

func NewOpticsProcessor() *OpticsProcessor {
	// modelTypesOfInterest := []string{"100 Gigabit Ethernet", "400 Gigabit Ethernet", "4 x 25 Gigabit Ethernet", "4 x 100 Gigabit Ethernet"}
	return &OpticsProcessor{}
}

func (op *OpticsProcessor) SetSKU(deviceSKU string) {
	op.sku = deviceSKU
}

func (op *OpticsProcessor) SetCableType(cableType string) {
	op.cableType = cableType
}

func (op *OpticsProcessor) SetDistance(distance string) {
	op.distance = distance
}

func (op *OpticsProcessor) SetSpeed(speed string) {
	op.speed = speed
}

func (op *OpticsProcessor) GetPossibleOptics() []string {
	return op.possibleOptics
}

func (op *OpticsProcessor) GetSelectedOptics() []string {
	return op.selectedOptics
}

func (op *OpticsProcessor) appendPossibleOptics(opticsModel string) {
	if !ContainsKey(opticsModel, op.possibleOptics) {
		op.possibleOptics = append(op.possibleOptics, opticsModel)
	}
}

func (op *OpticsProcessor) appendSelectedOptics(opticsModel string) {
	if !ContainsKey(opticsModel, op.selectedOptics) {
		op.selectedOptics = append(op.selectedOptics, opticsModel)
	}
}

func (op *OpticsProcessor) processOpticsToDevicesMapResp() {
	resp, err := ioutil.ReadFile("optics2devicesmap.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(resp))
	opticsToDevicesResp := &[]OpticsToDevicesResp{}
	json.Unmarshal([]byte(resp), opticsToDevicesResp)
	// fmt.Print(opticsToDevicesResp)
	for _, modelTypeToDevicesMap := range *opticsToDevicesResp {
		if ContainsKey(modelTypeToDevicesMap.ModelType, []string{}) {
			// fmt.Println(modelTypeToDevicesMap)
			for _, model := range modelTypeToDevicesMap.Models {
				op.processDeviceModels(model)
			}
		}
	}

}

func (op *OpticsProcessor) processDeviceModels(model DeviceModelDetails) {
	// fmt.Println(model.ModelNum)
	if ContainsKey(model.ModelNum, op.possibleOptics) {
		platforms := model.SupportedPlatforms
		// check based on family and process.
		for _, sw := range platforms.Switching {
			if sw.ProdName == op.sku {
				// fmt.Println(sw.ProdName, model.ModelNum)
				op.appendSelectedOptics(model.ModelNum)
			}
		}
		for _, rt := range platforms.Routing {
			if rt.ProdName == op.sku {
				// fmt.Println(rt.ProdName, model.ModelNum)
				op.appendSelectedOptics(model.ModelNum)
			}
		}
	}
}

func getAllOptics() OpticsResp {
	link := "https://apps.juniper.net/hct/optics/"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))

	opticsResp := &OpticsResp{}
	json.Unmarshal([]byte(s), opticsResp)

	return *opticsResp
}

func (op *OpticsProcessor) processOptics(optics []OpticDetails) {
	for _, optic := range optics {
		if ContainsKey(optic.CableType, []string{op.cableType}) &&
			IsDistanceSupported(optic.Distance) {
			// fmt.Println(optic.ModelNum, optic.Distance, optic.TransceiverType, optic.CableType)
			op.appendPossibleOptics(optic.ModelNum)
		}
		// if IsDistanceSupported(optic.Distance) {
		// 	fmt.Println(optic.ModelNum, optic.Distance, optic.TransceiverType, optic.CableType)
		// }
	}
}

func (op *OpticsProcessor) processOpticResponse() {
	resp, err := ioutil.ReadFile("optics.txt")
	// fmt.Println(string(resp))
	if err != nil {
		fmt.Println(err)
	}
	opticsResp := &OpticsResp{}
	json.Unmarshal([]byte(resp), opticsResp)
	// fmt.Print(opticsResp)

	for _, modelType := range opticsResp.OpticsResp {
		if ContainsKey(modelType.ModelType, []string{}) {
			// Process optics for this model type
			// fmt.Println(modelType.ModelType)
			op.processOptics(modelType.Optics)
		}
	}
}

func (op *OpticsProcessor) GetOpticsWithCableTypeAndSpeed() {
	resp, err := ioutil.ReadFile("optics.txt")
	// fmt.Println(string(resp))
	if err != nil {
		fmt.Println(err)
	}
	opticsResp := &OpticsResp{}
	json.Unmarshal([]byte(resp), opticsResp)
	// fmt.Print(opticsResp)

	for _, modelType := range opticsResp.OpticsResp {
		if ContainsKey(modelType.ModelType, []string{op.speed}) {
			// Process optics for this model type
			// fmt.Println(modelType.ModelType)
			op.processOptics(modelType.Optics)
		}
	}

	// Now call to get the Optics
	op.GetOptics()

}

func (op *OpticsProcessor) GetOptics() {
	resp, err := ioutil.ReadFile("optics2devicesmap.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(resp))
	opticsToDevicesResp := &[]OpticsToDevicesResp{}
	json.Unmarshal([]byte(resp), opticsToDevicesResp)
	// fmt.Print(opticsToDevicesResp)
	for _, modelTypeToDevicesMap := range *opticsToDevicesResp {
		if ContainsKey(modelTypeToDevicesMap.ModelType, []string{op.speed}) {
			// fmt.Println(modelTypeToDevicesMap)
			for _, model := range modelTypeToDevicesMap.Models {
				op.processDeviceModels(model)
			}
		}
	}

}
