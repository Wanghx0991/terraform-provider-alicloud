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

// ReleasePublicContactPoints invokes the cassandra.ReleasePublicContactPoints API synchronously
func (client *Client) ReleasePublicContactPoints(request *ReleasePublicContactPointsRequest) (response *ReleasePublicContactPointsResponse, err error) {
	response = CreateReleasePublicContactPointsResponse()
	err = client.DoAction(request, response)
	return
}

// ReleasePublicContactPointsWithChan invokes the cassandra.ReleasePublicContactPoints API asynchronously
func (client *Client) ReleasePublicContactPointsWithChan(request *ReleasePublicContactPointsRequest) (<-chan *ReleasePublicContactPointsResponse, <-chan error) {
	responseChan := make(chan *ReleasePublicContactPointsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleasePublicContactPoints(request)
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

// ReleasePublicContactPointsWithCallback invokes the cassandra.ReleasePublicContactPoints API asynchronously
func (client *Client) ReleasePublicContactPointsWithCallback(request *ReleasePublicContactPointsRequest, callback func(response *ReleasePublicContactPointsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleasePublicContactPointsResponse
		var err error
		defer close(result)
		response, err = client.ReleasePublicContactPoints(request)
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

// ReleasePublicContactPointsRequest is the request struct for api ReleasePublicContactPoints
type ReleasePublicContactPointsRequest struct {
	*requests.RpcRequest
	DataCenterId string `position:"Query" name:"DataCenterId"`
	ClusterId    string `position:"Query" name:"ClusterId"`
}

// ReleasePublicContactPointsResponse is the response struct for api ReleasePublicContactPoints
type ReleasePublicContactPointsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleasePublicContactPointsRequest creates a request to invoke ReleasePublicContactPoints API
func CreateReleasePublicContactPointsRequest() (request *ReleasePublicContactPointsRequest) {
	request = &ReleasePublicContactPointsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "ReleasePublicContactPoints", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleasePublicContactPointsResponse creates a response to parse from ReleasePublicContactPoints response
func CreateReleasePublicContactPointsResponse() (response *ReleasePublicContactPointsResponse) {
	response = &ReleasePublicContactPointsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
