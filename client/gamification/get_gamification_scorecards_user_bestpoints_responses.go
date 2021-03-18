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

// GetGamificationScorecardsUserBestpointsReader is a Reader for the GetGamificationScorecardsUserBestpoints structure.
type GetGamificationScorecardsUserBestpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationScorecardsUserBestpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationScorecardsUserBestpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationScorecardsUserBestpointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationScorecardsUserBestpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationScorecardsUserBestpointsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationScorecardsUserBestpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationScorecardsUserBestpointsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationScorecardsUserBestpointsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationScorecardsUserBestpointsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationScorecardsUserBestpointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationScorecardsUserBestpointsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationScorecardsUserBestpointsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationScorecardsUserBestpointsOK creates a GetGamificationScorecardsUserBestpointsOK with default headers values
func NewGetGamificationScorecardsUserBestpointsOK() *GetGamificationScorecardsUserBestpointsOK {
	return &GetGamificationScorecardsUserBestpointsOK{}
}

/*GetGamificationScorecardsUserBestpointsOK handles this case with default header values.

successful operation
*/
type GetGamificationScorecardsUserBestpointsOK struct {
	Payload *models.UserBestPoints
}

func (o *GetGamificationScorecardsUserBestpointsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsOK) GetPayload() *models.UserBestPoints {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserBestPoints)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsBadRequest creates a GetGamificationScorecardsUserBestpointsBadRequest with default headers values
func NewGetGamificationScorecardsUserBestpointsBadRequest() *GetGamificationScorecardsUserBestpointsBadRequest {
	return &GetGamificationScorecardsUserBestpointsBadRequest{}
}

/*GetGamificationScorecardsUserBestpointsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationScorecardsUserBestpointsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsUnauthorized creates a GetGamificationScorecardsUserBestpointsUnauthorized with default headers values
func NewGetGamificationScorecardsUserBestpointsUnauthorized() *GetGamificationScorecardsUserBestpointsUnauthorized {
	return &GetGamificationScorecardsUserBestpointsUnauthorized{}
}

/*GetGamificationScorecardsUserBestpointsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationScorecardsUserBestpointsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsForbidden creates a GetGamificationScorecardsUserBestpointsForbidden with default headers values
func NewGetGamificationScorecardsUserBestpointsForbidden() *GetGamificationScorecardsUserBestpointsForbidden {
	return &GetGamificationScorecardsUserBestpointsForbidden{}
}

/*GetGamificationScorecardsUserBestpointsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationScorecardsUserBestpointsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsNotFound creates a GetGamificationScorecardsUserBestpointsNotFound with default headers values
func NewGetGamificationScorecardsUserBestpointsNotFound() *GetGamificationScorecardsUserBestpointsNotFound {
	return &GetGamificationScorecardsUserBestpointsNotFound{}
}

/*GetGamificationScorecardsUserBestpointsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationScorecardsUserBestpointsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsRequestEntityTooLarge creates a GetGamificationScorecardsUserBestpointsRequestEntityTooLarge with default headers values
func NewGetGamificationScorecardsUserBestpointsRequestEntityTooLarge() *GetGamificationScorecardsUserBestpointsRequestEntityTooLarge {
	return &GetGamificationScorecardsUserBestpointsRequestEntityTooLarge{}
}

/*GetGamificationScorecardsUserBestpointsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationScorecardsUserBestpointsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsUnsupportedMediaType creates a GetGamificationScorecardsUserBestpointsUnsupportedMediaType with default headers values
func NewGetGamificationScorecardsUserBestpointsUnsupportedMediaType() *GetGamificationScorecardsUserBestpointsUnsupportedMediaType {
	return &GetGamificationScorecardsUserBestpointsUnsupportedMediaType{}
}

/*GetGamificationScorecardsUserBestpointsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationScorecardsUserBestpointsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsTooManyRequests creates a GetGamificationScorecardsUserBestpointsTooManyRequests with default headers values
func NewGetGamificationScorecardsUserBestpointsTooManyRequests() *GetGamificationScorecardsUserBestpointsTooManyRequests {
	return &GetGamificationScorecardsUserBestpointsTooManyRequests{}
}

/*GetGamificationScorecardsUserBestpointsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGamificationScorecardsUserBestpointsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsInternalServerError creates a GetGamificationScorecardsUserBestpointsInternalServerError with default headers values
func NewGetGamificationScorecardsUserBestpointsInternalServerError() *GetGamificationScorecardsUserBestpointsInternalServerError {
	return &GetGamificationScorecardsUserBestpointsInternalServerError{}
}

/*GetGamificationScorecardsUserBestpointsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationScorecardsUserBestpointsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsServiceUnavailable creates a GetGamificationScorecardsUserBestpointsServiceUnavailable with default headers values
func NewGetGamificationScorecardsUserBestpointsServiceUnavailable() *GetGamificationScorecardsUserBestpointsServiceUnavailable {
	return &GetGamificationScorecardsUserBestpointsServiceUnavailable{}
}

/*GetGamificationScorecardsUserBestpointsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationScorecardsUserBestpointsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationScorecardsUserBestpointsGatewayTimeout creates a GetGamificationScorecardsUserBestpointsGatewayTimeout with default headers values
func NewGetGamificationScorecardsUserBestpointsGatewayTimeout() *GetGamificationScorecardsUserBestpointsGatewayTimeout {
	return &GetGamificationScorecardsUserBestpointsGatewayTimeout{}
}

/*GetGamificationScorecardsUserBestpointsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationScorecardsUserBestpointsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationScorecardsUserBestpointsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/scorecards/users/{userId}/bestpoints][%d] getGamificationScorecardsUserBestpointsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationScorecardsUserBestpointsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationScorecardsUserBestpointsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}