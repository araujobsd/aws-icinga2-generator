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
