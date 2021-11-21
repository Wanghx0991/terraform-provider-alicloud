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

// QueryInstanceByTag invokes the bssopenapi.QueryInstanceByTag API synchronously
func (client *Client) QueryInstanceByTag(request *QueryInstanceByTagRequest) (response *QueryInstanceByTagResponse, err error) {
	response = CreateQueryInstanceByTagResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInstanceByTagWithChan invokes the bssopenapi.QueryInstanceByTag API asynchronously
func (client *Client) QueryInstanceByTagWithChan(request *QueryInstanceByTagRequest) (<-chan *QueryInstanceByTagResponse, <-chan error) {
	responseChan := make(chan *QueryInstanceByTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInstanceByTag(request)
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

// QueryInstanceByTagWithCallback invokes the bssopenapi.QueryInstanceByTag API asynchronously
func (client *Client) QueryInstanceByTagWithCallback(request *QueryInstanceByTagRequest, callback func(response *QueryInstanceByTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInstanceByTagResponse
		var err error
		defer close(result)
		response, err = client.QueryInstanceByTag(request)
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

// QueryInstanceByTagRequest is the request struct for api QueryInstanceByTag
type QueryInstanceByTagRequest struct {
	*requests.RpcRequest
	ResourceId   *[]string                `position:"Query" name:"ResourceId"  type:"Repeated"`
	Tag          *[]QueryInstanceByTagTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceType string                   `position:"Query" name:"ResourceType"`
}

// QueryInstanceByTagTag is a repeated param struct in QueryInstanceByTagRequest
type QueryInstanceByTagTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// QueryInstanceByTagResponse is the response struct for api QueryInstanceByTag
type QueryInstanceByTagResponse struct {
	*responses.BaseResponse
	Code        string            `json:"Code" xml:"Code"`
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	Success     bool              `json:"Success" xml:"Success"`
	Message     string            `json:"Message" xml:"Message"`
	NextToken   string            `json:"NextToken" xml:"NextToken"`
	TagResource []TagResourceItem `json:"TagResource" xml:"TagResource"`
}

// CreateQueryInstanceByTagRequest creates a request to invoke QueryInstanceByTag API
func CreateQueryInstanceByTagRequest() (request *QueryInstanceByTagRequest) {
	request = &QueryInstanceByTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryInstanceByTag", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryInstanceByTagResponse creates a response to parse from QueryInstanceByTag response
func CreateQueryInstanceByTagResponse() (response *QueryInstanceByTagResponse) {
	response = &QueryInstanceByTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
