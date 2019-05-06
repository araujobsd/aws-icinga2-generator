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
	"time"
)

const (
	awscmd = "/usr/bin/aws"
)

// AwsDataTags - Get all tags structure
type AwsDataTags struct {
	Tags []struct {
		Key          string `json:"Key"`
		ResourceID   string `json:"ResourceId"`
		ResourceType string `json:"ResourceType"`
		Value        string `json:"Value"`
	} `json:"Tags"`
}

// AwsData - Get all instances information
type AwsData struct {
	AmiLaunchIndex int       `json:"AmiLaunchIndex"`
	ImageID        string    `json:"ImageId"`
	InstanceID     string    `json:"InstanceId"`
	InstanceType   string    `json:"InstanceType"`
	KeyName        string    `json:"KeyName"`
	LaunchTime     time.Time `json:"LaunchTime"`
	Monitoring     struct {
		State string `json:"State"`
	} `json:"Monitoring"`
	Placement struct {
		AvailabilityZone string `json:"AvailabilityZone"`
		GroupName        string `json:"GroupName"`
		Tenancy          string `json:"Tenancy"`
	} `json:"Placement"`
	PrivateDNSName   string        `json:"PrivateDnsName"`
	PrivateIPAddress string        `json:"PrivateIpAddress"`
	ProductCodes     []interface{} `json:"ProductCodes"`
	PublicDNSName    string        `json:"PublicDnsName"`
	State            struct {
		Code int    `json:"Code"`
		Name string `json:"Name"`
	} `json:"State"`
	StateTransitionReason string `json:"StateTransitionReason"`
	SubnetID              string `json:"SubnetId"`
	VpcID                 string `json:"VpcId"`
	Architecture          string `json:"Architecture"`
	BlockDeviceMappings   []struct {
		DeviceName string `json:"DeviceName"`
		Ebs        struct {
			AttachTime          time.Time `json:"AttachTime"`
			DeleteOnTermination bool      `json:"DeleteOnTermination"`
			Status              string    `json:"Status"`
			VolumeID            string    `json:"VolumeId"`
		} `json:"Ebs"`
	} `json:"BlockDeviceMappings"`
	ClientToken       string `json:"ClientToken"`
	EbsOptimized      bool   `json:"EbsOptimized"`
	EnaSupport        bool   `json:"EnaSupport"`
	Hypervisor        string `json:"Hypervisor"`
	NetworkInterfaces []struct {
		Attachment struct {
			AttachTime          time.Time `json:"AttachTime"`
			AttachmentID        string    `json:"AttachmentId"`
			DeleteOnTermination bool      `json:"DeleteOnTermination"`
			DeviceIndex         int       `json:"DeviceIndex"`
			Status              string    `json:"Status"`
		} `json:"Attachment"`
		Description string `json:"Description"`
		Groups      []struct {
			GroupName string `json:"GroupName"`
			GroupID   string `json:"GroupId"`
		} `json:"Groups"`
		Ipv6Addresses      []interface{} `json:"Ipv6Addresses"`
		MacAddress         string        `json:"MacAddress"`
		NetworkInterfaceID string        `json:"NetworkInterfaceId"`
		OwnerID            string        `json:"OwnerId"`
		PrivateDNSName     string        `json:"PrivateDnsName"`
		PrivateIPAddress   string        `json:"PrivateIpAddress"`
		PrivateIPAddresses []struct {
			Primary          bool   `json:"Primary"`
			PrivateDNSName   string `json:"PrivateDnsName"`
			PrivateIPAddress string `json:"PrivateIpAddress"`
		} `json:"PrivateIpAddresses"`
		SourceDestCheck bool   `json:"SourceDestCheck"`
		Status          string `json:"Status"`
		SubnetID        string `json:"SubnetId"`
		VpcID           string `json:"VpcId"`
	} `json:"NetworkInterfaces"`
	RootDeviceName string `json:"RootDeviceName"`
	RootDeviceType string `json:"RootDeviceType"`
	SecurityGroups []struct {
		GroupName string `json:"GroupName"`
		GroupID   string `json:"GroupId"`
	} `json:"SecurityGroups"`
	SourceDestCheck bool `json:"SourceDestCheck"`
	StateReason     struct {
		Code    string `json:"Code"`
		Message string `json:"Message"`
	} `json:"StateReason"`
	Tags []struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"Tags"`
	VirtualizationType string `json:"VirtualizationType"`
	CPUOptions         struct {
		CoreCount      int `json:"CoreCount"`
		ThreadsPerCore int `json:"ThreadsPerCore"`
	} `json:"CpuOptions"`
}
