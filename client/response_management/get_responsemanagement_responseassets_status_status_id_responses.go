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

// GetResponsemanagementResponseassetsStatusStatusIDReader is a Reader for the GetResponsemanagementResponseassetsStatusStatusID structure.
type GetResponsemanagementResponseassetsStatusStatusIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponsemanagementResponseassetsStatusStatusIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResponsemanagementResponseassetsStatusStatusIDOK creates a GetResponsemanagementResponseassetsStatusStatusIDOK with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDOK() *GetResponsemanagementResponseassetsStatusStatusIDOK {
	return &GetResponsemanagementResponseassetsStatusStatusIDOK{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDOK handles this case with default header values.

successful operation
*/
type GetResponsemanagementResponseassetsStatusStatusIDOK struct {
	Payload *models.ResponseAssetStatus
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDOK) GetPayload() *models.ResponseAssetStatus {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseAssetStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDBadRequest creates a GetResponsemanagementResponseassetsStatusStatusIDBadRequest with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDBadRequest() *GetResponsemanagementResponseassetsStatusStatusIDBadRequest {
	return &GetResponsemanagementResponseassetsStatusStatusIDBadRequest{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetResponsemanagementResponseassetsStatusStatusIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDUnauthorized creates a GetResponsemanagementResponseassetsStatusStatusIDUnauthorized with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDUnauthorized() *GetResponsemanagementResponseassetsStatusStatusIDUnauthorized {
	return &GetResponsemanagementResponseassetsStatusStatusIDUnauthorized{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetResponsemanagementResponseassetsStatusStatusIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDForbidden creates a GetResponsemanagementResponseassetsStatusStatusIDForbidden with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDForbidden() *GetResponsemanagementResponseassetsStatusStatusIDForbidden {
	return &GetResponsemanagementResponseassetsStatusStatusIDForbidden{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetResponsemanagementResponseassetsStatusStatusIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDNotFound creates a GetResponsemanagementResponseassetsStatusStatusIDNotFound with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDNotFound() *GetResponsemanagementResponseassetsStatusStatusIDNotFound {
	return &GetResponsemanagementResponseassetsStatusStatusIDNotFound{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetResponsemanagementResponseassetsStatusStatusIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDRequestTimeout creates a GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDRequestTimeout() *GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout {
	return &GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge creates a GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge() *GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge {
	return &GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType creates a GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType() *GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType {
	return &GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDTooManyRequests creates a GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDTooManyRequests() *GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests {
	return &GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDInternalServerError creates a GetResponsemanagementResponseassetsStatusStatusIDInternalServerError with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDInternalServerError() *GetResponsemanagementResponseassetsStatusStatusIDInternalServerError {
	return &GetResponsemanagementResponseassetsStatusStatusIDInternalServerError{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetResponsemanagementResponseassetsStatusStatusIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable creates a GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable() *GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable {
	return &GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout creates a GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout with default headers values
func NewGetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout() *GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout {
	return &GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout{}
}

/*GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responseassets/status/{statusId}][%d] getResponsemanagementResponseassetsStatusStatusIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseassetsStatusStatusIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}