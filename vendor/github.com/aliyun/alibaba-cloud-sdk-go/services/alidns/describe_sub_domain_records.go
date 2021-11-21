package alidns

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

// DescribeSubDomainRecords invokes the alidns.DescribeSubDomainRecords API synchronously
func (client *Client) DescribeSubDomainRecords(request *DescribeSubDomainRecordsRequest) (response *DescribeSubDomainRecordsResponse, err error) {
	response = CreateDescribeSubDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSubDomainRecordsWithChan invokes the alidns.DescribeSubDomainRecords API asynchronously
func (client *Client) DescribeSubDomainRecordsWithChan(request *DescribeSubDomainRecordsRequest) (<-chan *DescribeSubDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeSubDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSubDomainRecords(request)
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

// DescribeSubDomainRecordsWithCallback invokes the alidns.DescribeSubDomainRecords API asynchronously
func (client *Client) DescribeSubDomainRecordsWithCallback(request *DescribeSubDomainRecordsRequest, callback func(response *DescribeSubDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSubDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSubDomainRecords(request)
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

// DescribeSubDomainRecordsRequest is the request struct for api DescribeSubDomainRecords
type DescribeSubDomainRecordsRequest struct {
	*requests.RpcRequest
	Line         string           `position:"Query" name:"Line"`
	DomainName   string           `position:"Query" name:"DomainName"`
	Type         string           `position:"Query" name:"Type"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	SubDomain    string           `position:"Query" name:"SubDomain"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DescribeSubDomainRecordsResponse is the response struct for api DescribeSubDomainRecords
type DescribeSubDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId     string                                  `json:"RequestId" xml:"RequestId"`
	TotalCount    int64                                   `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int64                                   `json:"PageNumber" xml:"PageNumber"`
	PageSize      int64                                   `json:"PageSize" xml:"PageSize"`
	DomainRecords DomainRecordsInDescribeSubDomainRecords `json:"DomainRecords" xml:"DomainRecords"`
}

// CreateDescribeSubDomainRecordsRequest creates a request to invoke DescribeSubDomainRecords API
func CreateDescribeSubDomainRecordsRequest() (request *DescribeSubDomainRecordsRequest) {
	request = &DescribeSubDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeSubDomainRecords", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSubDomainRecordsResponse creates a response to parse from DescribeSubDomainRecords response
func CreateDescribeSubDomainRecordsResponse() (response *DescribeSubDomainRecordsResponse) {
	response = &DescribeSubDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
