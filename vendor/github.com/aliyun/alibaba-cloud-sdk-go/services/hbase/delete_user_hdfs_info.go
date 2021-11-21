package hbase

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

// DeleteUserHdfsInfo invokes the hbase.DeleteUserHdfsInfo API synchronously
func (client *Client) DeleteUserHdfsInfo(request *DeleteUserHdfsInfoRequest) (response *DeleteUserHdfsInfoResponse, err error) {
	response = CreateDeleteUserHdfsInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserHdfsInfoWithChan invokes the hbase.DeleteUserHdfsInfo API asynchronously
func (client *Client) DeleteUserHdfsInfoWithChan(request *DeleteUserHdfsInfoRequest) (<-chan *DeleteUserHdfsInfoResponse, <-chan error) {
	responseChan := make(chan *DeleteUserHdfsInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUserHdfsInfo(request)
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

// DeleteUserHdfsInfoWithCallback invokes the hbase.DeleteUserHdfsInfo API asynchronously
func (client *Client) DeleteUserHdfsInfoWithCallback(request *DeleteUserHdfsInfoRequest, callback func(response *DeleteUserHdfsInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserHdfsInfoResponse
		var err error
		defer close(result)
		response, err = client.DeleteUserHdfsInfo(request)
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

// DeleteUserHdfsInfoRequest is the request struct for api DeleteUserHdfsInfo
type DeleteUserHdfsInfoRequest struct {
	*requests.RpcRequest
	ClusterId   string `position:"Query" name:"ClusterId"`
	NameService string `position:"Query" name:"NameService"`
}

// DeleteUserHdfsInfoResponse is the response struct for api DeleteUserHdfsInfo
type DeleteUserHdfsInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUserHdfsInfoRequest creates a request to invoke DeleteUserHdfsInfo API
func CreateDeleteUserHdfsInfoRequest() (request *DeleteUserHdfsInfoRequest) {
	request = &DeleteUserHdfsInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DeleteUserHdfsInfo", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteUserHdfsInfoResponse creates a response to parse from DeleteUserHdfsInfo response
func CreateDeleteUserHdfsInfoResponse() (response *DeleteUserHdfsInfoResponse) {
	response = &DeleteUserHdfsInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
