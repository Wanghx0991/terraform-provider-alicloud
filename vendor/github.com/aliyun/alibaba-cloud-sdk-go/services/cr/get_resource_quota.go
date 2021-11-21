package cr

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

// GetResourceQuota invokes the cr.GetResourceQuota API synchronously
func (client *Client) GetResourceQuota(request *GetResourceQuotaRequest) (response *GetResourceQuotaResponse, err error) {
	response = CreateGetResourceQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceQuotaWithChan invokes the cr.GetResourceQuota API asynchronously
func (client *Client) GetResourceQuotaWithChan(request *GetResourceQuotaRequest) (<-chan *GetResourceQuotaResponse, <-chan error) {
	responseChan := make(chan *GetResourceQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceQuota(request)
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

// GetResourceQuotaWithCallback invokes the cr.GetResourceQuota API asynchronously
func (client *Client) GetResourceQuotaWithCallback(request *GetResourceQuotaRequest, callback func(response *GetResourceQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceQuotaResponse
		var err error
		defer close(result)
		response, err = client.GetResourceQuota(request)
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

// GetResourceQuotaRequest is the request struct for api GetResourceQuota
type GetResourceQuotaRequest struct {
	*requests.RoaRequest
	ResourceName string `position:"Path" name:"ResourceName"`
}

// GetResourceQuotaResponse is the response struct for api GetResourceQuota
type GetResourceQuotaResponse struct {
	*responses.BaseResponse
}

// CreateGetResourceQuotaRequest creates a request to invoke GetResourceQuota API
func CreateGetResourceQuotaRequest() (request *GetResourceQuotaRequest) {
	request = &GetResourceQuotaRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetResourceQuota", "/resource/[ResourceName]", "acr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetResourceQuotaResponse creates a response to parse from GetResourceQuota response
func CreateGetResourceQuotaResponse() (response *GetResourceQuotaResponse) {
	response = &GetResourceQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
