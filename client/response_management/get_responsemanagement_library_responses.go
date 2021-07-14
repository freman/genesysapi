// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetResponsemanagementLibraryReader is a Reader for the GetResponsemanagementLibrary structure.
type GetResponsemanagementLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponsemanagementLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponsemanagementLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResponsemanagementLibraryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResponsemanagementLibraryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResponsemanagementLibraryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResponsemanagementLibraryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetResponsemanagementLibraryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetResponsemanagementLibraryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetResponsemanagementLibraryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetResponsemanagementLibraryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResponsemanagementLibraryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetResponsemanagementLibraryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResponsemanagementLibraryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResponsemanagementLibraryOK creates a GetResponsemanagementLibraryOK with default headers values
func NewGetResponsemanagementLibraryOK() *GetResponsemanagementLibraryOK {
	return &GetResponsemanagementLibraryOK{}
}

/*GetResponsemanagementLibraryOK handles this case with default header values.

successful operation
*/
type GetResponsemanagementLibraryOK struct {
	Payload *models.Library
}

func (o *GetResponsemanagementLibraryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementLibraryOK) GetPayload() *models.Library {
	return o.Payload
}

func (o *GetResponsemanagementLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Library)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryBadRequest creates a GetResponsemanagementLibraryBadRequest with default headers values
func NewGetResponsemanagementLibraryBadRequest() *GetResponsemanagementLibraryBadRequest {
	return &GetResponsemanagementLibraryBadRequest{}
}

/*GetResponsemanagementLibraryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetResponsemanagementLibraryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementLibraryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryUnauthorized creates a GetResponsemanagementLibraryUnauthorized with default headers values
func NewGetResponsemanagementLibraryUnauthorized() *GetResponsemanagementLibraryUnauthorized {
	return &GetResponsemanagementLibraryUnauthorized{}
}

/*GetResponsemanagementLibraryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetResponsemanagementLibraryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementLibraryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryForbidden creates a GetResponsemanagementLibraryForbidden with default headers values
func NewGetResponsemanagementLibraryForbidden() *GetResponsemanagementLibraryForbidden {
	return &GetResponsemanagementLibraryForbidden{}
}

/*GetResponsemanagementLibraryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetResponsemanagementLibraryForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementLibraryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryNotFound creates a GetResponsemanagementLibraryNotFound with default headers values
func NewGetResponsemanagementLibraryNotFound() *GetResponsemanagementLibraryNotFound {
	return &GetResponsemanagementLibraryNotFound{}
}

/*GetResponsemanagementLibraryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetResponsemanagementLibraryNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementLibraryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryRequestTimeout creates a GetResponsemanagementLibraryRequestTimeout with default headers values
func NewGetResponsemanagementLibraryRequestTimeout() *GetResponsemanagementLibraryRequestTimeout {
	return &GetResponsemanagementLibraryRequestTimeout{}
}

/*GetResponsemanagementLibraryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetResponsemanagementLibraryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementLibraryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryRequestEntityTooLarge creates a GetResponsemanagementLibraryRequestEntityTooLarge with default headers values
func NewGetResponsemanagementLibraryRequestEntityTooLarge() *GetResponsemanagementLibraryRequestEntityTooLarge {
	return &GetResponsemanagementLibraryRequestEntityTooLarge{}
}

/*GetResponsemanagementLibraryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetResponsemanagementLibraryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementLibraryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryUnsupportedMediaType creates a GetResponsemanagementLibraryUnsupportedMediaType with default headers values
func NewGetResponsemanagementLibraryUnsupportedMediaType() *GetResponsemanagementLibraryUnsupportedMediaType {
	return &GetResponsemanagementLibraryUnsupportedMediaType{}
}

/*GetResponsemanagementLibraryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetResponsemanagementLibraryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementLibraryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryTooManyRequests creates a GetResponsemanagementLibraryTooManyRequests with default headers values
func NewGetResponsemanagementLibraryTooManyRequests() *GetResponsemanagementLibraryTooManyRequests {
	return &GetResponsemanagementLibraryTooManyRequests{}
}

/*GetResponsemanagementLibraryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetResponsemanagementLibraryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementLibraryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryInternalServerError creates a GetResponsemanagementLibraryInternalServerError with default headers values
func NewGetResponsemanagementLibraryInternalServerError() *GetResponsemanagementLibraryInternalServerError {
	return &GetResponsemanagementLibraryInternalServerError{}
}

/*GetResponsemanagementLibraryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetResponsemanagementLibraryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementLibraryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryServiceUnavailable creates a GetResponsemanagementLibraryServiceUnavailable with default headers values
func NewGetResponsemanagementLibraryServiceUnavailable() *GetResponsemanagementLibraryServiceUnavailable {
	return &GetResponsemanagementLibraryServiceUnavailable{}
}

/*GetResponsemanagementLibraryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetResponsemanagementLibraryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementLibraryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementLibraryGatewayTimeout creates a GetResponsemanagementLibraryGatewayTimeout with default headers values
func NewGetResponsemanagementLibraryGatewayTimeout() *GetResponsemanagementLibraryGatewayTimeout {
	return &GetResponsemanagementLibraryGatewayTimeout{}
}

/*GetResponsemanagementLibraryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetResponsemanagementLibraryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementLibraryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/libraries/{libraryId}][%d] getResponsemanagementLibraryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementLibraryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementLibraryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
