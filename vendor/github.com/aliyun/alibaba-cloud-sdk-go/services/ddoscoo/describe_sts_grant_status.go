package ddoscoo

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

// DescribeStsGrantStatus invokes the ddoscoo.DescribeStsGrantStatus API synchronously
func (client *Client) DescribeStsGrantStatus(request *DescribeStsGrantStatusRequest) (response *DescribeStsGrantStatusResponse, err error) {
	response = CreateDescribeStsGrantStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStsGrantStatusWithChan invokes the ddoscoo.DescribeStsGrantStatus API asynchronously
func (client *Client) DescribeStsGrantStatusWithChan(request *DescribeStsGrantStatusRequest) (<-chan *DescribeStsGrantStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeStsGrantStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStsGrantStatus(request)
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

// DescribeStsGrantStatusWithCallback invokes the ddoscoo.DescribeStsGrantStatus API asynchronously
func (client *Client) DescribeStsGrantStatusWithCallback(request *DescribeStsGrantStatusRequest, callback func(response *DescribeStsGrantStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStsGrantStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeStsGrantStatus(request)
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

// DescribeStsGrantStatusRequest is the request struct for api DescribeStsGrantStatus
type DescribeStsGrantStatusRequest struct {
	*requests.RpcRequest
	Role            string `position:"Query" name:"Role"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

// DescribeStsGrantStatusResponse is the response struct for api DescribeStsGrantStatus
type DescribeStsGrantStatusResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	StsGrant  StsGrant `json:"StsGrant" xml:"StsGrant"`
}

// CreateDescribeStsGrantStatusRequest creates a request to invoke DescribeStsGrantStatus API
func CreateDescribeStsGrantStatusRequest() (request *DescribeStsGrantStatusRequest) {
	request = &DescribeStsGrantStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeStsGrantStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStsGrantStatusResponse creates a response to parse from DescribeStsGrantStatus response
func CreateDescribeStsGrantStatusResponse() (response *DescribeStsGrantStatusResponse) {
	response = &DescribeStsGrantStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
