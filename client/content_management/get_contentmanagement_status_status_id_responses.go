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

// GetContentmanagementStatusStatusIDReader is a Reader for the GetContentmanagementStatusStatusID structure.
type GetContentmanagementStatusStatusIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementStatusStatusIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementStatusStatusIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementStatusStatusIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementStatusStatusIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementStatusStatusIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementStatusStatusIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementStatusStatusIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementStatusStatusIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementStatusStatusIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementStatusStatusIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementStatusStatusIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementStatusStatusIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementStatusStatusIDOK creates a GetContentmanagementStatusStatusIDOK with default headers values
func NewGetContentmanagementStatusStatusIDOK() *GetContentmanagementStatusStatusIDOK {
	return &GetContentmanagementStatusStatusIDOK{}
}

/*GetContentmanagementStatusStatusIDOK handles this case with default header values.

successful operation
*/
type GetContentmanagementStatusStatusIDOK struct {
	Payload *models.CommandStatus
}

func (o *GetContentmanagementStatusStatusIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDOK) GetPayload() *models.CommandStatus {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommandStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDBadRequest creates a GetContentmanagementStatusStatusIDBadRequest with default headers values
func NewGetContentmanagementStatusStatusIDBadRequest() *GetContentmanagementStatusStatusIDBadRequest {
	return &GetContentmanagementStatusStatusIDBadRequest{}
}

/*GetContentmanagementStatusStatusIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementStatusStatusIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDUnauthorized creates a GetContentmanagementStatusStatusIDUnauthorized with default headers values
func NewGetContentmanagementStatusStatusIDUnauthorized() *GetContentmanagementStatusStatusIDUnauthorized {
	return &GetContentmanagementStatusStatusIDUnauthorized{}
}

/*GetContentmanagementStatusStatusIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementStatusStatusIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDForbidden creates a GetContentmanagementStatusStatusIDForbidden with default headers values
func NewGetContentmanagementStatusStatusIDForbidden() *GetContentmanagementStatusStatusIDForbidden {
	return &GetContentmanagementStatusStatusIDForbidden{}
}

/*GetContentmanagementStatusStatusIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementStatusStatusIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDNotFound creates a GetContentmanagementStatusStatusIDNotFound with default headers values
func NewGetContentmanagementStatusStatusIDNotFound() *GetContentmanagementStatusStatusIDNotFound {
	return &GetContentmanagementStatusStatusIDNotFound{}
}

/*GetContentmanagementStatusStatusIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetContentmanagementStatusStatusIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDRequestEntityTooLarge creates a GetContentmanagementStatusStatusIDRequestEntityTooLarge with default headers values
func NewGetContentmanagementStatusStatusIDRequestEntityTooLarge() *GetContentmanagementStatusStatusIDRequestEntityTooLarge {
	return &GetContentmanagementStatusStatusIDRequestEntityTooLarge{}
}

/*GetContentmanagementStatusStatusIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetContentmanagementStatusStatusIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDUnsupportedMediaType creates a GetContentmanagementStatusStatusIDUnsupportedMediaType with default headers values
func NewGetContentmanagementStatusStatusIDUnsupportedMediaType() *GetContentmanagementStatusStatusIDUnsupportedMediaType {
	return &GetContentmanagementStatusStatusIDUnsupportedMediaType{}
}

/*GetContentmanagementStatusStatusIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementStatusStatusIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDTooManyRequests creates a GetContentmanagementStatusStatusIDTooManyRequests with default headers values
func NewGetContentmanagementStatusStatusIDTooManyRequests() *GetContentmanagementStatusStatusIDTooManyRequests {
	return &GetContentmanagementStatusStatusIDTooManyRequests{}
}

/*GetContentmanagementStatusStatusIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetContentmanagementStatusStatusIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDInternalServerError creates a GetContentmanagementStatusStatusIDInternalServerError with default headers values
func NewGetContentmanagementStatusStatusIDInternalServerError() *GetContentmanagementStatusStatusIDInternalServerError {
	return &GetContentmanagementStatusStatusIDInternalServerError{}
}

/*GetContentmanagementStatusStatusIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementStatusStatusIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDServiceUnavailable creates a GetContentmanagementStatusStatusIDServiceUnavailable with default headers values
func NewGetContentmanagementStatusStatusIDServiceUnavailable() *GetContentmanagementStatusStatusIDServiceUnavailable {
	return &GetContentmanagementStatusStatusIDServiceUnavailable{}
}

/*GetContentmanagementStatusStatusIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementStatusStatusIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementStatusStatusIDGatewayTimeout creates a GetContentmanagementStatusStatusIDGatewayTimeout with default headers values
func NewGetContentmanagementStatusStatusIDGatewayTimeout() *GetContentmanagementStatusStatusIDGatewayTimeout {
	return &GetContentmanagementStatusStatusIDGatewayTimeout{}
}

/*GetContentmanagementStatusStatusIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetContentmanagementStatusStatusIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementStatusStatusIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/status/{statusId}][%d] getContentmanagementStatusStatusIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementStatusStatusIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementStatusStatusIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}