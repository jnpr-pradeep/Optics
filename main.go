package main

import (
	. "clover/optics/models"
	. "clover/optics/utils"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// processOpticsResp(getAllOptics())
	// getOptics2ModelMapping()

	// Intersted Model Types
	o := NewOpticsProcessor()

	o.processOpticResponse()
	o.processOpticsToDevicesMapResp()
}

type OpticsObject struct {
	OpticDetails
	Devices []string `json:"devices"`
}

type OpticsProcessor struct {
	modelTypesOfInterest []string
	speed2OpticsMap      map[string]OpticsObject
}

func NewOpticsProcessor() *OpticsProcessor {
	modelTypesOfInterest := []string{"100 Gigabit Ethernet", "400 Gigabit Ethernet", "4 x 25 Gigabit Ethernet", "4 x 100 Gigabit Ethernet"}

	return &OpticsProcessor{modelTypesOfInterest: modelTypesOfInterest}
}

func (op *OpticsProcessor) processOpticsToDevicesMapResp() {
	resp, err := ioutil.ReadFile("/tmp/clover_optics/optics2devicesmap.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(resp))
	opticsToDevicesResp := &[]OpticsToDevicesResp{}
	json.Unmarshal([]byte(resp), opticsToDevicesResp)
	// fmt.Print(opticsToDevicesResp)
	for _, modelTypeToDevicesMap := range *opticsToDevicesResp {
		if ContainsKey(modelTypeToDevicesMap.ModelType, op.modelTypesOfInterest) {
			// fmt.Println(modelTypeToDevicesMap)
			for _, model := range modelTypeToDevicesMap.Models {
				op.processDeviceModels(model)
			}
		}
	}

}

func (op *OpticsProcessor) processDeviceModels(model DeviceModelDetails) {
	// fmt.Println(model.ModelNum)
	platforms := model.SupportedPlatforms
	// check based on family and process.
	for _, sw := range platforms.Switching {
		fmt.Println(sw.ProdName)
	}
	for _, rt := range platforms.Routing {
		fmt.Println(rt.ProdName)
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
		if IsDistanceSupported(optic.Distance) {
			fmt.Println(optic.ModelNum, optic.Distance, optic.TransceiverType)
		}
	}
}

func (op *OpticsProcessor) processOpticResponse() {
	resp, err := ioutil.ReadFile("/tmp/clover_optics/optics.txt")
	// fmt.Println(string(resp))
	if err != nil {
		fmt.Println(err)
	}
	opticsResp := &OpticsResp{}
	json.Unmarshal([]byte(resp), opticsResp)
	// fmt.Print(opticsResp)

	for _, modelType := range opticsResp.OpticsResp {
		if ContainsKey(modelType.ModelType, op.modelTypesOfInterest) {
			// Process optics for this model type
			fmt.Println(modelType.ModelType)
			op.processOptics(modelType.Optics)
		}
	}
}
