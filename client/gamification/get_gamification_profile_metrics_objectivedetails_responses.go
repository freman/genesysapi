// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGamificationProfileMetricsObjectivedetailsReader is a Reader for the GetGamificationProfileMetricsObjectivedetails structure.
type GetGamificationProfileMetricsObjectivedetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationProfileMetricsObjectivedetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationProfileMetricsObjectivedetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationProfileMetricsObjectivedetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationProfileMetricsObjectivedetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationProfileMetricsObjectivedetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationProfileMetricsObjectivedetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationProfileMetricsObjectivedetailsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationProfileMetricsObjectivedetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationProfileMetricsObjectivedetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationProfileMetricsObjectivedetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationProfileMetricsObjectivedetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationProfileMetricsObjectivedetailsOK creates a GetGamificationProfileMetricsObjectivedetailsOK with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsOK() *GetGamificationProfileMetricsObjectivedetailsOK {
	return &GetGamificationProfileMetricsObjectivedetailsOK{}
}

/*GetGamificationProfileMetricsObjectivedetailsOK handles this case with default header values.

successful operation
*/
type GetGamificationProfileMetricsObjectivedetailsOK struct {
	Payload *models.GetMetricsResponse
}

func (o *GetGamificationProfileMetricsObjectivedetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsOK) GetPayload() *models.GetMetricsResponse {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsBadRequest creates a GetGamificationProfileMetricsObjectivedetailsBadRequest with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsBadRequest() *GetGamificationProfileMetricsObjectivedetailsBadRequest {
	return &GetGamificationProfileMetricsObjectivedetailsBadRequest{}
}

/*GetGamificationProfileMetricsObjectivedetailsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationProfileMetricsObjectivedetailsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsUnauthorized creates a GetGamificationProfileMetricsObjectivedetailsUnauthorized with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsUnauthorized() *GetGamificationProfileMetricsObjectivedetailsUnauthorized {
	return &GetGamificationProfileMetricsObjectivedetailsUnauthorized{}
}

/*GetGamificationProfileMetricsObjectivedetailsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationProfileMetricsObjectivedetailsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsForbidden creates a GetGamificationProfileMetricsObjectivedetailsForbidden with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsForbidden() *GetGamificationProfileMetricsObjectivedetailsForbidden {
	return &GetGamificationProfileMetricsObjectivedetailsForbidden{}
}

/*GetGamificationProfileMetricsObjectivedetailsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationProfileMetricsObjectivedetailsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsNotFound creates a GetGamificationProfileMetricsObjectivedetailsNotFound with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsNotFound() *GetGamificationProfileMetricsObjectivedetailsNotFound {
	return &GetGamificationProfileMetricsObjectivedetailsNotFound{}
}

/*GetGamificationProfileMetricsObjectivedetailsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationProfileMetricsObjectivedetailsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsRequestTimeout creates a GetGamificationProfileMetricsObjectivedetailsRequestTimeout with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsRequestTimeout() *GetGamificationProfileMetricsObjectivedetailsRequestTimeout {
	return &GetGamificationProfileMetricsObjectivedetailsRequestTimeout{}
}

/*GetGamificationProfileMetricsObjectivedetailsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationProfileMetricsObjectivedetailsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge creates a GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge() *GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge {
	return &GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge{}
}

/*GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType creates a GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType() *GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType {
	return &GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType{}
}

/*GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsTooManyRequests creates a GetGamificationProfileMetricsObjectivedetailsTooManyRequests with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsTooManyRequests() *GetGamificationProfileMetricsObjectivedetailsTooManyRequests {
	return &GetGamificationProfileMetricsObjectivedetailsTooManyRequests{}
}

/*GetGamificationProfileMetricsObjectivedetailsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationProfileMetricsObjectivedetailsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsInternalServerError creates a GetGamificationProfileMetricsObjectivedetailsInternalServerError with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsInternalServerError() *GetGamificationProfileMetricsObjectivedetailsInternalServerError {
	return &GetGamificationProfileMetricsObjectivedetailsInternalServerError{}
}

/*GetGamificationProfileMetricsObjectivedetailsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationProfileMetricsObjectivedetailsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsServiceUnavailable creates a GetGamificationProfileMetricsObjectivedetailsServiceUnavailable with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsServiceUnavailable() *GetGamificationProfileMetricsObjectivedetailsServiceUnavailable {
	return &GetGamificationProfileMetricsObjectivedetailsServiceUnavailable{}
}

/*GetGamificationProfileMetricsObjectivedetailsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationProfileMetricsObjectivedetailsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsObjectivedetailsGatewayTimeout creates a GetGamificationProfileMetricsObjectivedetailsGatewayTimeout with default headers values
func NewGetGamificationProfileMetricsObjectivedetailsGatewayTimeout() *GetGamificationProfileMetricsObjectivedetailsGatewayTimeout {
	return &GetGamificationProfileMetricsObjectivedetailsGatewayTimeout{}
}

/*GetGamificationProfileMetricsObjectivedetailsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationProfileMetricsObjectivedetailsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsObjectivedetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails][%d] getGamificationProfileMetricsObjectivedetailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationProfileMetricsObjectivedetailsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsObjectivedetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}