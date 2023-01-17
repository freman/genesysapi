// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetContentmanagementSharedSharedIDReader is a Reader for the GetContentmanagementSharedSharedID structure.
type GetContentmanagementSharedSharedIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementSharedSharedIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementSharedSharedIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetContentmanagementSharedSharedIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 307:
		result := NewGetContentmanagementSharedSharedIDTemporaryRedirect()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetContentmanagementSharedSharedIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementSharedSharedIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementSharedSharedIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementSharedSharedIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementSharedSharedIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementSharedSharedIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementSharedSharedIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementSharedSharedIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementSharedSharedIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementSharedSharedIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementSharedSharedIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementSharedSharedIDOK creates a GetContentmanagementSharedSharedIDOK with default headers values
func NewGetContentmanagementSharedSharedIDOK() *GetContentmanagementSharedSharedIDOK {
	return &GetContentmanagementSharedSharedIDOK{}
}

/*
GetContentmanagementSharedSharedIDOK describes a response with status code 200, with default header values.

Download location is returned in header, if redirect is set to false and disposition is not set to none. If disposition is none, location header will not be populated, DownloadUri and ViewUri will be populated.
*/
type GetContentmanagementSharedSharedIDOK struct {
	Payload *models.SharedResponse
}

// IsSuccess returns true when this get contentmanagement shared shared Id o k response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contentmanagement shared shared Id o k response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id o k response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shared shared Id o k response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id o k response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContentmanagementSharedSharedIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDOK) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDOK) GetPayload() *models.SharedResponse {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SharedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDAccepted creates a GetContentmanagementSharedSharedIDAccepted with default headers values
func NewGetContentmanagementSharedSharedIDAccepted() *GetContentmanagementSharedSharedIDAccepted {
	return &GetContentmanagementSharedSharedIDAccepted{}
}

/*
GetContentmanagementSharedSharedIDAccepted describes a response with status code 202, with default header values.

Accepted - Preparing file for download - try again soon.
*/
type GetContentmanagementSharedSharedIDAccepted struct {
}

// IsSuccess returns true when this get contentmanagement shared shared Id accepted response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contentmanagement shared shared Id accepted response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id accepted response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shared shared Id accepted response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id accepted response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GetContentmanagementSharedSharedIDAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdAccepted ", 202)
}

func (o *GetContentmanagementSharedSharedIDAccepted) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdAccepted ", 202)
}

func (o *GetContentmanagementSharedSharedIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContentmanagementSharedSharedIDTemporaryRedirect creates a GetContentmanagementSharedSharedIDTemporaryRedirect with default headers values
func NewGetContentmanagementSharedSharedIDTemporaryRedirect() *GetContentmanagementSharedSharedIDTemporaryRedirect {
	return &GetContentmanagementSharedSharedIDTemporaryRedirect{}
}

/*
GetContentmanagementSharedSharedIDTemporaryRedirect describes a response with status code 307, with default header values.

Redirected to download location, if redirect is set to true
*/
type GetContentmanagementSharedSharedIDTemporaryRedirect struct {
}

// IsSuccess returns true when this get contentmanagement shared shared Id temporary redirect response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id temporary redirect response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get contentmanagement shared shared Id temporary redirect response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shared shared Id temporary redirect response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id temporary redirect response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) IsCode(code int) bool {
	return code == 307
}

func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdTemporaryRedirect ", 307)
}

func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdTemporaryRedirect ", 307)
}

func (o *GetContentmanagementSharedSharedIDTemporaryRedirect) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContentmanagementSharedSharedIDBadRequest creates a GetContentmanagementSharedSharedIDBadRequest with default headers values
func NewGetContentmanagementSharedSharedIDBadRequest() *GetContentmanagementSharedSharedIDBadRequest {
	return &GetContentmanagementSharedSharedIDBadRequest{}
}

