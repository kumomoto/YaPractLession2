package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type yamlFile struct {
	ApiVersion string     `yaml:"apiVersion"`
	Kind       string     `yaml:"kind"`
	Metadata   ObjectMeta `yaml:"metadata"`
	Spec       PodSpec    `yaml:"spec"`
}

type ObjectMeta struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace,omitempty"`
	Labels    map[string]string `yaml:"labels,omitempty"`
}

type PodSpec struct {
	OS         PodOS       `yaml:"os,omitempty"`
	Containers []Container `yaml:"containers"`
}

type PodOS struct {
	Name string `yaml:"name"`
}

type Container struct {
	Name           string               `yaml:"name"`
	Image          string               `yaml:"image"`
	Ports          []ContainerPort      `yaml:"ports,omitempty"`
	ReadinessProbe Probe                `yaml:"readinessProbe,omitempty"`
	LivenessProbe  Probe                `yaml:"livenessProbe,omitempty"`
	Resources      ResourceRequirements `yaml:"resources,omitempty"`
}

type ContainerPort struct {
	ContainerPort int    `yaml:"containerPort"`
	Protocol      string `yaml:"protocol,omitempty"`
}

type Probe struct {
	HTGet HTTPGetAction `yaml:"httpGet"`
}

type ResourceRequirements struct {
	Limits   Limits   `yaml:"limits,omitempty"`
	Requests Requests `yaml:"requests,omitempty"`
}

type HTTPGetAction struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}

type Requests struct {
	Cpu    int    `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type Limits struct {
	Cpu    int    `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

func main() {

	filePath := os.Args[1]

	if filePath == "" {
		fmt.Errorf("Error! File is not specified.")
	}

	var manifest yamlFile
	var yamlNode yaml.Node

	contentOfFile, fileError := os.ReadFile(filePath)

	if fileError != nil {
		fmt.Errorf("cannot read file content: %w", fileError)
	}

	if unmarshalErr := yaml.Unmarshal(contentOfFile, &manifest); unmarshalErr != nil {
		fmt.Errorf("cannot unmarshal file content: %w", unmarshalErr)
	}

	if unmarshalErr := yaml.Unmarshal(contentOfFile, &yamlNode); unmarshalErr != nil {
		fmt.Errorf("cannot unmarshal file content: %w", unmarshalErr)
	}

	fmt.Println(manifest)

}

func getLine(yamlNode yaml.Node, value string) int {

	var line int

	for _, doc := range yamlNode.Content {
		if strings.Compare(doc.Value, value) == 0 {
			line = doc.Line
		} else {
			line = getLine(yamlNode, value)
		}
	}

	return line
}

func (y *yamlFile) checkReqFileds(yamlNode yaml.Node) map[int]string {
	output := make(map[int]string)

	if y.ApiVersion == "" {
		output[1] = fmt.Sprintf("%s is required", "apiVersion")
	}
	if y.Kind == "" {

	}
	return output
}

func (y *yamlFile) checkTypeOfFields(filePath string, yamlNode yaml.Node) {

}

func (y *yamlFile) validateValuesOfFields(filePath string, yamlNode yaml.Node) {

}

func (y *yamlFile) checkUnsupportedValue(filePath string, yamlNode yaml.Node) {

}

func (y *yamlFile) validateOutOfRange(filePath string, yamlNode yaml.Node) {

}
