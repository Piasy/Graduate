package models

import (
  "io"
  "time"
  "strconv"
  "math/rand"
  "crypto/md5"
  "encoding/hex"

  "github.com/astaxie/beego"
  "github.com/astaxie/beego/config"

  "models/types"
  "models/dbhelper"
)

const dbHost string = "localhost"
const dbName string = "GraduateDesignDB"
const playersCollName string = "PlayersColl"
var players []types.Player
var dbHelper *dbhelper.DBHelper
var username, password, token string
var rander *rand.Rand
var rawTrainRecord map[string] types.RawTrainRecord
var processedRemark map[string] types.TrainRecord

func init() {
  //open db
  var err error
  dbHelper, err = dbhelper.Open(dbHost, dbName, playersCollName)
  if err != nil {
    panic(err)
  }

  //preload all players
  players, err = dbHelper.QueryAllPlayer()
  if err != nil {
    if players == nil {
      panic(err)
    } //else
    beego.Warn("Error happened when query all player: ", err)
  }

  if players == nil {
    players = make([]types.Player, 0)
  }

  //auth
  appConf, err := config.NewConfig("ini", "conf/app.conf")
  if err != nil {
    panic(err)
  }

  username = appConf.String("admin::username")
  password = appConf.String("admin::password")
  rander = rand.New(rand.NewSource(time.Now().UnixNano()))
  ResetToken()

  //raw train record
  rawTrainRecord = make(map[string]types.RawTrainRecord)
  //processed remark
  processedRemark = make(map[string]types.TrainRecord)
}

func VerifyAdmin(user, pass string) bool {
  return user == username && pass == password
}

func GetToken() string {
  return token
}

func ResetToken() {
  h := md5.New()
  io.WriteString(h, strconv.Itoa(rander.Int()))
  token = hex.EncodeToString(h.Sum(nil))
  beego.Info("new token: ", token)
}

func QueryPlayer(page, num int) *types.QueryResult {
  var next int
  var result []types.Player
  if 0 <= page * num && page * num < len(players) {
    if (page + 1) * num <= len(players) {
      next = page + 1
      result = players[page * num : (page + 1) * num]
    } else {
      next = -1
      result = players[page * num :]
    }
  } else {
    result = make([]types.Player, 0)
  }

  return &types.QueryResult {result, len(players), page - 1, page, next}
}


func QueryTrainRecord(coll string, page, num int) *types.QueryResult {
  var result []types.TrainRecord
  var next, count int
  var err error

  if page < 0 || num < 0 {
    result, count, err = dbHelper.QueryTrainRecord(coll, 0, 10)
    if err != nil {
      return &types.QueryResult {make([]types.Player, 0), 0, page - 1, page, -1}
    }
    if count < 10 {
      next = -1
    } else {
      next = page + 1
    }
    return &types.QueryResult {result, count, page - 1, page, next}
  }

  result, count, err = dbHelper.QueryTrainRecord(coll, page, num)
  if err != nil {
    return &types.QueryResult {make([]types.Player, 0), 0, page - 1, page, -1}
  }

  if len(result) < num {
    next = -1
  } else {
    next = page + 1
  }

  return &types.QueryResult {result, count, page - 1, page, next}
}

func InsertPlayer(p *types.Player) error {
  err := dbHelper.InsertPlayer(p)
  if err != nil {
    return err
  }

  players = append(players, *p)
  return nil
}

func UpdatePlayer(p *types.Player) error {
  err := dbHelper.UpdatePlayer(p.ObjId, p)
  if err != nil {
    return err
  }

  for i, _ := range players {
    if players[i].ObjId == p.ObjId {
      players[i] = *p
      break
    }
  }
  return nil
}

func AppendRawTrainData(coll string, data *types.RawTrainRecord) {
  //when no key-value pair, tmp will be an empty RawTrainRecord
  tmp := rawTrainRecord[coll]

  //when tmp is an empty RawTrainRecord, its GpsData field is nil;
  //but this usage of append is legal
  tmp.GpsData = append(rawTrainRecord[coll].GpsData, data.GpsData...)
  tmp.AccData = append(rawTrainRecord[coll].AccData, data.AccData...)
  tmp.GyroData = append(rawTrainRecord[coll].GyroData, data.GyroData...)
  tmp.HeartRateData = append(rawTrainRecord[coll].HeartRateData, data.HeartRateData...)
  rawTrainRecord[coll] = tmp

  //TODO calculate remarks
}

func GetTrainSpeed(coll string, page, num int) *types.QueryResult {
  data := processedRemark[coll]
  var next int
  var result []float32
  if 0 <= page * num && page * num < len(data.Speed) {
    if (page + 1) * num <= len(data.Speed) {
      next = page + 1
      result = data.Speed[page * num : (page + 1) * num]
    } else {
      next = -1
      result = data.Speed[page * num :]
    }
  } else {
    result = make([]float32, 0)
  }

  return &types.QueryResult {result, len(data.Speed), page - 1, page, next}
}

func GetTrainDistance(coll string, page, num int) *types.QueryResult {
  data := processedRemark[coll]
  var next int
  var result []float32
  if 0 <= page * num && page * num < len(data.Distance) {
    if (page + 1) * num <= len(data.Distance) {
      next = page + 1
      result = data.Distance[page * num : (page + 1) * num]
    } else {
      next = -1
      result = data.Distance[page * num :]
    }
  } else {
    result = make([]float32, 0)
  }

  return &types.QueryResult {result, len(data.Distance), page - 1, page, next}
}

func FlushRawTrainData() error {
  for coll, data := range rawTrainRecord {
    err := dbHelper.InsertRawTrainRecord(coll, &data)
    if err != nil {
      return err
    }
  }

  for coll, _ := range rawTrainRecord {
    delete(rawTrainRecord, coll)
  }
  return nil
}

func CloseDB() {
  if dbHelper != nil {
    dbHelper.Close()
  } else {
    beego.Info("DBHelper is nil!")
  }
}
