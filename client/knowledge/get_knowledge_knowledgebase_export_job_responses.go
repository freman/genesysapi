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

// GetKnowledgeKnowledgebaseExportJobReader is a Reader for the GetKnowledgeKnowledgebaseExportJob structure.
type GetKnowledgeKnowledgebaseExportJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseExportJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseExportJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetKnowledgeKnowledgebaseExportJobAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseExportJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseExportJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseExportJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseExportJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseExportJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseExportJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseExportJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseExportJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseExportJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseExportJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseExportJobOK creates a GetKnowledgeKnowledgebaseExportJobOK with default headers values
func NewGetKnowledgeKnowledgebaseExportJobOK() *GetKnowledgeKnowledgebaseExportJobOK {
	return &GetKnowledgeKnowledgebaseExportJobOK{}
}

/*
GetKnowledgeKnowledgebaseExportJobOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseExportJobOK struct {
	Payload *models.KnowledgeExportJobResponse
}

// IsSuccess returns true when this get knowledge knowledgebase export job o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase export job o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase export job o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseExportJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobOK) GetPayload() *models.KnowledgeExportJobResponse {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeExportJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobAccepted creates a GetKnowledgeKnowledgebaseExportJobAccepted with default headers values
func NewGetKnowledgeKnowledgebaseExportJobAccepted() *GetKnowledgeKnowledgebaseExportJobAccepted {
	return &GetKnowledgeKnowledgebaseExportJobAccepted{}
}

/*
GetKnowledgeKnowledgebaseExportJobAccepted describes a response with status code 202, with default header values.

Export job in progress
*/
type GetKnowledgeKnowledgebaseExportJobAccepted struct {
	Payload *models.KnowledgeExportJobResponse
}

// IsSuccess returns true when this get knowledge knowledgebase export job accepted response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase export job accepted response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job accepted response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase export job accepted response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job accepted response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GetKnowledgeKnowledgebaseExportJobAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobAccepted  %+v", 202, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobAccepted) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobAccepted  %+v", 202, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobAccepted) GetPayload() *models.KnowledgeExportJobResponse {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeExportJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobBadRequest creates a GetKnowledgeKnowledgebaseExportJobBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseExportJobBadRequest() *GetKnowledgeKnowledgebaseExportJobBadRequest {
	return &GetKnowledgeKnowledgebaseExportJobBadRequest{}
}

/*
GetKnowledgeKnowledgebaseExportJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseExportJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobUnauthorized creates a GetKnowledgeKnowledgebaseExportJobUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseExportJobUnauthorized() *GetKnowledgeKnowledgebaseExportJobUnauthorized {
	return &GetKnowledgeKnowledgebaseExportJobUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseExportJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseExportJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobForbidden creates a GetKnowledgeKnowledgebaseExportJobForbidden with default headers values
func NewGetKnowledgeKnowledgebaseExportJobForbidden() *GetKnowledgeKnowledgebaseExportJobForbidden {
	return &GetKnowledgeKnowledgebaseExportJobForbidden{}
}

/*
GetKnowledgeKnowledgebaseExportJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseExportJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseExportJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobNotFound creates a GetKnowledgeKnowledgebaseExportJobNotFound with default headers values
func NewGetKnowledgeKnowledgebaseExportJobNotFound() *GetKnowledgeKnowledgebaseExportJobNotFound {
	return &GetKnowledgeKnowledgebaseExportJobNotFound{}
}

/*
GetKnowledgeKnowledgebaseExportJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseExportJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseExportJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobRequestTimeout creates a GetKnowledgeKnowledgebaseExportJobRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseExportJobRequestTimeout() *GetKnowledgeKnowledgebaseExportJobRequestTimeout {
	return &GetKnowledgeKnowledgebaseExportJobRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseExportJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseExportJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge() *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobUnsupportedMediaType creates a GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseExportJobUnsupportedMediaType() *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobTooManyRequests creates a GetKnowledgeKnowledgebaseExportJobTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseExportJobTooManyRequests() *GetKnowledgeKnowledgebaseExportJobTooManyRequests {
	return &GetKnowledgeKnowledgebaseExportJobTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseExportJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseExportJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase export job too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase export job too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobInternalServerError creates a GetKnowledgeKnowledgebaseExportJobInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseExportJobInternalServerError() *GetKnowledgeKnowledgebaseExportJobInternalServerError {
	return &GetKnowledgeKnowledgebaseExportJobInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseExportJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseExportJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase export job internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase export job internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobServiceUnavailable creates a GetKnowledgeKnowledgebaseExportJobServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseExportJobServiceUnavailable() *GetKnowledgeKnowledgebaseExportJobServiceUnavailable {
	return &GetKnowledgeKnowledgebaseExportJobServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseExportJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseExportJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase export job service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase export job service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseExportJobGatewayTimeout creates a GetKnowledgeKnowledgebaseExportJobGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseExportJobGatewayTimeout() *GetKnowledgeKnowledgebaseExportJobGatewayTimeout {
	return &GetKnowledgeKnowledgebaseExportJobGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseExportJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseExportJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase export job gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase export job gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase export job gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase export job gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase export job gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}][%d] getKnowledgeKnowledgebaseExportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseExportJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
