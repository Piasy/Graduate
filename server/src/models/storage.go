package models

import (
  "io"
  "time"
  "strconv"
  "math/rand"
  "crypto/md5"
  "encoding/hex"
  "sync"

  "github.com/astaxie/beego"
  "github.com/astaxie/beego/config"

  "models/types"
  "models/dbhelper"
  "models/processor"
)

const dbHost string = "localhost"
const dbName string = "GraduateDesignDB"
const playersCollName string = "PlayersColl"
var dbHelper *dbhelper.DBHelper
var username, password, token string
var rander *rand.Rand
var rawDataProcessor processor.Processor

//sync access
var players []types.Player
var rawTrainRecord map[string] types.RawTrainRecord
var processedRecord map[string] types.TrainRecord
var playersLock, rawLock, recordLock sync.RWMutex

func init() {
  //open db
  //the access of db needn't sync, it's done by mongo
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

  //these two maps should sync, ideally it needn't, cause every device access
  //its own k-v pair. but it can not be guaranteed
  //raw train record
  rawTrainRecord = make(map[string]types.RawTrainRecord)
  //processed record
  processedRecord = make(map[string]types.TrainRecord)
  //raw train record processor
  rawDataProcessor = processor.NaiveProcessor{}
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
  playersLock.RLock()
  defer playersLock.RUnlock()
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

  playersLock.Lock()
  players = append(players, *p)
  playersLock.Unlock()
  return nil
}

func UpdatePlayer(p *types.Player) error {
  err := dbHelper.UpdatePlayer(p.ObjId, p)
  if err != nil {
    return err
  }

  playersLock.Lock()
  for i, _ := range players {
    if players[i].ObjId == p.ObjId {
      players[i] = *p
      break
    }
  }
  playersLock.Unlock()
  return nil
}

func AppendRawTrainData(coll string, data *types.RawTrainRecord) {
  //when no key-value pair, tmpRaw will be an empty RawTrainRecord
  rawLock.Lock()
  tmpRaw := rawTrainRecord[coll]

  //when tmpRaw is an empty RawTrainRecord, its GpsData field is nil;
  //but this usage of append is legal
  tmpRaw.GpsData = append(tmpRaw.GpsData, data.GpsData...)
  tmpRaw.AccData = append(tmpRaw.AccData, data.AccData...)
  tmpRaw.GyroData = append(tmpRaw.GyroData, data.GyroData...)
  tmpRaw.HRData = append(tmpRaw.HRData, data.HRData...)
  rawTrainRecord[coll] = tmpRaw
  rawLock.Unlock()

  //calculate record
  recordLock.Lock()
  tmp := processedRecord[coll]
  processed := rawDataProcessor.RawData2Record(data)

  tmp.TimeStamp = append(tmp.TimeStamp, processed.TimeStamp...)

  tmp.Speed = append(tmp.Speed, processed.Speed...)
  tmp.CurSpeed = tmp.Speed[len(tmp.Speed) - 1]
  tmp.HIRun.Interval = append(tmp.HIRun.Interval, processed.HIRun.Interval...)
  if processed.HIRun.Times > 0 {
    tmp.HIRun.AveInterval = (tmp.HIRun.AveInterval * float32(tmp.HIRun.Times) +
      processed.HIRun.AveInterval * float32(processed.HIRun.Times)) /
      float32(tmp.HIRun.Times + processed.HIRun.Times)
    tmp.HIRun.Times += processed.HIRun.Times
  }

  tmp.Distance = append(tmp.Distance, processed.Distance...)
  tmp.CurDistance = tmp.Distance[len(tmp.Distance) - 1]
  if tmp.DistWithSpeed == nil {
    tmp.DistWithSpeed = make([]float64, len(processed.DistWithSpeed))
  }
  for i, _ := range tmp.DistWithSpeed {
    tmp.DistWithSpeed[i] += processed.DistWithSpeed[i]
  }
  if tmp.DistWithHR == nil {
    tmp.DistWithHR = make([]float64, len(processed.DistWithHR))
  }
  for i, _ := range tmp.DistWithHR {
    tmp.DistWithHR[i] += processed.DistWithHR[i]
  }

  tmp.HeartRate = append(tmp.HeartRate, processed.HeartRate...)
  tmp.CurHeartRate = tmp.HeartRate[len(tmp.HeartRate) - 1]
  if tmp.HeartRateElapse == nil {
    tmp.HeartRateElapse = make([]float32, len(processed.HeartRateElapse))
  }
  for i, _ := range tmp.HeartRateElapse {
    tmp.HeartRateElapse[i] += processed.HeartRateElapse[i]
  }
  tmp.MergeHRWithSpeed(processed)

  tmp.Position = append(tmp.Position, processed.Position...)

  processedRecord[coll] = tmp
  recordLock.Unlock()
}

func GetTrainRecord(coll string) *types.TrainRecord {
  recordLock.RLock()
  data := processedRecord[coll]
  recordLock.RUnlock()

  return &data
}

func FlushRawTrainData() error {
  //raw train record
  rawLock.Lock()
  for coll, data := range rawTrainRecord {
    err := dbHelper.InsertRawTrainRecord(coll, &data)
    if err != nil {
      return err
    }
  }

  for coll, _ := range rawTrainRecord {
    delete(rawTrainRecord, coll)
  }
  rawLock.Unlock()

  //train record
  recordLock.Lock()
  for coll, data := range processedRecord {
    err := dbHelper.InsertTrainRecord(coll, &data)
    if err != nil {
      return err
    }
  }

  //TODO when to clear this map?
  /*for coll, _ := range processedRecord {
    delete(processedRecord, coll)
  }*/
  recordLock.Unlock()
  return nil
}

func CloseDB() {
  if dbHelper != nil {
    dbHelper.Close()
  } else {
    beego.Info("DBHelper is nil!")
  }
}
