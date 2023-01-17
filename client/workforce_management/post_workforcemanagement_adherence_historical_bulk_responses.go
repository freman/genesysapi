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

// PostWorkforcemanagementAdherenceHistoricalBulkReader is a Reader for the PostWorkforcemanagementAdherenceHistoricalBulk structure.
type PostWorkforcemanagementAdherenceHistoricalBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementAdherenceHistoricalBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkAccepted creates a PostWorkforcemanagementAdherenceHistoricalBulkAccepted with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkAccepted() *PostWorkforcemanagementAdherenceHistoricalBulkAccepted {
	return &PostWorkforcemanagementAdherenceHistoricalBulkAccepted{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkAccepted describes a response with status code 202, with default header values.

Processing request
*/
type PostWorkforcemanagementAdherenceHistoricalBulkAccepted struct {
	Payload *models.WfmHistoricalAdherenceBulkResponse
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk accepted response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk accepted response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk accepted response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk accepted response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk accepted response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkAccepted  %+v", 202, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkAccepted  %+v", 202, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) GetPayload() *models.WfmHistoricalAdherenceBulkResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmHistoricalAdherenceBulkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkBadRequest creates a PostWorkforcemanagementAdherenceHistoricalBulkBadRequest with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkBadRequest() *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest {
	return &PostWorkforcemanagementAdherenceHistoricalBulkBadRequest{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk bad request response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk bad request response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk bad request response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk bad request response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk bad request response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkUnauthorized creates a PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkUnauthorized() *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized {
	return &PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkForbidden creates a PostWorkforcemanagementAdherenceHistoricalBulkForbidden with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkForbidden() *PostWorkforcemanagementAdherenceHistoricalBulkForbidden {
	return &PostWorkforcemanagementAdherenceHistoricalBulkForbidden{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk forbidden response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk forbidden response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk forbidden response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk forbidden response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk forbidden response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkNotFound creates a PostWorkforcemanagementAdherenceHistoricalBulkNotFound with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkNotFound() *PostWorkforcemanagementAdherenceHistoricalBulkNotFound {
	return &PostWorkforcemanagementAdherenceHistoricalBulkNotFound{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk not found response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk not found response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk not found response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk not found response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk not found response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout creates a PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout() *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout {
	return &PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk request timeout response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk request timeout response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk request timeout response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk request timeout response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk request timeout response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge creates a PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge() *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge {
	return &PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType creates a PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType() *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType {
	return &PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests creates a PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests() *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests {
	return &PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk too many requests response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk too many requests response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk too many requests response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk too many requests response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement adherence historical bulk too many requests response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkInternalServerError creates a PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkInternalServerError() *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError {
	return &PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk internal server error response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk internal server error response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk internal server error response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk internal server error response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement adherence historical bulk internal server error response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable creates a PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable() *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable {
	return &PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement adherence historical bulk service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout creates a PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout with default headers values
func NewPostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout() *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout {
	return &PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout{}
}

/*
PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement adherence historical bulk gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement adherence historical bulk gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement adherence historical bulk gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement adherence historical bulk gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement adherence historical bulk gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/adherence/historical/bulk][%d] postWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementAdherenceHistoricalBulkGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
