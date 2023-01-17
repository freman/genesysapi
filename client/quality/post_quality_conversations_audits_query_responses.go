// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostQualityConversationsAuditsQueryReader is a Reader for the PostQualityConversationsAuditsQuery structure.
type PostQualityConversationsAuditsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityConversationsAuditsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityConversationsAuditsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostQualityConversationsAuditsQueryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityConversationsAuditsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityConversationsAuditsQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityConversationsAuditsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityConversationsAuditsQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityConversationsAuditsQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityConversationsAuditsQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityConversationsAuditsQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityConversationsAuditsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityConversationsAuditsQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityConversationsAuditsQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityConversationsAuditsQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityConversationsAuditsQueryOK creates a PostQualityConversationsAuditsQueryOK with default headers values
func NewPostQualityConversationsAuditsQueryOK() *PostQualityConversationsAuditsQueryOK {
	return &PostQualityConversationsAuditsQueryOK{}
}

/*
PostQualityConversationsAuditsQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type PostQualityConversationsAuditsQueryOK struct {
	Payload *models.QualityAuditQueryExecutionStatusResponse
}

// IsSuccess returns true when this post quality conversations audits query o k response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post quality conversations audits query o k response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query o k response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality conversations audits query o k response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query o k response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostQualityConversationsAuditsQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryOK  %+v", 200, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryOK) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryOK  %+v", 200, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryOK) GetPayload() *models.QualityAuditQueryExecutionStatusResponse {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QualityAuditQueryExecutionStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryAccepted creates a PostQualityConversationsAuditsQueryAccepted with default headers values
func NewPostQualityConversationsAuditsQueryAccepted() *PostQualityConversationsAuditsQueryAccepted {
	return &PostQualityConversationsAuditsQueryAccepted{}
}

/*
PostQualityConversationsAuditsQueryAccepted describes a response with status code 202, with default header values.

Accepted - Query execution is accepted.
*/
type PostQualityConversationsAuditsQueryAccepted struct {
	Payload *models.QualityAuditQueryExecutionStatusResponse
}

// IsSuccess returns true when this post quality conversations audits query accepted response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post quality conversations audits query accepted response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query accepted response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality conversations audits query accepted response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query accepted response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostQualityConversationsAuditsQueryAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryAccepted  %+v", 202, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryAccepted  %+v", 202, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryAccepted) GetPayload() *models.QualityAuditQueryExecutionStatusResponse {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QualityAuditQueryExecutionStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryBadRequest creates a PostQualityConversationsAuditsQueryBadRequest with default headers values
func NewPostQualityConversationsAuditsQueryBadRequest() *PostQualityConversationsAuditsQueryBadRequest {
	return &PostQualityConversationsAuditsQueryBadRequest{}
}

