// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFlowsDivisionviewsReader is a Reader for the GetFlowsDivisionviews structure.
type GetFlowsDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowsDivisionviewsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetFlowsDivisionviewsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDivisionviewsOK creates a GetFlowsDivisionviewsOK with default headers values
func NewGetFlowsDivisionviewsOK() *GetFlowsDivisionviewsOK {
	return &GetFlowsDivisionviewsOK{}
}

/*GetFlowsDivisionviewsOK handles this case with default header values.

successful operation
*/
type GetFlowsDivisionviewsOK struct {
	Payload *models.FlowDivisionViewEntityListing
}

func (o *GetFlowsDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDivisionviewsOK) GetPayload() *models.FlowDivisionViewEntityListing {
	return o.Payload
}

func (o *GetFlowsDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowDivisionViewEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsBadRequest creates a GetFlowsDivisionviewsBadRequest with default headers values
func NewGetFlowsDivisionviewsBadRequest() *GetFlowsDivisionviewsBadRequest {
	return &GetFlowsDivisionviewsBadRequest{}
}

/*GetFlowsDivisionviewsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsUnauthorized creates a GetFlowsDivisionviewsUnauthorized with default headers values
func NewGetFlowsDivisionviewsUnauthorized() *GetFlowsDivisionviewsUnauthorized {
	return &GetFlowsDivisionviewsUnauthorized{}
}

/*GetFlowsDivisionviewsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsForbidden creates a GetFlowsDivisionviewsForbidden with default headers values
func NewGetFlowsDivisionviewsForbidden() *GetFlowsDivisionviewsForbidden {
	return &GetFlowsDivisionviewsForbidden{}
}

/*GetFlowsDivisionviewsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsNotFound creates a GetFlowsDivisionviewsNotFound with default headers values
func NewGetFlowsDivisionviewsNotFound() *GetFlowsDivisionviewsNotFound {
	return &GetFlowsDivisionviewsNotFound{}
}

/*GetFlowsDivisionviewsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsMethodNotAllowed creates a GetFlowsDivisionviewsMethodNotAllowed with default headers values
func NewGetFlowsDivisionviewsMethodNotAllowed() *GetFlowsDivisionviewsMethodNotAllowed {
	return &GetFlowsDivisionviewsMethodNotAllowed{}
}

/*GetFlowsDivisionviewsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetFlowsDivisionviewsMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowsDivisionviewsMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsRequestTimeout creates a GetFlowsDivisionviewsRequestTimeout with default headers values
func NewGetFlowsDivisionviewsRequestTimeout() *GetFlowsDivisionviewsRequestTimeout {
	return &GetFlowsDivisionviewsRequestTimeout{}
}

/*GetFlowsDivisionviewsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsRequestEntityTooLarge creates a GetFlowsDivisionviewsRequestEntityTooLarge with default headers values
func NewGetFlowsDivisionviewsRequestEntityTooLarge() *GetFlowsDivisionviewsRequestEntityTooLarge {
	return &GetFlowsDivisionviewsRequestEntityTooLarge{}
}

/*GetFlowsDivisionviewsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsUnsupportedMediaType creates a GetFlowsDivisionviewsUnsupportedMediaType with default headers values
func NewGetFlowsDivisionviewsUnsupportedMediaType() *GetFlowsDivisionviewsUnsupportedMediaType {
	return &GetFlowsDivisionviewsUnsupportedMediaType{}
}

/*GetFlowsDivisionviewsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsTooManyRequests creates a GetFlowsDivisionviewsTooManyRequests with default headers values
func NewGetFlowsDivisionviewsTooManyRequests() *GetFlowsDivisionviewsTooManyRequests {
	return &GetFlowsDivisionviewsTooManyRequests{}
}

/*GetFlowsDivisionviewsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsInternalServerError creates a GetFlowsDivisionviewsInternalServerError with default headers values
func NewGetFlowsDivisionviewsInternalServerError() *GetFlowsDivisionviewsInternalServerError {
	return &GetFlowsDivisionviewsInternalServerError{}
}

/*GetFlowsDivisionviewsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsNotImplemented creates a GetFlowsDivisionviewsNotImplemented with default headers values
func NewGetFlowsDivisionviewsNotImplemented() *GetFlowsDivisionviewsNotImplemented {
	return &GetFlowsDivisionviewsNotImplemented{}
}

/*GetFlowsDivisionviewsNotImplemented handles this case with default header values.

Not Implemented
*/
type GetFlowsDivisionviewsNotImplemented struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsNotImplemented) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsNotImplemented  %+v", 501, o.Payload)
}

func (o *GetFlowsDivisionviewsNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsServiceUnavailable creates a GetFlowsDivisionviewsServiceUnavailable with default headers values
func NewGetFlowsDivisionviewsServiceUnavailable() *GetFlowsDivisionviewsServiceUnavailable {
	return &GetFlowsDivisionviewsServiceUnavailable{}
}

/*GetFlowsDivisionviewsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDivisionviewsGatewayTimeout creates a GetFlowsDivisionviewsGatewayTimeout with default headers values
func NewGetFlowsDivisionviewsGatewayTimeout() *GetFlowsDivisionviewsGatewayTimeout {
	return &GetFlowsDivisionviewsGatewayTimeout{}
}

/*GetFlowsDivisionviewsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/divisionviews][%d] getFlowsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
