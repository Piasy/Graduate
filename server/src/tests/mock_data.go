package test

import (
  "controllers/errors"
  "models/types"
)

var mocked_auth_err errors.Error = errors.Error {"You need login first", "20007", ""}

var mocked_auth_info types.AuthInfo = types.AuthInfo {"admin", "admin"}
