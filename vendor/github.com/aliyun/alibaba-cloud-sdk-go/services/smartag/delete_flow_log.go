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

// DeleteFlowLog invokes the smartag.DeleteFlowLog API synchronously
func (client *Client) DeleteFlowLog(request *DeleteFlowLogRequest) (response *DeleteFlowLogResponse, err error) {
	response = CreateDeleteFlowLogResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFlowLogWithChan invokes the smartag.DeleteFlowLog API asynchronously
func (client *Client) DeleteFlowLogWithChan(request *DeleteFlowLogRequest) (<-chan *DeleteFlowLogResponse, <-chan error) {
	responseChan := make(chan *DeleteFlowLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFlowLog(request)
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

// DeleteFlowLogWithCallback invokes the smartag.DeleteFlowLog API asynchronously
func (client *Client) DeleteFlowLogWithCallback(request *DeleteFlowLogRequest, callback func(response *DeleteFlowLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFlowLogResponse
		var err error
		defer close(result)
		response, err = client.DeleteFlowLog(request)
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

// DeleteFlowLogRequest is the request struct for api DeleteFlowLog
type DeleteFlowLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
}

// DeleteFlowLogResponse is the response struct for api DeleteFlowLog
type DeleteFlowLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFlowLogRequest creates a request to invoke DeleteFlowLog API
func CreateDeleteFlowLogRequest() (request *DeleteFlowLogRequest) {
	request = &DeleteFlowLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteFlowLog", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFlowLogResponse creates a response to parse from DeleteFlowLog response
func CreateDeleteFlowLogResponse() (response *DeleteFlowLogResponse) {
	response = &DeleteFlowLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
