// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteOutboundCampaignReader is a Reader for the DeleteOutboundCampaign structure.
type DeleteOutboundCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundCampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundCampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundCampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundCampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundCampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundCampaignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundCampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundCampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundCampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundCampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundCampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundCampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundCampaignOK creates a DeleteOutboundCampaignOK with default headers values
func NewDeleteOutboundCampaignOK() *DeleteOutboundCampaignOK {
	return &DeleteOutboundCampaignOK{}
}

/*
DeleteOutboundCampaignOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteOutboundCampaignOK struct {
	Payload *models.Campaign
}

// IsSuccess returns true when this delete outbound campaign o k response has a 2xx status code
func (o *DeleteOutboundCampaignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete outbound campaign o k response has a 3xx status code
func (o *DeleteOutboundCampaignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign o k response has a 4xx status code
func (o *DeleteOutboundCampaignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound campaign o k response has a 5xx status code
func (o *DeleteOutboundCampaignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign o k response a status code equal to that given
func (o *DeleteOutboundCampaignOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOutboundCampaignOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignOK  %+v", 200, o.Payload)
}

func (o *DeleteOutboundCampaignOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignOK  %+v", 200, o.Payload)
}

func (o *DeleteOutboundCampaignOK) GetPayload() *models.Campaign {
	return o.Payload
}

func (o *DeleteOutboundCampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Campaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignBadRequest creates a DeleteOutboundCampaignBadRequest with default headers values
func NewDeleteOutboundCampaignBadRequest() *DeleteOutboundCampaignBadRequest {
	return &DeleteOutboundCampaignBadRequest{}
}

/*
DeleteOutboundCampaignBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundCampaignBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign bad request response has a 2xx status code
func (o *DeleteOutboundCampaignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign bad request response has a 3xx status code
func (o *DeleteOutboundCampaignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign bad request response has a 4xx status code
func (o *DeleteOutboundCampaignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign bad request response has a 5xx status code
func (o *DeleteOutboundCampaignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign bad request response a status code equal to that given
func (o *DeleteOutboundCampaignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteOutboundCampaignBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundCampaignBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundCampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignUnauthorized creates a DeleteOutboundCampaignUnauthorized with default headers values
func NewDeleteOutboundCampaignUnauthorized() *DeleteOutboundCampaignUnauthorized {
	return &DeleteOutboundCampaignUnauthorized{}
}

/*
DeleteOutboundCampaignUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundCampaignUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign unauthorized response has a 2xx status code
func (o *DeleteOutboundCampaignUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign unauthorized response has a 3xx status code
func (o *DeleteOutboundCampaignUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign unauthorized response has a 4xx status code
func (o *DeleteOutboundCampaignUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign unauthorized response has a 5xx status code
func (o *DeleteOutboundCampaignUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign unauthorized response a status code equal to that given
func (o *DeleteOutboundCampaignUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteOutboundCampaignUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundCampaignUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundCampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignForbidden creates a DeleteOutboundCampaignForbidden with default headers values
func NewDeleteOutboundCampaignForbidden() *DeleteOutboundCampaignForbidden {
	return &DeleteOutboundCampaignForbidden{}
}

/*
DeleteOutboundCampaignForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundCampaignForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign forbidden response has a 2xx status code
func (o *DeleteOutboundCampaignForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign forbidden response has a 3xx status code
func (o *DeleteOutboundCampaignForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign forbidden response has a 4xx status code
func (o *DeleteOutboundCampaignForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign forbidden response has a 5xx status code
func (o *DeleteOutboundCampaignForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign forbidden response a status code equal to that given
func (o *DeleteOutboundCampaignForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteOutboundCampaignForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundCampaignForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundCampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignNotFound creates a DeleteOutboundCampaignNotFound with default headers values
func NewDeleteOutboundCampaignNotFound() *DeleteOutboundCampaignNotFound {
	return &DeleteOutboundCampaignNotFound{}
}

/*
DeleteOutboundCampaignNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteOutboundCampaignNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign not found response has a 2xx status code
func (o *DeleteOutboundCampaignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign not found response has a 3xx status code
func (o *DeleteOutboundCampaignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign not found response has a 4xx status code
func (o *DeleteOutboundCampaignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign not found response has a 5xx status code
func (o *DeleteOutboundCampaignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign not found response a status code equal to that given
func (o *DeleteOutboundCampaignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteOutboundCampaignNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundCampaignNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundCampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignRequestTimeout creates a DeleteOutboundCampaignRequestTimeout with default headers values
func NewDeleteOutboundCampaignRequestTimeout() *DeleteOutboundCampaignRequestTimeout {
	return &DeleteOutboundCampaignRequestTimeout{}
}

/*
DeleteOutboundCampaignRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundCampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign request timeout response has a 2xx status code
func (o *DeleteOutboundCampaignRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign request timeout response has a 3xx status code
func (o *DeleteOutboundCampaignRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign request timeout response has a 4xx status code
func (o *DeleteOutboundCampaignRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign request timeout response has a 5xx status code
func (o *DeleteOutboundCampaignRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign request timeout response a status code equal to that given
func (o *DeleteOutboundCampaignRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteOutboundCampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundCampaignRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundCampaignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignRequestEntityTooLarge creates a DeleteOutboundCampaignRequestEntityTooLarge with default headers values
func NewDeleteOutboundCampaignRequestEntityTooLarge() *DeleteOutboundCampaignRequestEntityTooLarge {
	return &DeleteOutboundCampaignRequestEntityTooLarge{}
}

/*
DeleteOutboundCampaignRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteOutboundCampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign request entity too large response has a 2xx status code
func (o *DeleteOutboundCampaignRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign request entity too large response has a 3xx status code
func (o *DeleteOutboundCampaignRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign request entity too large response has a 4xx status code
func (o *DeleteOutboundCampaignRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign request entity too large response has a 5xx status code
func (o *DeleteOutboundCampaignRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign request entity too large response a status code equal to that given
func (o *DeleteOutboundCampaignRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteOutboundCampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundCampaignRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundCampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignUnsupportedMediaType creates a DeleteOutboundCampaignUnsupportedMediaType with default headers values
func NewDeleteOutboundCampaignUnsupportedMediaType() *DeleteOutboundCampaignUnsupportedMediaType {
	return &DeleteOutboundCampaignUnsupportedMediaType{}
}

/*
DeleteOutboundCampaignUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundCampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign unsupported media type response has a 2xx status code
func (o *DeleteOutboundCampaignUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign unsupported media type response has a 3xx status code
func (o *DeleteOutboundCampaignUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign unsupported media type response has a 4xx status code
func (o *DeleteOutboundCampaignUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign unsupported media type response has a 5xx status code
func (o *DeleteOutboundCampaignUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign unsupported media type response a status code equal to that given
func (o *DeleteOutboundCampaignUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteOutboundCampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundCampaignUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundCampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignTooManyRequests creates a DeleteOutboundCampaignTooManyRequests with default headers values
func NewDeleteOutboundCampaignTooManyRequests() *DeleteOutboundCampaignTooManyRequests {
	return &DeleteOutboundCampaignTooManyRequests{}
}

/*
DeleteOutboundCampaignTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundCampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign too many requests response has a 2xx status code
func (o *DeleteOutboundCampaignTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign too many requests response has a 3xx status code
func (o *DeleteOutboundCampaignTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign too many requests response has a 4xx status code
func (o *DeleteOutboundCampaignTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound campaign too many requests response has a 5xx status code
func (o *DeleteOutboundCampaignTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound campaign too many requests response a status code equal to that given
func (o *DeleteOutboundCampaignTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteOutboundCampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundCampaignTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundCampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignInternalServerError creates a DeleteOutboundCampaignInternalServerError with default headers values
func NewDeleteOutboundCampaignInternalServerError() *DeleteOutboundCampaignInternalServerError {
	return &DeleteOutboundCampaignInternalServerError{}
}

/*
DeleteOutboundCampaignInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundCampaignInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign internal server error response has a 2xx status code
func (o *DeleteOutboundCampaignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign internal server error response has a 3xx status code
func (o *DeleteOutboundCampaignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign internal server error response has a 4xx status code
func (o *DeleteOutboundCampaignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound campaign internal server error response has a 5xx status code
func (o *DeleteOutboundCampaignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound campaign internal server error response a status code equal to that given
func (o *DeleteOutboundCampaignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteOutboundCampaignInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundCampaignInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundCampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignServiceUnavailable creates a DeleteOutboundCampaignServiceUnavailable with default headers values
func NewDeleteOutboundCampaignServiceUnavailable() *DeleteOutboundCampaignServiceUnavailable {
	return &DeleteOutboundCampaignServiceUnavailable{}
}

/*
DeleteOutboundCampaignServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundCampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign service unavailable response has a 2xx status code
func (o *DeleteOutboundCampaignServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign service unavailable response has a 3xx status code
func (o *DeleteOutboundCampaignServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign service unavailable response has a 4xx status code
func (o *DeleteOutboundCampaignServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound campaign service unavailable response has a 5xx status code
func (o *DeleteOutboundCampaignServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound campaign service unavailable response a status code equal to that given
func (o *DeleteOutboundCampaignServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteOutboundCampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundCampaignServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundCampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundCampaignGatewayTimeout creates a DeleteOutboundCampaignGatewayTimeout with default headers values
func NewDeleteOutboundCampaignGatewayTimeout() *DeleteOutboundCampaignGatewayTimeout {
	return &DeleteOutboundCampaignGatewayTimeout{}
}

/*
DeleteOutboundCampaignGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteOutboundCampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound campaign gateway timeout response has a 2xx status code
func (o *DeleteOutboundCampaignGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound campaign gateway timeout response has a 3xx status code
func (o *DeleteOutboundCampaignGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound campaign gateway timeout response has a 4xx status code
func (o *DeleteOutboundCampaignGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound campaign gateway timeout response has a 5xx status code
func (o *DeleteOutboundCampaignGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound campaign gateway timeout response a status code equal to that given
func (o *DeleteOutboundCampaignGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteOutboundCampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundCampaignGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/campaigns/{campaignId}][%d] deleteOutboundCampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundCampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundCampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
