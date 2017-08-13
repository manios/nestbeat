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

type ModelError struct {

	// Short error message format.
	Error_ int32 `json:"error,omitempty"`

	// Provides a URL to detailed information about the error condition (this page).
	Type_ string `json:"type,omitempty"`

	// Long error message format that may use variables to provide additional details. When a variable is included in the message, it will appear in the **details** object.
	Message string `json:"message,omitempty"`

	// A text string that holds an error identifier that is unique to each individual call. We may ask you for the instance number if you report an issue with the service.
	Instance string `json:"instance,omitempty"`

	// Optional. Contains variables that are inserted into the message. Messages can contain multiple variables.
	Details string `json:"details,omitempty"`
}
