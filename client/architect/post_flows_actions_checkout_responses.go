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

// PostFlowsActionsCheckoutReader is a Reader for the PostFlowsActionsCheckout structure.
type PostFlowsActionsCheckoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlowsActionsCheckoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlowsActionsCheckoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFlowsActionsCheckoutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFlowsActionsCheckoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlowsActionsCheckoutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFlowsActionsCheckoutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostFlowsActionsCheckoutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostFlowsActionsCheckoutRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostFlowsActionsCheckoutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostFlowsActionsCheckoutGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostFlowsActionsCheckoutRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostFlowsActionsCheckoutUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFlowsActionsCheckoutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlowsActionsCheckoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFlowsActionsCheckoutServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFlowsActionsCheckoutGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlowsActionsCheckoutOK creates a PostFlowsActionsCheckoutOK with default headers values
func NewPostFlowsActionsCheckoutOK() *PostFlowsActionsCheckoutOK {
	return &PostFlowsActionsCheckoutOK{}
}

/*
PostFlowsActionsCheckoutOK describes a response with status code 200, with default header values.

successful operation
*/
type PostFlowsActionsCheckoutOK struct {
	Payload *models.Flow
}

// IsSuccess returns true when this post flows actions checkout o k response has a 2xx status code
func (o *PostFlowsActionsCheckoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post flows actions checkout o k response has a 3xx status code
func (o *PostFlowsActionsCheckoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout o k response has a 4xx status code
func (o *PostFlowsActionsCheckoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows actions checkout o k response has a 5xx status code
func (o *PostFlowsActionsCheckoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout o k response a status code equal to that given
func (o *PostFlowsActionsCheckoutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostFlowsActionsCheckoutOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutOK  %+v", 200, o.Payload)
}

func (o *PostFlowsActionsCheckoutOK) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutOK  %+v", 200, o.Payload)
}

func (o *PostFlowsActionsCheckoutOK) GetPayload() *models.Flow {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutBadRequest creates a PostFlowsActionsCheckoutBadRequest with default headers values
func NewPostFlowsActionsCheckoutBadRequest() *PostFlowsActionsCheckoutBadRequest {
	return &PostFlowsActionsCheckoutBadRequest{}
}

/*
PostFlowsActionsCheckoutBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostFlowsActionsCheckoutBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout bad request response has a 2xx status code
func (o *PostFlowsActionsCheckoutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout bad request response has a 3xx status code
func (o *PostFlowsActionsCheckoutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout bad request response has a 4xx status code
func (o *PostFlowsActionsCheckoutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout bad request response has a 5xx status code
func (o *PostFlowsActionsCheckoutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout bad request response a status code equal to that given
func (o *PostFlowsActionsCheckoutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostFlowsActionsCheckoutBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsActionsCheckoutBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsActionsCheckoutBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutUnauthorized creates a PostFlowsActionsCheckoutUnauthorized with default headers values
func NewPostFlowsActionsCheckoutUnauthorized() *PostFlowsActionsCheckoutUnauthorized {
	return &PostFlowsActionsCheckoutUnauthorized{}
}

/*
PostFlowsActionsCheckoutUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostFlowsActionsCheckoutUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout unauthorized response has a 2xx status code
func (o *PostFlowsActionsCheckoutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout unauthorized response has a 3xx status code
func (o *PostFlowsActionsCheckoutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout unauthorized response has a 4xx status code
func (o *PostFlowsActionsCheckoutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout unauthorized response has a 5xx status code
func (o *PostFlowsActionsCheckoutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout unauthorized response a status code equal to that given
func (o *PostFlowsActionsCheckoutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostFlowsActionsCheckoutUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsActionsCheckoutUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsActionsCheckoutUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutForbidden creates a PostFlowsActionsCheckoutForbidden with default headers values
func NewPostFlowsActionsCheckoutForbidden() *PostFlowsActionsCheckoutForbidden {
	return &PostFlowsActionsCheckoutForbidden{}
}

/*
PostFlowsActionsCheckoutForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostFlowsActionsCheckoutForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout forbidden response has a 2xx status code
func (o *PostFlowsActionsCheckoutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout forbidden response has a 3xx status code
func (o *PostFlowsActionsCheckoutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout forbidden response has a 4xx status code
func (o *PostFlowsActionsCheckoutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout forbidden response has a 5xx status code
func (o *PostFlowsActionsCheckoutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout forbidden response a status code equal to that given
func (o *PostFlowsActionsCheckoutForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostFlowsActionsCheckoutForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsActionsCheckoutForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsActionsCheckoutForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutNotFound creates a PostFlowsActionsCheckoutNotFound with default headers values
func NewPostFlowsActionsCheckoutNotFound() *PostFlowsActionsCheckoutNotFound {
	return &PostFlowsActionsCheckoutNotFound{}
}

/*
PostFlowsActionsCheckoutNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostFlowsActionsCheckoutNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout not found response has a 2xx status code
func (o *PostFlowsActionsCheckoutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout not found response has a 3xx status code
func (o *PostFlowsActionsCheckoutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout not found response has a 4xx status code
func (o *PostFlowsActionsCheckoutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout not found response has a 5xx status code
func (o *PostFlowsActionsCheckoutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout not found response a status code equal to that given
func (o *PostFlowsActionsCheckoutNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostFlowsActionsCheckoutNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsActionsCheckoutNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsActionsCheckoutNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutMethodNotAllowed creates a PostFlowsActionsCheckoutMethodNotAllowed with default headers values
func NewPostFlowsActionsCheckoutMethodNotAllowed() *PostFlowsActionsCheckoutMethodNotAllowed {
	return &PostFlowsActionsCheckoutMethodNotAllowed{}
}

/*
PostFlowsActionsCheckoutMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type PostFlowsActionsCheckoutMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout method not allowed response has a 2xx status code
func (o *PostFlowsActionsCheckoutMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout method not allowed response has a 3xx status code
func (o *PostFlowsActionsCheckoutMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout method not allowed response has a 4xx status code
func (o *PostFlowsActionsCheckoutMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout method not allowed response has a 5xx status code
func (o *PostFlowsActionsCheckoutMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout method not allowed response a status code equal to that given
func (o *PostFlowsActionsCheckoutMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *PostFlowsActionsCheckoutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowsActionsCheckoutMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostFlowsActionsCheckoutMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutRequestTimeout creates a PostFlowsActionsCheckoutRequestTimeout with default headers values
func NewPostFlowsActionsCheckoutRequestTimeout() *PostFlowsActionsCheckoutRequestTimeout {
	return &PostFlowsActionsCheckoutRequestTimeout{}
}

/*
PostFlowsActionsCheckoutRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostFlowsActionsCheckoutRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout request timeout response has a 2xx status code
func (o *PostFlowsActionsCheckoutRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout request timeout response has a 3xx status code
func (o *PostFlowsActionsCheckoutRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout request timeout response has a 4xx status code
func (o *PostFlowsActionsCheckoutRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout request timeout response has a 5xx status code
func (o *PostFlowsActionsCheckoutRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout request timeout response a status code equal to that given
func (o *PostFlowsActionsCheckoutRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostFlowsActionsCheckoutRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostFlowsActionsCheckoutRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostFlowsActionsCheckoutRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutConflict creates a PostFlowsActionsCheckoutConflict with default headers values
func NewPostFlowsActionsCheckoutConflict() *PostFlowsActionsCheckoutConflict {
	return &PostFlowsActionsCheckoutConflict{}
}

/*
PostFlowsActionsCheckoutConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostFlowsActionsCheckoutConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout conflict response has a 2xx status code
func (o *PostFlowsActionsCheckoutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout conflict response has a 3xx status code
func (o *PostFlowsActionsCheckoutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout conflict response has a 4xx status code
func (o *PostFlowsActionsCheckoutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout conflict response has a 5xx status code
func (o *PostFlowsActionsCheckoutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout conflict response a status code equal to that given
func (o *PostFlowsActionsCheckoutConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostFlowsActionsCheckoutConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutConflict  %+v", 409, o.Payload)
}

func (o *PostFlowsActionsCheckoutConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutConflict  %+v", 409, o.Payload)
}

func (o *PostFlowsActionsCheckoutConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutGone creates a PostFlowsActionsCheckoutGone with default headers values
func NewPostFlowsActionsCheckoutGone() *PostFlowsActionsCheckoutGone {
	return &PostFlowsActionsCheckoutGone{}
}

/*
PostFlowsActionsCheckoutGone describes a response with status code 410, with default header values.

Gone
*/
type PostFlowsActionsCheckoutGone struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout gone response has a 2xx status code
func (o *PostFlowsActionsCheckoutGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout gone response has a 3xx status code
func (o *PostFlowsActionsCheckoutGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout gone response has a 4xx status code
func (o *PostFlowsActionsCheckoutGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout gone response has a 5xx status code
func (o *PostFlowsActionsCheckoutGone) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout gone response a status code equal to that given
func (o *PostFlowsActionsCheckoutGone) IsCode(code int) bool {
	return code == 410
}

func (o *PostFlowsActionsCheckoutGone) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutGone  %+v", 410, o.Payload)
}

func (o *PostFlowsActionsCheckoutGone) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutGone  %+v", 410, o.Payload)
}

func (o *PostFlowsActionsCheckoutGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutRequestEntityTooLarge creates a PostFlowsActionsCheckoutRequestEntityTooLarge with default headers values
func NewPostFlowsActionsCheckoutRequestEntityTooLarge() *PostFlowsActionsCheckoutRequestEntityTooLarge {
	return &PostFlowsActionsCheckoutRequestEntityTooLarge{}
}

/*
PostFlowsActionsCheckoutRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostFlowsActionsCheckoutRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout request entity too large response has a 2xx status code
func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout request entity too large response has a 3xx status code
func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout request entity too large response has a 4xx status code
func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout request entity too large response has a 5xx status code
func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout request entity too large response a status code equal to that given
func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutUnsupportedMediaType creates a PostFlowsActionsCheckoutUnsupportedMediaType with default headers values
func NewPostFlowsActionsCheckoutUnsupportedMediaType() *PostFlowsActionsCheckoutUnsupportedMediaType {
	return &PostFlowsActionsCheckoutUnsupportedMediaType{}
}

/*
PostFlowsActionsCheckoutUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostFlowsActionsCheckoutUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout unsupported media type response has a 2xx status code
func (o *PostFlowsActionsCheckoutUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout unsupported media type response has a 3xx status code
func (o *PostFlowsActionsCheckoutUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout unsupported media type response has a 4xx status code
func (o *PostFlowsActionsCheckoutUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout unsupported media type response has a 5xx status code
func (o *PostFlowsActionsCheckoutUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout unsupported media type response a status code equal to that given
func (o *PostFlowsActionsCheckoutUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostFlowsActionsCheckoutUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsActionsCheckoutUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsActionsCheckoutUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutTooManyRequests creates a PostFlowsActionsCheckoutTooManyRequests with default headers values
func NewPostFlowsActionsCheckoutTooManyRequests() *PostFlowsActionsCheckoutTooManyRequests {
	return &PostFlowsActionsCheckoutTooManyRequests{}
}

/*
PostFlowsActionsCheckoutTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostFlowsActionsCheckoutTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout too many requests response has a 2xx status code
func (o *PostFlowsActionsCheckoutTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout too many requests response has a 3xx status code
func (o *PostFlowsActionsCheckoutTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout too many requests response has a 4xx status code
func (o *PostFlowsActionsCheckoutTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post flows actions checkout too many requests response has a 5xx status code
func (o *PostFlowsActionsCheckoutTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post flows actions checkout too many requests response a status code equal to that given
func (o *PostFlowsActionsCheckoutTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostFlowsActionsCheckoutTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsActionsCheckoutTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsActionsCheckoutTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutInternalServerError creates a PostFlowsActionsCheckoutInternalServerError with default headers values
func NewPostFlowsActionsCheckoutInternalServerError() *PostFlowsActionsCheckoutInternalServerError {
	return &PostFlowsActionsCheckoutInternalServerError{}
}

/*
PostFlowsActionsCheckoutInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostFlowsActionsCheckoutInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout internal server error response has a 2xx status code
func (o *PostFlowsActionsCheckoutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout internal server error response has a 3xx status code
func (o *PostFlowsActionsCheckoutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout internal server error response has a 4xx status code
func (o *PostFlowsActionsCheckoutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows actions checkout internal server error response has a 5xx status code
func (o *PostFlowsActionsCheckoutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post flows actions checkout internal server error response a status code equal to that given
func (o *PostFlowsActionsCheckoutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostFlowsActionsCheckoutInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsActionsCheckoutInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsActionsCheckoutInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutServiceUnavailable creates a PostFlowsActionsCheckoutServiceUnavailable with default headers values
func NewPostFlowsActionsCheckoutServiceUnavailable() *PostFlowsActionsCheckoutServiceUnavailable {
	return &PostFlowsActionsCheckoutServiceUnavailable{}
}

/*
PostFlowsActionsCheckoutServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostFlowsActionsCheckoutServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout service unavailable response has a 2xx status code
func (o *PostFlowsActionsCheckoutServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout service unavailable response has a 3xx status code
func (o *PostFlowsActionsCheckoutServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout service unavailable response has a 4xx status code
func (o *PostFlowsActionsCheckoutServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows actions checkout service unavailable response has a 5xx status code
func (o *PostFlowsActionsCheckoutServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post flows actions checkout service unavailable response a status code equal to that given
func (o *PostFlowsActionsCheckoutServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostFlowsActionsCheckoutServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsActionsCheckoutServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsActionsCheckoutServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsActionsCheckoutGatewayTimeout creates a PostFlowsActionsCheckoutGatewayTimeout with default headers values
func NewPostFlowsActionsCheckoutGatewayTimeout() *PostFlowsActionsCheckoutGatewayTimeout {
	return &PostFlowsActionsCheckoutGatewayTimeout{}
}

/*
PostFlowsActionsCheckoutGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostFlowsActionsCheckoutGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post flows actions checkout gateway timeout response has a 2xx status code
func (o *PostFlowsActionsCheckoutGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post flows actions checkout gateway timeout response has a 3xx status code
func (o *PostFlowsActionsCheckoutGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post flows actions checkout gateway timeout response has a 4xx status code
func (o *PostFlowsActionsCheckoutGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post flows actions checkout gateway timeout response has a 5xx status code
func (o *PostFlowsActionsCheckoutGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post flows actions checkout gateway timeout response a status code equal to that given
func (o *PostFlowsActionsCheckoutGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostFlowsActionsCheckoutGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsActionsCheckoutGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/flows/actions/checkout][%d] postFlowsActionsCheckoutGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsActionsCheckoutGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsActionsCheckoutGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
