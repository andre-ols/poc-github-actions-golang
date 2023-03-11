package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GetDeviceFields(dev_description string) map[string]string {
	// The fields are divided by a new line
	//fmt.Println(dev_description)
	fields_list := strings.Split(dev_description, "\r\n")

	parsed_devinfo := make(map[string]string)
	var field_split []string
	var key string
	var value string

	for _, field := range fields_list {
		field_split = strings.Split(field, ":")
		if len(field_split) >= 2 {
			//fmt.Println(field_split)

			// Standardize the spaces and adds the key/value pair
			key = strings.Join(strings.Fields(field_split[0]), " ")

			if len(key) > 0 {

				value = strings.Join(strings.Fields(field_split[1]), " ")
				parsed_devinfo[string(key)] = string(value)

			}

		}

	}

	return parsed_devinfo
}

func FilterFields(parsed_devinfo map[string]string) map[string]string {

	important_fields := [...]string{"Name", "Status", "Class", "ClassGuid", "HardwareID", "Description"}
	filtered_dev_list := make(map[string]string)

	value := ""
	found := false

	if len(parsed_devinfo) > 0 {
		for _, key := range important_fields {

			value, found = parsed_devinfo[key]

			if found == true {
				filtered_dev_list[key] = value
			}
		}
	}
	return filtered_dev_list
}

func ParseWinDeviceList(device_list []string) []map[string]string {

	var parsed_dev_list []map[string]string

	for idx, dev := range device_list {
		// the first lines are usually empty
		if idx > 1 {
			parsed_devinfo := GetDeviceFields(dev)
			parsed_devinfo = FilterFields(parsed_devinfo)
			if len(parsed_devinfo) > 0 {
				parsed_dev_list = append(parsed_dev_list, parsed_devinfo)
			}
		}
	}

	return parsed_dev_list
}

func main() {
	cmd := exec.Command("powershell", "Get-PnpDevice -PresentOnly | Where-Object { $_.InstanceId -match '^USB' } | Format-List")
	//cmd1 := exec.Command("powershell", "Get-PnpDevice -PresentOnly")

	res, _ := cmd.Output()

	output := string(res)

	output_list := strings.Split(output, "\r\n\r\n")

	result := ParseWinDeviceList(output_list)

	/*for key := range result[1] {
		fmt.Print(key + " :")
		fmt.Print(" " + result[1][key] + "\n")
	}*/

	fmt.Print("Número de Devices Conectados ao computador: ")
	fmt.Println(len(result))

	write_on_file, _ := json.Marshal(result)
	_ = ioutil.WriteFile("devices.json", write_on_file, 0644)
}
