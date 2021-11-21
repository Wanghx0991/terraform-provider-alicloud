package emr

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

// UpdateTag invokes the emr.UpdateTag API synchronously
func (client *Client) UpdateTag(request *UpdateTagRequest) (response *UpdateTagResponse, err error) {
	response = CreateUpdateTagResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTagWithChan invokes the emr.UpdateTag API asynchronously
func (client *Client) UpdateTagWithChan(request *UpdateTagRequest) (<-chan *UpdateTagResponse, <-chan error) {
	responseChan := make(chan *UpdateTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTag(request)
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

// UpdateTagWithCallback invokes the emr.UpdateTag API asynchronously
func (client *Client) UpdateTagWithCallback(request *UpdateTagRequest, callback func(response *UpdateTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTagResponse
		var err error
		defer close(result)
		response, err = client.UpdateTag(request)
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

// UpdateTagRequest is the request struct for api UpdateTag
type UpdateTagRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description     string           `position:"Query" name:"Description"`
	Name            string           `position:"Query" name:"Name"`
	Id              requests.Integer `position:"Query" name:"Id"`
	Category        string           `position:"Query" name:"Category"`
}

// UpdateTagResponse is the response struct for api UpdateTag
type UpdateTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateUpdateTagRequest creates a request to invoke UpdateTag API
func CreateUpdateTagRequest() (request *UpdateTagRequest) {
	request = &UpdateTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "UpdateTag", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTagResponse creates a response to parse from UpdateTag response
func CreateUpdateTagResponse() (response *UpdateTagResponse) {
	response = &UpdateTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