/*
PostQualityConversationsAuditsQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityConversationsAuditsQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query bad request response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query bad request response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query bad request response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query bad request response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query bad request response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostQualityConversationsAuditsQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryUnauthorized creates a PostQualityConversationsAuditsQueryUnauthorized with default headers values
func NewPostQualityConversationsAuditsQueryUnauthorized() *PostQualityConversationsAuditsQueryUnauthorized {
	return &PostQualityConversationsAuditsQueryUnauthorized{}
}

/*
PostQualityConversationsAuditsQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityConversationsAuditsQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query unauthorized response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query unauthorized response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query unauthorized response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query unauthorized response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query unauthorized response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostQualityConversationsAuditsQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryForbidden creates a PostQualityConversationsAuditsQueryForbidden with default headers values
func NewPostQualityConversationsAuditsQueryForbidden() *PostQualityConversationsAuditsQueryForbidden {
	return &PostQualityConversationsAuditsQueryForbidden{}
}

/*
PostQualityConversationsAuditsQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityConversationsAuditsQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query forbidden response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query forbidden response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query forbidden response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query forbidden response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query forbidden response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostQualityConversationsAuditsQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryNotFound creates a PostQualityConversationsAuditsQueryNotFound with default headers values
func NewPostQualityConversationsAuditsQueryNotFound() *PostQualityConversationsAuditsQueryNotFound {
	return &PostQualityConversationsAuditsQueryNotFound{}
}

/*
PostQualityConversationsAuditsQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostQualityConversationsAuditsQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query not found response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query not found response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query not found response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query not found response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query not found response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostQualityConversationsAuditsQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryRequestTimeout creates a PostQualityConversationsAuditsQueryRequestTimeout with default headers values
func NewPostQualityConversationsAuditsQueryRequestTimeout() *PostQualityConversationsAuditsQueryRequestTimeout {
	return &PostQualityConversationsAuditsQueryRequestTimeout{}
}

/*
PostQualityConversationsAuditsQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityConversationsAuditsQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query request timeout response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query request timeout response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query request timeout response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query request timeout response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query request timeout response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostQualityConversationsAuditsQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryRequestEntityTooLarge creates a PostQualityConversationsAuditsQueryRequestEntityTooLarge with default headers values
func NewPostQualityConversationsAuditsQueryRequestEntityTooLarge() *PostQualityConversationsAuditsQueryRequestEntityTooLarge {
	return &PostQualityConversationsAuditsQueryRequestEntityTooLarge{}
}

/*
PostQualityConversationsAuditsQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostQualityConversationsAuditsQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query request entity too large response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query request entity too large response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query request entity too large response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query request entity too large response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query request entity too large response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryUnsupportedMediaType creates a PostQualityConversationsAuditsQueryUnsupportedMediaType with default headers values
func NewPostQualityConversationsAuditsQueryUnsupportedMediaType() *PostQualityConversationsAuditsQueryUnsupportedMediaType {
	return &PostQualityConversationsAuditsQueryUnsupportedMediaType{}
}

/*
PostQualityConversationsAuditsQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityConversationsAuditsQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query unsupported media type response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query unsupported media type response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query unsupported media type response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query unsupported media type response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query unsupported media type response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryTooManyRequests creates a PostQualityConversationsAuditsQueryTooManyRequests with default headers values
func NewPostQualityConversationsAuditsQueryTooManyRequests() *PostQualityConversationsAuditsQueryTooManyRequests {
	return &PostQualityConversationsAuditsQueryTooManyRequests{}
}

/*
PostQualityConversationsAuditsQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityConversationsAuditsQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query too many requests response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query too many requests response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query too many requests response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality conversations audits query too many requests response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality conversations audits query too many requests response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostQualityConversationsAuditsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryInternalServerError creates a PostQualityConversationsAuditsQueryInternalServerError with default headers values
func NewPostQualityConversationsAuditsQueryInternalServerError() *PostQualityConversationsAuditsQueryInternalServerError {
	return &PostQualityConversationsAuditsQueryInternalServerError{}
}

/*
PostQualityConversationsAuditsQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityConversationsAuditsQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query internal server error response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query internal server error response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query internal server error response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality conversations audits query internal server error response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality conversations audits query internal server error response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostQualityConversationsAuditsQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryServiceUnavailable creates a PostQualityConversationsAuditsQueryServiceUnavailable with default headers values
func NewPostQualityConversationsAuditsQueryServiceUnavailable() *PostQualityConversationsAuditsQueryServiceUnavailable {
	return &PostQualityConversationsAuditsQueryServiceUnavailable{}
}

/*
PostQualityConversationsAuditsQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityConversationsAuditsQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query service unavailable response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query service unavailable response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query service unavailable response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality conversations audits query service unavailable response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality conversations audits query service unavailable response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostQualityConversationsAuditsQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityConversationsAuditsQueryGatewayTimeout creates a PostQualityConversationsAuditsQueryGatewayTimeout with default headers values
func NewPostQualityConversationsAuditsQueryGatewayTimeout() *PostQualityConversationsAuditsQueryGatewayTimeout {
	return &PostQualityConversationsAuditsQueryGatewayTimeout{}
}

/*
PostQualityConversationsAuditsQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostQualityConversationsAuditsQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality conversations audits query gateway timeout response has a 2xx status code
func (o *PostQualityConversationsAuditsQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality conversations audits query gateway timeout response has a 3xx status code
func (o *PostQualityConversationsAuditsQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality conversations audits query gateway timeout response has a 4xx status code
func (o *PostQualityConversationsAuditsQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality conversations audits query gateway timeout response has a 5xx status code
func (o *PostQualityConversationsAuditsQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality conversations audits query gateway timeout response a status code equal to that given
func (o *PostQualityConversationsAuditsQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostQualityConversationsAuditsQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/conversations/audits/query][%d] postQualityConversationsAuditsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityConversationsAuditsQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityConversationsAuditsQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
