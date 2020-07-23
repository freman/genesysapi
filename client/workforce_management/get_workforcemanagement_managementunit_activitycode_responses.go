// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementManagementunitActivitycodeReader is a Reader for the GetWorkforcemanagementManagementunitActivitycode structure.
type GetWorkforcemanagementManagementunitActivitycodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitActivitycodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitActivitycodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitActivitycodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitActivitycodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitActivitycodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitActivitycodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetWorkforcemanagementManagementunitActivitycodeGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitActivitycodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitActivitycodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitActivitycodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitActivitycodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitActivitycodeOK creates a GetWorkforcemanagementManagementunitActivitycodeOK with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeOK() *GetWorkforcemanagementManagementunitActivitycodeOK {
	return &GetWorkforcemanagementManagementunitActivitycodeOK{}
}

/*GetWorkforcemanagementManagementunitActivitycodeOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitActivitycodeOK struct {
	Payload *models.ActivityCode
}

func (o *GetWorkforcemanagementManagementunitActivitycodeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeOK) GetPayload() *models.ActivityCode {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeBadRequest creates a GetWorkforcemanagementManagementunitActivitycodeBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeBadRequest() *GetWorkforcemanagementManagementunitActivitycodeBadRequest {
	return &GetWorkforcemanagementManagementunitActivitycodeBadRequest{}
}

/*GetWorkforcemanagementManagementunitActivitycodeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitActivitycodeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeUnauthorized creates a GetWorkforcemanagementManagementunitActivitycodeUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeUnauthorized() *GetWorkforcemanagementManagementunitActivitycodeUnauthorized {
	return &GetWorkforcemanagementManagementunitActivitycodeUnauthorized{}
}

/*GetWorkforcemanagementManagementunitActivitycodeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitActivitycodeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeForbidden creates a GetWorkforcemanagementManagementunitActivitycodeForbidden with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeForbidden() *GetWorkforcemanagementManagementunitActivitycodeForbidden {
	return &GetWorkforcemanagementManagementunitActivitycodeForbidden{}
}

/*GetWorkforcemanagementManagementunitActivitycodeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitActivitycodeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeNotFound creates a GetWorkforcemanagementManagementunitActivitycodeNotFound with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeNotFound() *GetWorkforcemanagementManagementunitActivitycodeNotFound {
	return &GetWorkforcemanagementManagementunitActivitycodeNotFound{}
}

/*GetWorkforcemanagementManagementunitActivitycodeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitActivitycodeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeGone creates a GetWorkforcemanagementManagementunitActivitycodeGone with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeGone() *GetWorkforcemanagementManagementunitActivitycodeGone {
	return &GetWorkforcemanagementManagementunitActivitycodeGone{}
}

/*GetWorkforcemanagementManagementunitActivitycodeGone handles this case with default header values.

Gone
*/
type GetWorkforcemanagementManagementunitActivitycodeGone struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeGone) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeGone  %+v", 410, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge() *GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType creates a GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType() *GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeTooManyRequests creates a GetWorkforcemanagementManagementunitActivitycodeTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeTooManyRequests() *GetWorkforcemanagementManagementunitActivitycodeTooManyRequests {
	return &GetWorkforcemanagementManagementunitActivitycodeTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitActivitycodeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitActivitycodeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeInternalServerError creates a GetWorkforcemanagementManagementunitActivitycodeInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeInternalServerError() *GetWorkforcemanagementManagementunitActivitycodeInternalServerError {
	return &GetWorkforcemanagementManagementunitActivitycodeInternalServerError{}
}

/*GetWorkforcemanagementManagementunitActivitycodeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitActivitycodeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeServiceUnavailable creates a GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeServiceUnavailable() *GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable {
	return &GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodeGatewayTimeout creates a GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodeGatewayTimeout() *GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout {
	return &GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] getWorkforcemanagementManagementunitActivitycodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
