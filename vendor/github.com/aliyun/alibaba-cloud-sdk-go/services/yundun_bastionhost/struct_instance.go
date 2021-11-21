package yundun_bastionhost

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Instance is a nested struct in yundun_bastionhost response
type Instance struct {
	EcsIntranetIp       string `json:"EcsIntranetIp" xml:"EcsIntranetIp"`
	EcsInternetIp       string `json:"EcsInternetIp" xml:"EcsInternetIp"`
	ImageVersionName    string `json:"ImageVersionName" xml:"ImageVersionName"`
	StartTime           int64  `json:"StartTime" xml:"StartTime"`
	RegionId            string `json:"RegionId" xml:"RegionId"`
	InternetEndpoint    string `json:"InternetEndpoint" xml:"InternetEndpoint"`
	RegionNo            string `json:"RegionNo" xml:"RegionNo"`
	InstanceId          string `json:"InstanceId" xml:"InstanceId"`
	VswitchId           string `json:"VswitchId" xml:"VswitchId"`
	InstanceStatus      string `json:"InstanceStatus" xml:"InstanceStatus"`
	VpcId               string `json:"VpcId" xml:"VpcId"`
	NetworkType         string `json:"NetworkType" xml:"NetworkType"`
	IntranetIp          string `json:"IntranetIp" xml:"IntranetIp"`
	Description         string `json:"Description" xml:"Description"`
	SeriesCode          string `json:"SeriesCode" xml:"SeriesCode"`
	LicenseCode         string `json:"LicenseCode" xml:"LicenseCode"`
	Legacy              bool   `json:"Legacy" xml:"Legacy"`
	Status              int    `json:"Status" xml:"Status"`
	PublicNetworkAccess bool   `json:"PublicNetworkAccess" xml:"PublicNetworkAccess"`
	InternetIp          string `json:"InternetIp" xml:"InternetIp"`
	EcsNetworkType      string `json:"EcsNetworkType" xml:"EcsNetworkType"`
	Renewable           bool   `json:"Renewable" xml:"Renewable"`
	Upgradeable         bool   `json:"Upgradeable" xml:"Upgradeable"`
	CustomName          string `json:"CustomName" xml:"CustomName"`
	Operatable          bool   `json:"Operatable" xml:"Operatable"`
	EcsInstanceId       string `json:"EcsInstanceId" xml:"EcsInstanceId"`
	PlanUpgradeable     bool   `json:"PlanUpgradeable" xml:"PlanUpgradeable"`
	IntranetEndpoint    string `json:"IntranetEndpoint" xml:"IntranetEndpoint"`
	EcsStatus           string `json:"EcsStatus" xml:"EcsStatus"`
	PlanUpgradeStatus   int    `json:"PlanUpgradeStatus" xml:"PlanUpgradeStatus"`
	ExpireTime          int64  `json:"ExpireTime" xml:"ExpireTime"`
	UpgradeStatus       int    `json:"UpgradeStatus" xml:"UpgradeStatus"`
}
