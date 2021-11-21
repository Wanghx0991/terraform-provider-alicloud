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

// ModifyWebCCRule invokes the ddoscoo.ModifyWebCCRule API synchronously
func (client *Client) ModifyWebCCRule(request *ModifyWebCCRuleRequest) (response *ModifyWebCCRuleResponse, err error) {
	response = CreateModifyWebCCRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebCCRuleWithChan invokes the ddoscoo.ModifyWebCCRule API asynchronously
func (client *Client) ModifyWebCCRuleWithChan(request *ModifyWebCCRuleRequest) (<-chan *ModifyWebCCRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyWebCCRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebCCRule(request)
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

// ModifyWebCCRuleWithCallback invokes the ddoscoo.ModifyWebCCRule API asynchronously
func (client *Client) ModifyWebCCRuleWithCallback(request *ModifyWebCCRuleRequest, callback func(response *ModifyWebCCRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebCCRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebCCRule(request)
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

// ModifyWebCCRuleRequest is the request struct for api ModifyWebCCRule
type ModifyWebCCRuleRequest struct {
	*requests.RpcRequest
	Mode            string           `position:"Query" name:"Mode"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	Act             string           `position:"Query" name:"Act"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Count           requests.Integer `position:"Query" name:"Count"`
	Ttl             requests.Integer `position:"Query" name:"Ttl"`
	Uri             string           `position:"Query" name:"Uri"`
	Domain          string           `position:"Query" name:"Domain"`
	Name            string           `position:"Query" name:"Name"`
	Interval        requests.Integer `position:"Query" name:"Interval"`
}

// ModifyWebCCRuleResponse is the response struct for api ModifyWebCCRule
type ModifyWebCCRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebCCRuleRequest creates a request to invoke ModifyWebCCRule API
func CreateModifyWebCCRuleRequest() (request *ModifyWebCCRuleRequest) {
	request = &ModifyWebCCRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyWebCCRule", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyWebCCRuleResponse creates a response to parse from ModifyWebCCRule response
func CreateModifyWebCCRuleResponse() (response *ModifyWebCCRuleResponse) {
	response = &ModifyWebCCRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
