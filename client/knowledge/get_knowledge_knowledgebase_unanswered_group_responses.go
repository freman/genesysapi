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

// GetKnowledgeKnowledgebaseUnansweredGroupReader is a Reader for the GetKnowledgeKnowledgebaseUnansweredGroup structure.
type GetKnowledgeKnowledgebaseUnansweredGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseUnansweredGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupOK creates a GetKnowledgeKnowledgebaseUnansweredGroupOK with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupOK() *GetKnowledgeKnowledgebaseUnansweredGroupOK {
	return &GetKnowledgeKnowledgebaseUnansweredGroupOK{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseUnansweredGroupOK struct {
	Payload *models.UnansweredGroup
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) GetPayload() *models.UnansweredGroup {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnansweredGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupBadRequest creates a GetKnowledgeKnowledgebaseUnansweredGroupBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupBadRequest() *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest {
	return &GetKnowledgeKnowledgebaseUnansweredGroupBadRequest{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupUnauthorized creates a GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupUnauthorized() *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized {
	return &GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupForbidden creates a GetKnowledgeKnowledgebaseUnansweredGroupForbidden with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupForbidden() *GetKnowledgeKnowledgebaseUnansweredGroupForbidden {
	return &GetKnowledgeKnowledgebaseUnansweredGroupForbidden{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupNotFound creates a GetKnowledgeKnowledgebaseUnansweredGroupNotFound with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupNotFound() *GetKnowledgeKnowledgebaseUnansweredGroupNotFound {
	return &GetKnowledgeKnowledgebaseUnansweredGroupNotFound{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout creates a GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout() *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout {
	return &GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge() *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType creates a GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType() *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests creates a GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests() *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests {
	return &GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase unanswered group too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupInternalServerError creates a GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupInternalServerError() *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError {
	return &GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase unanswered group internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable creates a GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable() *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable {
	return &GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase unanswered group service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout creates a GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout() *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout {
	return &GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase unanswered group gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase unanswered group gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase unanswered group gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase unanswered group gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase unanswered group gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}][%d] getKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseUnansweredGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
