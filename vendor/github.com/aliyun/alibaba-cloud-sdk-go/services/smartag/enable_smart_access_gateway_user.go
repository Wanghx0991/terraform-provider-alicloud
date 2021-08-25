package smartag

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

// EnableSmartAccessGatewayUser invokes the smartag.EnableSmartAccessGatewayUser API synchronously
func (client *Client) EnableSmartAccessGatewayUser(request *EnableSmartAccessGatewayUserRequest) (response *EnableSmartAccessGatewayUserResponse, err error) {
	response = CreateEnableSmartAccessGatewayUserResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSmartAccessGatewayUserWithChan invokes the smartag.EnableSmartAccessGatewayUser API asynchronously
func (client *Client) EnableSmartAccessGatewayUserWithChan(request *EnableSmartAccessGatewayUserRequest) (<-chan *EnableSmartAccessGatewayUserResponse, <-chan error) {
	responseChan := make(chan *EnableSmartAccessGatewayUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSmartAccessGatewayUser(request)
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

// EnableSmartAccessGatewayUserWithCallback invokes the smartag.EnableSmartAccessGatewayUser API asynchronously
func (client *Client) EnableSmartAccessGatewayUserWithCallback(request *EnableSmartAccessGatewayUserRequest, callback func(response *EnableSmartAccessGatewayUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSmartAccessGatewayUserResponse
		var err error
		defer close(result)
		response, err = client.EnableSmartAccessGatewayUser(request)
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

// EnableSmartAccessGatewayUserRequest is the request struct for api EnableSmartAccessGatewayUser
type EnableSmartAccessGatewayUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// EnableSmartAccessGatewayUserResponse is the response struct for api EnableSmartAccessGatewayUser
type EnableSmartAccessGatewayUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableSmartAccessGatewayUserRequest creates a request to invoke EnableSmartAccessGatewayUser API
func CreateEnableSmartAccessGatewayUserRequest() (request *EnableSmartAccessGatewayUserRequest) {
	request = &EnableSmartAccessGatewayUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "EnableSmartAccessGatewayUser", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableSmartAccessGatewayUserResponse creates a response to parse from EnableSmartAccessGatewayUser response
func CreateEnableSmartAccessGatewayUserResponse() (response *EnableSmartAccessGatewayUserResponse) {
	response = &EnableSmartAccessGatewayUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
