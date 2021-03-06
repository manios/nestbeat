/* 
 * Nest API
 *
 * The Nest API models a physical home or building as a structure, with Nest Learning Thermostats, Nest Protects, and Nest Cams as devices in the structure. This structure also contains information about the home as a whole (such as Away or ETA state, or active Rush Hours).  > **Key Point**: All devices share a common base set of information: a > user-supplied name, physical location in the home, software version, and > online status.  Every data element in the structure is addressable by a resource URL (called \"data locations\") in a shared JSON document. Each data location can store strings, numbers, booleans, parent/child objects, or arrays.  > Note: Use this root URL when making Nest API calls: ```https://developer-api.nest.com```   From the API, you can sync data from locations at multiple levels in the data model, for example:  * an entire structure, including all devices  * a single device in a structure  * a group of data values (current and ambient temperature)  * a single data value (battery health state) 
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package gonest

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"net/url"
	"io/ioutil"
	"gopkg.in/go-resty/resty.v0"
)

type APIClient struct {
	config *Configuration
}

func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {

	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *APIClient) SelectHeaderAccept(accepts []string) string {

	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) (*resty.Response, error) {

	rClient := c.prepareClient()
	request := c.prepareRequest(rClient, postBody, headerParams, queryParams, formParams, fileName, fileBytes)

	switch strings.ToUpper(method) {
	case "GET":
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err
	}

	return nil, fmt.Errorf("invalid method %v", method)
}

func (c *APIClient) ParameterToString(obj interface{}, collectionFormat string) string {
	delimiter := ""
	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

func (c *APIClient) prepareClient() *resty.Client {

	rClient := resty.New()

	rClient.SetDebug(c.config.Debug)
	if c.config.Transport != nil {
		rClient.SetTransport(c.config.Transport)
	}

	if c.config.Timeout != nil {
		rClient.SetTimeout(*c.config.Timeout)
	}
	rClient.SetLogger(ioutil.Discard)
	return rClient
}

func (c *APIClient) prepareRequest(
	rClient *resty.Client,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) *resty.Request {


	request := rClient.R()
	request.SetBody(postBody)

	if c.config.UserAgent != "" {
		request.SetHeader("User-Agent", c.config.UserAgent)
	}

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetMultiValueQueryParams(queryParams)
	}

	// add form parameter, if any
	if len(formParams) > 0 {
		request.SetFormData(formParams)
	}

	if len(fileBytes) > 0 && fileName != "" {
		_, fileNm := filepath.Split(fileName)
		request.SetFileReader("file", fileNm, bytes.NewReader(fileBytes))
	}
	return request
}
