package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"os/exec"
)

type configdata struct {
	Hostname    string
	Dns         string
	Displayname string
	Instanceid  string
	Hostdns     string
}

func createTemplate(confdata configdata) error {
	var err error

	newcp := fmt.Sprintf("output/%s.conf", confdata.Instanceid)

	t, err := template.ParseFiles("template/node.tmpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(newcp)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	err = t.Execute(f, confdata)
	if err != nil {
		panic(err)
	}

	return nil
}

func describeTags(instanceID string) AwsDataTags {
	var awstags AwsDataTags

	cmdout, err := exec.Command(aws_cmd, "ec2", "describe-tags", "--filters",
		fmt.Sprintf("Name=resource-id,Values=%s", instanceID)).Output()

	if err != nil {
		panic(err)
	}

	json.Unmarshal(cmdout, &awstags)

	return awstags
}

func listEC2() error {
	var awsdata [][]AwsData
	var configdata configdata

	cmdout, err := exec.Command(aws_cmd, "ec2", "describe-instances",
		"--query", "Reservations[*].Instances[*]",
		"--output=json").Output()

	if err != nil {
		panic(err)
	}

	json.Unmarshal(cmdout, &awsdata)

	for _, data := range awsdata {
		if data[0].State.Name == "running" {
			awsdatatags := describeTags(data[0].InstanceID)
			if len(awsdatatags.Tags) > 1 {
				configdata.Displayname = fmt.Sprintf("%s-%s", awsdatatags.Tags[1].Value, data[0].InstanceID)
			} else {
				configdata.Displayname = fmt.Sprintf("%s-%s", awsdatatags.Tags[0].Value, data[0].InstanceID)
			}

			configdata.Hostname = data[0].PublicDNSName
			configdata.Dns = data[0].PublicDNSName
			configdata.Instanceid = data[0].InstanceID
			configdata.Hostdns = data[0].PublicDNSName

			fmt.Printf("===> Instance: %s\n", configdata.Displayname)

			createTemplate(configdata)

		}
	}

	return nil
}

func main() {
	awsdata := listEC2()
	if awsdata != nil {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
