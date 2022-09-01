// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostTeamsReader is a Reader for the PostTeams structure.
type PostTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTeamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTeamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTeamsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTeamsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTeamsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTeamsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTeamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTeamsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTeamsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTeamsOK creates a PostTeamsOK with default headers values
func NewPostTeamsOK() *PostTeamsOK {
	return &PostTeamsOK{}
}

/*PostTeamsOK handles this case with default header values.

successful operation
*/
type PostTeamsOK struct {
	Payload *models.Team
}

func (o *PostTeamsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsOK  %+v", 200, o.Payload)
}

func (o *PostTeamsOK) GetPayload() *models.Team {
	return o.Payload
}

func (o *PostTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsBadRequest creates a PostTeamsBadRequest with default headers values
func NewPostTeamsBadRequest() *PostTeamsBadRequest {
	return &PostTeamsBadRequest{}
}

/*PostTeamsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTeamsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsBadRequest  %+v", 400, o.Payload)
}

func (o *PostTeamsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsUnauthorized creates a PostTeamsUnauthorized with default headers values
func NewPostTeamsUnauthorized() *PostTeamsUnauthorized {
	return &PostTeamsUnauthorized{}
}

/*PostTeamsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTeamsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTeamsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsForbidden creates a PostTeamsForbidden with default headers values
func NewPostTeamsForbidden() *PostTeamsForbidden {
	return &PostTeamsForbidden{}
}

/*PostTeamsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTeamsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsForbidden  %+v", 403, o.Payload)
}

func (o *PostTeamsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsNotFound creates a PostTeamsNotFound with default headers values
func NewPostTeamsNotFound() *PostTeamsNotFound {
	return &PostTeamsNotFound{}
}

/*PostTeamsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTeamsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsNotFound  %+v", 404, o.Payload)
}

func (o *PostTeamsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsRequestTimeout creates a PostTeamsRequestTimeout with default headers values
func NewPostTeamsRequestTimeout() *PostTeamsRequestTimeout {
	return &PostTeamsRequestTimeout{}
}

/*PostTeamsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTeamsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTeamsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsRequestEntityTooLarge creates a PostTeamsRequestEntityTooLarge with default headers values
func NewPostTeamsRequestEntityTooLarge() *PostTeamsRequestEntityTooLarge {
	return &PostTeamsRequestEntityTooLarge{}
}

/*PostTeamsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostTeamsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTeamsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsUnsupportedMediaType creates a PostTeamsUnsupportedMediaType with default headers values
func NewPostTeamsUnsupportedMediaType() *PostTeamsUnsupportedMediaType {
	return &PostTeamsUnsupportedMediaType{}
}

/*PostTeamsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTeamsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTeamsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsTooManyRequests creates a PostTeamsTooManyRequests with default headers values
func NewPostTeamsTooManyRequests() *PostTeamsTooManyRequests {
	return &PostTeamsTooManyRequests{}
}

/*PostTeamsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTeamsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTeamsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsInternalServerError creates a PostTeamsInternalServerError with default headers values
func NewPostTeamsInternalServerError() *PostTeamsInternalServerError {
	return &PostTeamsInternalServerError{}
}

/*PostTeamsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTeamsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTeamsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsServiceUnavailable creates a PostTeamsServiceUnavailable with default headers values
func NewPostTeamsServiceUnavailable() *PostTeamsServiceUnavailable {
	return &PostTeamsServiceUnavailable{}
}

/*PostTeamsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTeamsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTeamsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsGatewayTimeout creates a PostTeamsGatewayTimeout with default headers values
func NewPostTeamsGatewayTimeout() *PostTeamsGatewayTimeout {
	return &PostTeamsGatewayTimeout{}
}

/*PostTeamsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTeamsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTeamsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams][%d] postTeamsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTeamsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}