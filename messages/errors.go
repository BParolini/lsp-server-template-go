package messages

type ErrorCode int

const (
	ParseError     ErrorCode = -32_700
	InvalidRequest ErrorCode = -32_600
	MethodNotFound ErrorCode = -32_601
	InvalidParams  ErrorCode = -32_602
	InternalError  ErrorCode = -32_603

	jsonrpcReservedErrorRangeStart ErrorCode = -32_099
	serverErrorStart               ErrorCode = jsonrpcReservedErrorRangeStart

	ServerNotInitialized ErrorCode = -32_002
	UnknownErrorCode     ErrorCode = -32_001

	jsonrpcReservedErrorRangeEnd ErrorCode = -32_000
	serverErrorEnd               ErrorCode = jsonrpcReservedErrorRangeEnd

	lspReservedErrorRangeStart ErrorCode = -32899

	RequestFailed    ErrorCode = -32803
	ServerCancelled  ErrorCode = -32802
	ContentModified  ErrorCode = -32801
	RequestCancelled ErrorCode = -32800

	lspReservedErrorRangeEnd ErrorCode = -32800
)
