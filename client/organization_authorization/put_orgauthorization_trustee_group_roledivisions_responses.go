// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOrgauthorizationTrusteeGroupRoledivisionsReader is a Reader for the PutOrgauthorizationTrusteeGroupRoledivisions structure.
type PutOrgauthorizationTrusteeGroupRoledivisionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsOK creates a PutOrgauthorizationTrusteeGroupRoledivisionsOK with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsOK() *PutOrgauthorizationTrusteeGroupRoledivisionsOK {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsOK{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsOK describes a response with status code 200, with default header values.

successful operation
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsOK struct {
	Payload *models.UserAuthorization
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions o k response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions o k response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions o k response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions o k response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions o k response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsOK  %+v", 200, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsOK  %+v", 200, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) GetPayload() *models.UserAuthorization {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsBadRequest creates a PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsBadRequest() *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions bad request response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions bad request response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions bad request response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions bad request response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions bad request response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsBadRequest  %+v", 400, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsBadRequest  %+v", 400, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized creates a PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized() *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions unauthorized response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions unauthorized response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions unauthorized response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions unauthorized response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions unauthorized response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsForbidden creates a PutOrgauthorizationTrusteeGroupRoledivisionsForbidden with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsForbidden() *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsForbidden{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions forbidden response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions forbidden response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions forbidden response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions forbidden response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions forbidden response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsForbidden  %+v", 403, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsForbidden  %+v", 403, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsNotFound creates a PutOrgauthorizationTrusteeGroupRoledivisionsNotFound with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsNotFound() *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsNotFound{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions not found response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions not found response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions not found response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions not found response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions not found response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsNotFound  %+v", 404, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsNotFound  %+v", 404, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout creates a PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout() *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions request timeout response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions request timeout response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions request timeout response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions request timeout response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions request timeout response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge creates a PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge() *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions request entity too large response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions request entity too large response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions request entity too large response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions request entity too large response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions request entity too large response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType creates a PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType() *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions unsupported media type response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions unsupported media type response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions unsupported media type response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions unsupported media type response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions unsupported media type response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests creates a PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests() *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions too many requests response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions too many requests response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions too many requests response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions too many requests response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put orgauthorization trustee group roledivisions too many requests response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError creates a PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError() *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions internal server error response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions internal server error response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions internal server error response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions internal server error response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put orgauthorization trustee group roledivisions internal server error response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable creates a PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable() *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions service unavailable response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions service unavailable response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions service unavailable response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions service unavailable response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put orgauthorization trustee group roledivisions service unavailable response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout creates a PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout with default headers values
func NewPutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout() *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout {
	return &PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout{}
}

/*
PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put orgauthorization trustee group roledivisions gateway timeout response has a 2xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put orgauthorization trustee group roledivisions gateway timeout response has a 3xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put orgauthorization trustee group roledivisions gateway timeout response has a 4xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put orgauthorization trustee group roledivisions gateway timeout response has a 5xx status code
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put orgauthorization trustee group roledivisions gateway timeout response a status code equal to that given
func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions][%d] putOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOrgauthorizationTrusteeGroupRoledivisionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}