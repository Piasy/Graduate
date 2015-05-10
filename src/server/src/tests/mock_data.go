package test

import (
  "controllers/errors"
  "models/types"
)

var mocked_auth_err errors.Error = errors.Error {"You need login first", "20007", ""}

var mocked_auth_info types.AuthInfo = types.AuthInfo {"admin", "admin"}

var mocked_train_post_data_str string = "{\"op\": \"append\", \"player\": \"5508f6b522c3c263c5dfe6cf\", \"data\": {\"gpsdata\":[{\"latitude\":1,\"longitude\":2,\"height\":3},{\"latitude\":1,\"longitude\":2,\"height\":3}],\"accdata\":[{\"xacc\":1,\"yacc\":2,\"zacc\":3},{\"xacc\":1,\"yacc\":2,\"zacc\":3}],\"gyrodata\":[{\"xgyro\":1,\"ygyro\":2,\"zgyro\":3},{\"xgyro\":1,\"ygyro\":2,\"zgyro\":3}],\"heartratedata\":[1,2,3]}}"

var mocked_player types.Player = types.Player {Name: "Zhang San", DetailInfo: types.Person {"Zhang San", types.MALE, 28, 175.5, 67.5, 88, 164}, OverallRemark: types.Remark {5.5, 88}, History: "5508f6b522c3c263c5dfe6cf"}
