package cassandra

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeBackupPlan invokes the cassandra.DescribeBackupPlan API synchronously
func (client *Client) DescribeBackupPlan(request *DescribeBackupPlanRequest) (response *DescribeBackupPlanResponse, err error) {
	response = CreateDescribeBackupPlanResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupPlanWithChan invokes the cassandra.DescribeBackupPlan API asynchronously
func (client *Client) DescribeBackupPlanWithChan(request *DescribeBackupPlanRequest) (<-chan *DescribeBackupPlanResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupPlan(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeBackupPlanWithCallback invokes the cassandra.DescribeBackupPlan API asynchronously
func (client *Client) DescribeBackupPlanWithCallback(request *DescribeBackupPlanRequest, callback func(response *DescribeBackupPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupPlanResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupPlan(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeBackupPlanRequest is the request struct for api DescribeBackupPlan
type DescribeBackupPlanRequest struct {
	*requests.RpcRequest
	DataCenterId string `position:"Query" name:"DataCenterId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
}

// DescribeBackupPlanResponse is the response struct for api DescribeBackupPlan
type DescribeBackupPlanResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	BackupPlan BackupPlan `json:"BackupPlan" xml:"BackupPlan"`
}

// CreateDescribeBackupPlanRequest creates a request to invoke DescribeBackupPlan API
func CreateDescribeBackupPlanRequest() (request *DescribeBackupPlanRequest) {
	request = &DescribeBackupPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeBackupPlan", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupPlanResponse creates a response to parse from DescribeBackupPlan response
func CreateDescribeBackupPlanResponse() (response *DescribeBackupPlanResponse) {
	response = &DescribeBackupPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
