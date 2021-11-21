package bssopenapi

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

// EnableBillGeneration invokes the bssopenapi.EnableBillGeneration API synchronously
func (client *Client) EnableBillGeneration(request *EnableBillGenerationRequest) (response *EnableBillGenerationResponse, err error) {
	response = CreateEnableBillGenerationResponse()
	err = client.DoAction(request, response)
	return
}

// EnableBillGenerationWithChan invokes the bssopenapi.EnableBillGeneration API asynchronously
func (client *Client) EnableBillGenerationWithChan(request *EnableBillGenerationRequest) (<-chan *EnableBillGenerationResponse, <-chan error) {
	responseChan := make(chan *EnableBillGenerationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableBillGeneration(request)
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

// EnableBillGenerationWithCallback invokes the bssopenapi.EnableBillGeneration API asynchronously
func (client *Client) EnableBillGenerationWithCallback(request *EnableBillGenerationRequest, callback func(response *EnableBillGenerationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableBillGenerationResponse
		var err error
		defer close(result)
		response, err = client.EnableBillGeneration(request)
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

// EnableBillGenerationRequest is the request struct for api EnableBillGeneration
type EnableBillGenerationRequest struct {
	*requests.RpcRequest
	ProductCode string           `position:"Query" name:"ProductCode"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// EnableBillGenerationResponse is the response struct for api EnableBillGeneration
type EnableBillGenerationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnableBillGenerationRequest creates a request to invoke EnableBillGeneration API
func CreateEnableBillGenerationRequest() (request *EnableBillGenerationRequest) {
	request = &EnableBillGenerationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "EnableBillGeneration", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableBillGenerationResponse creates a response to parse from EnableBillGeneration response
func CreateEnableBillGenerationResponse() (response *EnableBillGenerationResponse) {
	response = &EnableBillGenerationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
