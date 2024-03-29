// Code generated by go-swagger; DO NOT EDIT.

package scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScriptsPublishedScriptIDPagesReader is a Reader for the GetScriptsPublishedScriptIDPages structure.
type GetScriptsPublishedScriptIDPagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptsPublishedScriptIDPagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptsPublishedScriptIDPagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScriptsPublishedScriptIDPagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScriptsPublishedScriptIDPagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptsPublishedScriptIDPagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScriptsPublishedScriptIDPagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScriptsPublishedScriptIDPagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScriptsPublishedScriptIDPagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScriptsPublishedScriptIDPagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScriptsPublishedScriptIDPagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScriptsPublishedScriptIDPagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScriptsPublishedScriptIDPagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScriptsPublishedScriptIDPagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScriptsPublishedScriptIDPagesOK creates a GetScriptsPublishedScriptIDPagesOK with default headers values
func NewGetScriptsPublishedScriptIDPagesOK() *GetScriptsPublishedScriptIDPagesOK {
	return &GetScriptsPublishedScriptIDPagesOK{}
}

/*
GetScriptsPublishedScriptIDPagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScriptsPublishedScriptIDPagesOK struct {
	Payload []*models.Page
}

// IsSuccess returns true when this get scripts published script Id pages o k response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scripts published script Id pages o k response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages o k response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id pages o k response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages o k response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScriptsPublishedScriptIDPagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesOK) GetPayload() []*models.Page {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesBadRequest creates a GetScriptsPublishedScriptIDPagesBadRequest with default headers values
func NewGetScriptsPublishedScriptIDPagesBadRequest() *GetScriptsPublishedScriptIDPagesBadRequest {
	return &GetScriptsPublishedScriptIDPagesBadRequest{}
}

/*
GetScriptsPublishedScriptIDPagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsPublishedScriptIDPagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages bad request response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages bad request response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages bad request response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages bad request response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages bad request response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesUnauthorized creates a GetScriptsPublishedScriptIDPagesUnauthorized with default headers values
func NewGetScriptsPublishedScriptIDPagesUnauthorized() *GetScriptsPublishedScriptIDPagesUnauthorized {
	return &GetScriptsPublishedScriptIDPagesUnauthorized{}
}

/*
GetScriptsPublishedScriptIDPagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsPublishedScriptIDPagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages unauthorized response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages unauthorized response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages unauthorized response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages unauthorized response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages unauthorized response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesForbidden creates a GetScriptsPublishedScriptIDPagesForbidden with default headers values
func NewGetScriptsPublishedScriptIDPagesForbidden() *GetScriptsPublishedScriptIDPagesForbidden {
	return &GetScriptsPublishedScriptIDPagesForbidden{}
}

/*
GetScriptsPublishedScriptIDPagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsPublishedScriptIDPagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages forbidden response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages forbidden response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages forbidden response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages forbidden response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages forbidden response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesNotFound creates a GetScriptsPublishedScriptIDPagesNotFound with default headers values
func NewGetScriptsPublishedScriptIDPagesNotFound() *GetScriptsPublishedScriptIDPagesNotFound {
	return &GetScriptsPublishedScriptIDPagesNotFound{}
}

/*
GetScriptsPublishedScriptIDPagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetScriptsPublishedScriptIDPagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages not found response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages not found response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages not found response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages not found response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages not found response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesRequestTimeout creates a GetScriptsPublishedScriptIDPagesRequestTimeout with default headers values
func NewGetScriptsPublishedScriptIDPagesRequestTimeout() *GetScriptsPublishedScriptIDPagesRequestTimeout {
	return &GetScriptsPublishedScriptIDPagesRequestTimeout{}
}

/*
GetScriptsPublishedScriptIDPagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScriptsPublishedScriptIDPagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages request timeout response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages request timeout response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages request timeout response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages request timeout response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages request timeout response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesRequestEntityTooLarge creates a GetScriptsPublishedScriptIDPagesRequestEntityTooLarge with default headers values
func NewGetScriptsPublishedScriptIDPagesRequestEntityTooLarge() *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge {
	return &GetScriptsPublishedScriptIDPagesRequestEntityTooLarge{}
}

/*
GetScriptsPublishedScriptIDPagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetScriptsPublishedScriptIDPagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages request entity too large response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages request entity too large response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages request entity too large response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages request entity too large response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages request entity too large response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesUnsupportedMediaType creates a GetScriptsPublishedScriptIDPagesUnsupportedMediaType with default headers values
func NewGetScriptsPublishedScriptIDPagesUnsupportedMediaType() *GetScriptsPublishedScriptIDPagesUnsupportedMediaType {
	return &GetScriptsPublishedScriptIDPagesUnsupportedMediaType{}
}

/*
GetScriptsPublishedScriptIDPagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsPublishedScriptIDPagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages unsupported media type response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages unsupported media type response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages unsupported media type response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages unsupported media type response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages unsupported media type response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesTooManyRequests creates a GetScriptsPublishedScriptIDPagesTooManyRequests with default headers values
func NewGetScriptsPublishedScriptIDPagesTooManyRequests() *GetScriptsPublishedScriptIDPagesTooManyRequests {
	return &GetScriptsPublishedScriptIDPagesTooManyRequests{}
}

/*
GetScriptsPublishedScriptIDPagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScriptsPublishedScriptIDPagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages too many requests response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages too many requests response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages too many requests response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id pages too many requests response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id pages too many requests response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesInternalServerError creates a GetScriptsPublishedScriptIDPagesInternalServerError with default headers values
func NewGetScriptsPublishedScriptIDPagesInternalServerError() *GetScriptsPublishedScriptIDPagesInternalServerError {
	return &GetScriptsPublishedScriptIDPagesInternalServerError{}
}

/*
GetScriptsPublishedScriptIDPagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsPublishedScriptIDPagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages internal server error response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages internal server error response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages internal server error response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id pages internal server error response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published script Id pages internal server error response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesServiceUnavailable creates a GetScriptsPublishedScriptIDPagesServiceUnavailable with default headers values
func NewGetScriptsPublishedScriptIDPagesServiceUnavailable() *GetScriptsPublishedScriptIDPagesServiceUnavailable {
	return &GetScriptsPublishedScriptIDPagesServiceUnavailable{}
}

/*
GetScriptsPublishedScriptIDPagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsPublishedScriptIDPagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages service unavailable response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages service unavailable response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages service unavailable response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id pages service unavailable response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published script Id pages service unavailable response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesGatewayTimeout creates a GetScriptsPublishedScriptIDPagesGatewayTimeout with default headers values
func NewGetScriptsPublishedScriptIDPagesGatewayTimeout() *GetScriptsPublishedScriptIDPagesGatewayTimeout {
	return &GetScriptsPublishedScriptIDPagesGatewayTimeout{}
}

/*
GetScriptsPublishedScriptIDPagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetScriptsPublishedScriptIDPagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id pages gateway timeout response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id pages gateway timeout response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id pages gateway timeout response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id pages gateway timeout response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published script Id pages gateway timeout response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
