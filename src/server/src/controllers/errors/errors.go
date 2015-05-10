package errors

/**
ErrorCode definition:
20501
2: service error (1 is system error);
05: module code
01: detail error

module code:
player        01
trainrecord   02
*/
type Error struct {
  ErrorMsg string     `json:"error"`
  ErrorCode string    `json:"error_code"`
  Request string      `json:"request"`
}

func Issue(msg, code, req string) *Error {
  return &Error {msg, code, req}
}

const E_TYPE_SYSTEM string =      "1"
const E_TYPE_SERVICE string =     "2"

const E_MODULE_ANY string =           "00"
const E_MODULE_PLAYER string =        "01"
const E_MODULE_TRAINRECORD string =   "02"
const E_MODULE_TRAINREMARK string =   "03"
const E_MODULE_TRAINHISTORY string =  "04"

const E_DETAIL_INTERNAL_DB_ERROR string =     "01"
const E_DETAIL_PARSE_PARAM_ERROR string =     "02"
const E_DETAIL_NEED_LESS_PARAM string =       "03"
const E_DETAIL_NEED_MORE_PARAM string =       "04"
const E_DETAIL_ILLEGAL_PARAM string =         "05"
const E_DETAIL_INCOMPLETE_DATA string =       "06"
const E_DETAIL_NEED_LOGIN string =            "07"
