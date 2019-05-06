/*-
 * Copyright 2019 Bitmark, Inc.
 * Copyright 2019 by Marcelo Araujo <araujo@FreeBSD.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted providing that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

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

	fmt.Printf("===> Instance: %s\n", confdata.Displayname)
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

	cmdout, err := exec.Command(awscmd, "ec2", "describe-tags", "--filters",
		fmt.Sprintf("Name=resource-id,Values=%s", instanceID)).Output()

	if err != nil {
		panic(err)
	}

	json.Unmarshal(cmdout, &awstags)

	return awstags
}

func queueConsumer(awsdata [][]AwsData, awsdata_cp chan [][]AwsData) {
	if len(awsdata) > 1 {
		awsdata = awsdata[1:]
		awsdata_cp <- awsdata
	}
}

func createConfig(awsdata [][]AwsData) {
	var configdata configdata

	awsdata_cp := make(chan [][]AwsData)
	for {
		if len(awsdata) <= 1 {
			break
		}

		go queueConsumer(awsdata, awsdata_cp)

		if awsdata_cp == nil {
			break
		}

		awsdata = <-awsdata_cp
		if awsdata[0][0].State.Name == "running" {
			awsdatags := describeTags(awsdata[0][0].InstanceID)
			if len(awsdatags.Tags) > 1 {
				configdata.Displayname = fmt.Sprintf("%s-%s",
					awsdatags.Tags[1].Value, awsdata[0][0].InstanceID)
			} else {
				configdata.Displayname = fmt.Sprintf("%s-%s",
					awsdatags.Tags[0].Value, awsdata[0][0].InstanceID)
			}
			configdata.Hostname = awsdata[0][0].PublicDNSName
			configdata.Dns = awsdata[0][0].PublicDNSName
			configdata.Instanceid = awsdata[0][0].InstanceID
			configdata.Hostname = awsdata[0][0].PublicDNSName

			go createTemplate(configdata)
		}
	}
}

func listEC2() [][]AwsData {
	var awsdata [][]AwsData

	cmdout, err := exec.Command(awscmd, "ec2", "describe-instances",
		"--query", "Reservations[*].Instances[*]",
		"--output=json").Output()

	if err != nil {
		panic(err)
	}

	json.Unmarshal(cmdout, &awsdata)

	return awsdata
}

func main() {
	awsdata := listEC2()
	if awsdata != nil {
		createConfig(awsdata)
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
