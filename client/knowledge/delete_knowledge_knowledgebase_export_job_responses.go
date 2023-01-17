// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteKnowledgeKnowledgebaseExportJobReader is a Reader for the DeleteKnowledgeKnowledgebaseExportJob structure.
type DeleteKnowledgeKnowledgebaseExportJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseExportJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseExportJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseExportJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseExportJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseExportJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseExportJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteKnowledgeKnowledgebaseExportJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseExportJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseExportJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseExportJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseExportJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseExportJobNoContent creates a DeleteKnowledgeKnowledgebaseExportJobNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobNoContent() *DeleteKnowledgeKnowledgebaseExportJobNoContent {
	return &DeleteKnowledgeKnowledgebaseExportJobNoContent{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobNoContent describes a response with status code 204, with default header values.

Export job deleted
*/
type DeleteKnowledgeKnowledgebaseExportJobNoContent struct {
}

// IsSuccess returns true when this delete knowledge knowledgebase export job no content response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete knowledge knowledgebase export job no content response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job no content response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase export job no content response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job no content response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobBadRequest creates a DeleteKnowledgeKnowledgebaseExportJobBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobBadRequest() *DeleteKnowledgeKnowledgebaseExportJobBadRequest {
	return &DeleteKnowledgeKnowledgebaseExportJobBadRequest{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseExportJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job bad request response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job bad request response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job bad request response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job bad request response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job bad request response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobUnauthorized creates a DeleteKnowledgeKnowledgebaseExportJobUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobUnauthorized() *DeleteKnowledgeKnowledgebaseExportJobUnauthorized {
	return &DeleteKnowledgeKnowledgebaseExportJobUnauthorized{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseExportJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job unauthorized response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job unauthorized response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job unauthorized response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job unauthorized response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job unauthorized response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobForbidden creates a DeleteKnowledgeKnowledgebaseExportJobForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobForbidden() *DeleteKnowledgeKnowledgebaseExportJobForbidden {
	return &DeleteKnowledgeKnowledgebaseExportJobForbidden{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseExportJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job forbidden response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job forbidden response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job forbidden response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job forbidden response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job forbidden response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobNotFound creates a DeleteKnowledgeKnowledgebaseExportJobNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobNotFound() *DeleteKnowledgeKnowledgebaseExportJobNotFound {
	return &DeleteKnowledgeKnowledgebaseExportJobNotFound{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseExportJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job not found response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job not found response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job not found response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job not found response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job not found response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobRequestTimeout creates a DeleteKnowledgeKnowledgebaseExportJobRequestTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobRequestTimeout() *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout {
	return &DeleteKnowledgeKnowledgebaseExportJobRequestTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteKnowledgeKnowledgebaseExportJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job request timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job request timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job request timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job request timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job request timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job request entity too large response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job request entity too large response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job request entity too large response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job request entity too large response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job request entity too large response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job unsupported media type response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job unsupported media type response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job unsupported media type response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job unsupported media type response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job unsupported media type response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobTooManyRequests creates a DeleteKnowledgeKnowledgebaseExportJobTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobTooManyRequests() *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseExportJobTooManyRequests{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseExportJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job too many requests response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job too many requests response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job too many requests response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase export job too many requests response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase export job too many requests response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobInternalServerError creates a DeleteKnowledgeKnowledgebaseExportJobInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobInternalServerError() *DeleteKnowledgeKnowledgebaseExportJobInternalServerError {
	return &DeleteKnowledgeKnowledgebaseExportJobInternalServerError{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseExportJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job internal server error response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job internal server error response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job internal server error response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase export job internal server error response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase export job internal server error response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobServiceUnavailable creates a DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobServiceUnavailable() *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job service unavailable response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job service unavailable response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job service unavailable response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase export job service unavailable response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase export job service unavailable response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseExportJobGatewayTimeout creates a DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseExportJobGatewayTimeout() *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase export job gateway timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase export job gateway timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase export job gateway timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase export job gateway timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase export job gateway timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] deleteKnowledgeKnowledgebaseExportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseExportJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
