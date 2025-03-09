package models

// ResultCode : API 응답 코드 상수
type ResultCode int

const (
	// Success Codes (1000~1999)
	SUCCESS              ResultCode = 1000
	SUCCESS_WITH_MESSAGE ResultCode = 1001

	// Error Codes (2000~2999)
	ERROR_INVALID_INPUT    ResultCode = 2000
	ERROR_USER_NOT_FOUND   ResultCode = 2001
	ERROR_DUPLICATE_EMAIL  ResultCode = 2002
	ERROR_INVALID_PASSWORD ResultCode = 2003
	ERROR_INTERNAL_SERVER  ResultCode = 2004
	ERROR_UNAUTHORIZED     ResultCode = 2005
)

// ResultCodeMessages : 결과 코드에 대한 메시지 맵핑
var ResultCodeMessages = map[ResultCode]string{
	SUCCESS:              "Success",
	SUCCESS_WITH_MESSAGE: "Success with message",

	ERROR_INVALID_INPUT:    "Invalid input data",
	ERROR_USER_NOT_FOUND:   "User not found",
	ERROR_DUPLICATE_EMAIL:  "Email already exists",
	ERROR_INVALID_PASSWORD: "Invalid password",
	ERROR_INTERNAL_SERVER:  "Internal server error",
	ERROR_UNAUTHORIZED:     "Unauthorized access",
}

// APIResponse : 공통 API 응답 구조체
type APIResponse struct {
	Code    ResultCode  `json:"code"`           // 결과 코드
	Message string      `json:"message"`        // 결과 메시지
	Data    interface{} `json:"data,omitempty"` // 응답 데이터 (없으면 생략)
}

// NewSuccessResponse : 성공 응답 생성
func NewSuccessResponse(data interface{}) *APIResponse {
	return &APIResponse{
		Code:    SUCCESS,
		Message: ResultCodeMessages[SUCCESS],
		Data:    data,
	}
}

// NewErrorResponse : 에러 응답 생성
func NewErrorResponse(code ResultCode) *APIResponse {
	return &APIResponse{
		Code:    code,
		Message: ResultCodeMessages[code],
	}
}

// NewCustomResponse : 커스텀 메시지를 포함한 응답 생성
func NewCustomResponse(code ResultCode, message string, data interface{}) *APIResponse {
	return &APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
