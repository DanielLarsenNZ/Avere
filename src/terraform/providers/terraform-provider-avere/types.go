// Copyright (C) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See LICENSE-CODE in the project root for license information.
package main

import (
	"golang.org/x/crypto/ssh"
)

type IaasPlatform interface {
	CreateVfxt(avereVfxt *AvereVfxt) error
	AddIaasNodeToCluster(avereVfxt *AvereVfxt) error
	DestroyVfxt(avereVfxt *AvereVfxt) error
	DeleteVfxtIaasNode(avereVfxt *AvereVfxt, nodeName string) error
}

type AvereVfxt struct {
	ControllerAddress string
	ControllerUsename string

	SshAuthMethod ssh.AuthMethod

	Platform IaasPlatform

	AvereVfxtName      string
	AvereAdminPassword string
	NodeCount          int

	NtpServers *[]string

	ProxyUri        string
	ClusterProxyUri string

	ImageId string

	ManagementIP       string
	VServerIPAddresses *[]string
	NodeNames          *[]string
}

///////////////////////////////////////////////////////////
// The following types are used to parse json from
// averecmd.
///////////////////////////////////////////////////////////

type NFSExport struct {
	Path string `json:"path"`
}

type Node struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type VServerClientIPHome struct {
	NodeName  string `json:"current"`
	IPAddress string `json:"ip"`
}

type Activity struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	State   string `json:"state"`
	Percent string `json:"percent"`
}

type Alert struct {
	Name     string `json:"name"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type CoreFilerGeneric struct {
	Name         string `json:"name"`
	NetworkName  string `json:"networkName"`
	PolicyName   string `json:"policyName"`
	InternalName string `json:"internalName"`
	FilerClass   string `json:"filerClass"`
	Bucket       string `json:"bucket"`
}

type CoreFiler struct {
	Name            string `json:"name"`
	FqdnOrPrimaryIp string `json:"networkName"`
	CachePolicy     string `json:"policyName"`
	CustomSettings  []*CustomSetting
}

// an Azure Storage Account Filer can be used from a vFXT running in
// any platform
type AzureStorageFiler struct {
	AccountName    string
	Container      string
	CustomSettings []*CustomSetting
}

type Junction struct {
	NameSpacePath    string `json:"path"`
	CoreFilerName    string `json:"mass"`
	CoreFilerExport  string `json:"export"`
	SharePermissions string
}

type CustomSetting struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CheckCode string `json:"checkCode"`
}
