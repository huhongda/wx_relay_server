package consts

const (
	// etcd相关操作错误码
	ETCD_CREATE_DIR_ERROR = 3001
	ETCD_CREATE_KEY_ERROR = 3002
	ETCD_READ_KEY_ERROR   = 3003

	// http
	ERROR_CODE__SOURCE_DATA__ILLEGAL      = 1101 // 外部传入参数错误
	ERROR_CODE__GRPC__FAILED              = 1102 // grpc 调用失败
	ERROR_CODE__HTTP__CALL_FAILD_EXTERNAL = 1103 // http外部调用失败
	ERROR_CODE__HTTP__CALL_FAILD_INTERNAL = 1104 // http内部调用失败
	ERROR_CODE__JSON__PARSE_FAILED        = 1105 // JSON解析失败
	ERROR_CODE__HTTP__INPUT_EMPTY         = 1106
)
