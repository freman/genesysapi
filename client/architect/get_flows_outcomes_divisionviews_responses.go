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

// GetFlowsOutcomesDivisionviewsReader is a Reader for the GetFlowsOutcomesDivisionviews structure.
type GetFlowsOutcomesDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsOutcomesDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsOutcomesDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsOutcomesDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsOutcomesDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsOutcomesDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsOutcomesDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsOutcomesDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsOutcomesDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsOutcomesDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsOutcomesDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsOutcomesDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetFlowsOutcomesDivisionviewsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsOutcomesDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsOutcomesDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsOutcomesDivisionviewsOK creates a GetFlowsOutcomesDivisionviewsOK with default headers values
func NewGetFlowsOutcomesDivisionviewsOK() *GetFlowsOutcomesDivisionviewsOK {
	return &GetFlowsOutcomesDivisionviewsOK{}
}

/*GetFlowsOutcomesDivisionviewsOK handles this case with default header values.

successful operation
*/
type GetFlowsOutcomesDivisionviewsOK struct {
	Payload *models.FlowOutcomeDivisionViewEntityListing
}

func (o *GetFlowsOutcomesDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsOK) GetPayload() *models.FlowOutcomeDivisionViewEntityListing {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowOutcomeDivisionViewEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsBadRequest creates a GetFlowsOutcomesDivisionviewsBadRequest with default headers values
func NewGetFlowsOutcomesDivisionviewsBadRequest() *GetFlowsOutcomesDivisionviewsBadRequest {
	return &GetFlowsOutcomesDivisionviewsBadRequest{}
}

/*GetFlowsOutcomesDivisionviewsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsOutcomesDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsUnauthorized creates a GetFlowsOutcomesDivisionviewsUnauthorized with default headers values
func NewGetFlowsOutcomesDivisionviewsUnauthorized() *GetFlowsOutcomesDivisionviewsUnauthorized {
	return &GetFlowsOutcomesDivisionviewsUnauthorized{}
}

/*GetFlowsOutcomesDivisionviewsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsOutcomesDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsForbidden creates a GetFlowsOutcomesDivisionviewsForbidden with default headers values
func NewGetFlowsOutcomesDivisionviewsForbidden() *GetFlowsOutcomesDivisionviewsForbidden {
	return &GetFlowsOutcomesDivisionviewsForbidden{}
}

/*GetFlowsOutcomesDivisionviewsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsOutcomesDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsNotFound creates a GetFlowsOutcomesDivisionviewsNotFound with default headers values
func NewGetFlowsOutcomesDivisionviewsNotFound() *GetFlowsOutcomesDivisionviewsNotFound {
	return &GetFlowsOutcomesDivisionviewsNotFound{}
}

/*GetFlowsOutcomesDivisionviewsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsOutcomesDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsRequestTimeout creates a GetFlowsOutcomesDivisionviewsRequestTimeout with default headers values
func NewGetFlowsOutcomesDivisionviewsRequestTimeout() *GetFlowsOutcomesDivisionviewsRequestTimeout {
	return &GetFlowsOutcomesDivisionviewsRequestTimeout{}
}

/*GetFlowsOutcomesDivisionviewsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsOutcomesDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsRequestEntityTooLarge creates a GetFlowsOutcomesDivisionviewsRequestEntityTooLarge with default headers values
func NewGetFlowsOutcomesDivisionviewsRequestEntityTooLarge() *GetFlowsOutcomesDivisionviewsRequestEntityTooLarge {
	return &GetFlowsOutcomesDivisionviewsRequestEntityTooLarge{}
}

/*GetFlowsOutcomesDivisionviewsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsOutcomesDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsUnsupportedMediaType creates a GetFlowsOutcomesDivisionviewsUnsupportedMediaType with default headers values
func NewGetFlowsOutcomesDivisionviewsUnsupportedMediaType() *GetFlowsOutcomesDivisionviewsUnsupportedMediaType {
	return &GetFlowsOutcomesDivisionviewsUnsupportedMediaType{}
}

/*GetFlowsOutcomesDivisionviewsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsOutcomesDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsTooManyRequests creates a GetFlowsOutcomesDivisionviewsTooManyRequests with default headers values
func NewGetFlowsOutcomesDivisionviewsTooManyRequests() *GetFlowsOutcomesDivisionviewsTooManyRequests {
	return &GetFlowsOutcomesDivisionviewsTooManyRequests{}
}

/*GetFlowsOutcomesDivisionviewsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsOutcomesDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsInternalServerError creates a GetFlowsOutcomesDivisionviewsInternalServerError with default headers values
func NewGetFlowsOutcomesDivisionviewsInternalServerError() *GetFlowsOutcomesDivisionviewsInternalServerError {
	return &GetFlowsOutcomesDivisionviewsInternalServerError{}
}

/*GetFlowsOutcomesDivisionviewsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsOutcomesDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsNotImplemented creates a GetFlowsOutcomesDivisionviewsNotImplemented with default headers values
func NewGetFlowsOutcomesDivisionviewsNotImplemented() *GetFlowsOutcomesDivisionviewsNotImplemented {
	return &GetFlowsOutcomesDivisionviewsNotImplemented{}
}

/*GetFlowsOutcomesDivisionviewsNotImplemented handles this case with default header values.

Not Implemented
*/
type GetFlowsOutcomesDivisionviewsNotImplemented struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsNotImplemented) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsNotImplemented  %+v", 501, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsServiceUnavailable creates a GetFlowsOutcomesDivisionviewsServiceUnavailable with default headers values
func NewGetFlowsOutcomesDivisionviewsServiceUnavailable() *GetFlowsOutcomesDivisionviewsServiceUnavailable {
	return &GetFlowsOutcomesDivisionviewsServiceUnavailable{}
}

/*GetFlowsOutcomesDivisionviewsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsOutcomesDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesDivisionviewsGatewayTimeout creates a GetFlowsOutcomesDivisionviewsGatewayTimeout with default headers values
func NewGetFlowsOutcomesDivisionviewsGatewayTimeout() *GetFlowsOutcomesDivisionviewsGatewayTimeout {
	return &GetFlowsOutcomesDivisionviewsGatewayTimeout{}
}

/*GetFlowsOutcomesDivisionviewsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsOutcomesDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes/divisionviews][%d] getFlowsOutcomesDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsOutcomesDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