/*
GetContentmanagementSharedSharedIDBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementSharedSharedIDBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id bad request response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id bad request response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id bad request response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id bad request response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id bad request response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetContentmanagementSharedSharedIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDUnauthorized creates a GetContentmanagementSharedSharedIDUnauthorized with default headers values
func NewGetContentmanagementSharedSharedIDUnauthorized() *GetContentmanagementSharedSharedIDUnauthorized {
	return &GetContentmanagementSharedSharedIDUnauthorized{}
}

/*
GetContentmanagementSharedSharedIDUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementSharedSharedIDUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id unauthorized response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id unauthorized response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id unauthorized response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id unauthorized response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id unauthorized response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetContentmanagementSharedSharedIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDForbidden creates a GetContentmanagementSharedSharedIDForbidden with default headers values
func NewGetContentmanagementSharedSharedIDForbidden() *GetContentmanagementSharedSharedIDForbidden {
	return &GetContentmanagementSharedSharedIDForbidden{}
}

/*
GetContentmanagementSharedSharedIDForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementSharedSharedIDForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id forbidden response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id forbidden response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id forbidden response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id forbidden response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id forbidden response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetContentmanagementSharedSharedIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDNotFound creates a GetContentmanagementSharedSharedIDNotFound with default headers values
func NewGetContentmanagementSharedSharedIDNotFound() *GetContentmanagementSharedSharedIDNotFound {
	return &GetContentmanagementSharedSharedIDNotFound{}
}

/*
GetContentmanagementSharedSharedIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetContentmanagementSharedSharedIDNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id not found response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id not found response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id not found response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id not found response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id not found response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetContentmanagementSharedSharedIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDRequestTimeout creates a GetContentmanagementSharedSharedIDRequestTimeout with default headers values
func NewGetContentmanagementSharedSharedIDRequestTimeout() *GetContentmanagementSharedSharedIDRequestTimeout {
	return &GetContentmanagementSharedSharedIDRequestTimeout{}
}

/*
GetContentmanagementSharedSharedIDRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementSharedSharedIDRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id request timeout response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id request timeout response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id request timeout response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id request timeout response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id request timeout response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetContentmanagementSharedSharedIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDRequestEntityTooLarge creates a GetContentmanagementSharedSharedIDRequestEntityTooLarge with default headers values
func NewGetContentmanagementSharedSharedIDRequestEntityTooLarge() *GetContentmanagementSharedSharedIDRequestEntityTooLarge {
	return &GetContentmanagementSharedSharedIDRequestEntityTooLarge{}
}

/*
GetContentmanagementSharedSharedIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetContentmanagementSharedSharedIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id request entity too large response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id request entity too large response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id request entity too large response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id request entity too large response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id request entity too large response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDUnsupportedMediaType creates a GetContentmanagementSharedSharedIDUnsupportedMediaType with default headers values
func NewGetContentmanagementSharedSharedIDUnsupportedMediaType() *GetContentmanagementSharedSharedIDUnsupportedMediaType {
	return &GetContentmanagementSharedSharedIDUnsupportedMediaType{}
}

/*
GetContentmanagementSharedSharedIDUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementSharedSharedIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id unsupported media type response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id unsupported media type response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id unsupported media type response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id unsupported media type response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id unsupported media type response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDTooManyRequests creates a GetContentmanagementSharedSharedIDTooManyRequests with default headers values
func NewGetContentmanagementSharedSharedIDTooManyRequests() *GetContentmanagementSharedSharedIDTooManyRequests {
	return &GetContentmanagementSharedSharedIDTooManyRequests{}
}

/*
GetContentmanagementSharedSharedIDTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementSharedSharedIDTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id too many requests response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id too many requests response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id too many requests response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement shared shared Id too many requests response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement shared shared Id too many requests response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetContentmanagementSharedSharedIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDInternalServerError creates a GetContentmanagementSharedSharedIDInternalServerError with default headers values
func NewGetContentmanagementSharedSharedIDInternalServerError() *GetContentmanagementSharedSharedIDInternalServerError {
	return &GetContentmanagementSharedSharedIDInternalServerError{}
}

/*
GetContentmanagementSharedSharedIDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementSharedSharedIDInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id internal server error response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id internal server error response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id internal server error response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shared shared Id internal server error response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement shared shared Id internal server error response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetContentmanagementSharedSharedIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDServiceUnavailable creates a GetContentmanagementSharedSharedIDServiceUnavailable with default headers values
func NewGetContentmanagementSharedSharedIDServiceUnavailable() *GetContentmanagementSharedSharedIDServiceUnavailable {
	return &GetContentmanagementSharedSharedIDServiceUnavailable{}
}

/*
GetContentmanagementSharedSharedIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementSharedSharedIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id service unavailable response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id service unavailable response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id service unavailable response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shared shared Id service unavailable response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement shared shared Id service unavailable response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetContentmanagementSharedSharedIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSharedSharedIDGatewayTimeout creates a GetContentmanagementSharedSharedIDGatewayTimeout with default headers values
func NewGetContentmanagementSharedSharedIDGatewayTimeout() *GetContentmanagementSharedSharedIDGatewayTimeout {
	return &GetContentmanagementSharedSharedIDGatewayTimeout{}
}

/*
GetContentmanagementSharedSharedIDGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetContentmanagementSharedSharedIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement shared shared Id gateway timeout response has a 2xx status code
func (o *GetContentmanagementSharedSharedIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement shared shared Id gateway timeout response has a 3xx status code
func (o *GetContentmanagementSharedSharedIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement shared shared Id gateway timeout response has a 4xx status code
func (o *GetContentmanagementSharedSharedIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement shared shared Id gateway timeout response has a 5xx status code
func (o *GetContentmanagementSharedSharedIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement shared shared Id gateway timeout response a status code equal to that given
func (o *GetContentmanagementSharedSharedIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetContentmanagementSharedSharedIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shared/{sharedId}][%d] getContentmanagementSharedSharedIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementSharedSharedIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSharedSharedIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
