package types

import (
  "gopkg.in/mgo.v2/bson"
)

const (
  MALE = iota
  FEMALE
)

type Person struct {
  Name string                   `bson:"name" json:"name"`
  Gender int                    `bson:"gender" json:"gender"`
  Age int                       `bson:"age" json:"age"`
  Height float32                `bson:"height" json:"height"`
  Weight float32                `bson:"weight" json:"weight"`
  HeartRate int                 `bson:"heartrate" json:"heartrate"`
  MaxHeartRate int              `bson:"maxheartrate" json:"maxheartrate"`
  Avatar string                 `bson:"avatar" json:"avatar"`
}

func (p *Person) Valid() bool {
  return p.Name != "" && (p.Gender == MALE || p.Gender == FEMALE) &&
    p.Age > 0 && p.Height > 0 && p.Weight > 0 && p.HeartRate > 0 &&
    p.MaxHeartRate > 0
}

type Remark struct {
  Speed float64                 `bson:"speed" json:"speed"`
  HeartRate int                 `bson:"heartrate" json:"heartrate"`
  /**
  more remark criterion
  */
}

type Player struct {
  ObjId bson.ObjectId            `bson:"_id,omitempty" json:"_id"`
  Name string                    `bson:"name" json:"name"`
  DetailInfo Person              `bson:"detailinfo" json:"detailinfo"`
  OverallRemark Remark           `bson:"overallremark" json:"overallremark"`
  History string                 `bson:"history" json:"history"`
  Position string                `bson:"position" json:"position"`
  DeviceID string                `bson:"deviceid" json:"deviceid"`
}

func (p *Player) Valid() bool {
  if p.ObjId.Hex() == "" {
    return p.Name != "" && p.DetailInfo.Valid()
  } else {
    return p.Name != "" && p.DetailInfo.Valid() && p.History != ""
  }
}

func (p1 *Player) Equals(p2 *Player) bool {
  return p1.Name == p2.Name &&
    p1.DetailInfo == p2.DetailInfo && p1.OverallRemark == p2.OverallRemark &&
    p1.History == p2.History
}

type TrainDesc struct {
  TimeStamp int64                 `bson:"timestamp" json:"timestamp"`
  Title string                    `bson:"title" json:"title"`
  Time string                     `bson:"time" json:"time"`
  Place string                    `bson:"place" json:"place"`
  Desc string                     `bson:"desc" json:"desc"`
}

type HighIntensityRun struct {
  Times int                       `bson:"times" json:"times"`
  Interval []float32              `bson:"interval" json:"interval"`
  AveInterval float32             `bson:"aveinterval" json:"aveinterval"`
}

type TrainRecord struct {
  ObjId bson.ObjectId             `bson:"_id,omitempty" json:"_id"`
  Desc TrainDesc                  `bson:"desc" json:"desc"`
  TimeStamp []int64               `bson:"timestamp" json:"timestamp"`

  Speed []float64                 `bson:"speed" json:"speed"`
  CurSpeed float64                `bson:"curspeed" json:"curspeed"`
  HIRun HighIntensityRun          `bson:"hirun" json:"hirun"`

  Distance []float64              `bson:"distance" json:"distance"`
  CurDistance float64             `bson:"curdistance" json:"curdistance"`
  DistWithSpeed []float64         `bson:"distwithspeed" json:"distwithspeed"`
  DistWithHR []float64            `bson:"distwithhr" json:"distwithhr"`

  HeartRate []int                 `bson:"heartrate" json:"heartrate"`
  CurHeartRate int                `bson:"curheartrate" json:"curheartrate"`
  HeartRateElapse []float32       `bson:"hrelapse" json:"hrelapse"`
  HRWithSpeed []int               `bson:"hrwithspeed" json:"hrwithspeed"`

  TrainTime int64                 `bsob:"traintime" json:"traintime"`
  MaxHeartRate int                `bsob:"maxheartrate" json:"maxheartrate"`
  AveHeartRate int                `bsob:"aveheartrate" json:"aveheartrate"`

  //private
  hrWithSpeedCount []int

  Position []GPSData              `bson:"position" json:"position"`
  /**
  more display criterion
  */
}

func (r *TrainRecord) SetHRWithSpeedCount(count []int) {
  r.hrWithSpeedCount = count
}

func (r1 *TrainRecord) MergeHRWithSpeed(r2 *TrainRecord) {
  //assume legal data
  if r1.HRWithSpeed == nil {
    r1.HRWithSpeed = make([]int, len(r2.HRWithSpeed))
  }
  if r1.hrWithSpeedCount == nil {
    r1.hrWithSpeedCount = make([]int, len(r2.hrWithSpeedCount))
  }
  for i, _ := range r1.HRWithSpeed {
    if r2.hrWithSpeedCount[i] > 0 {
      r1.HRWithSpeed[i] = (r1.HRWithSpeed[i] * r1.hrWithSpeedCount[i] +
        r2.HRWithSpeed[i] * r2.hrWithSpeedCount[i]) /
        (r1.hrWithSpeedCount[i] + r2.hrWithSpeedCount[i])
      r1.hrWithSpeedCount[i] += r2.hrWithSpeedCount[i]
    }
  }
}

func (p1 *TrainRecord) Equals(p2 *TrainRecord) bool {
  if len(p1.Speed) != len(p2.Speed) ||
      len(p1.Distance) != len(p2.Distance) {
    return false
  }

  for i, _ := range p1.Speed {
    if p1.Speed[i] != p2.Speed[i] {
      return false
    }
  }

  for i, _ := range p1.Distance {
    if p1.Distance[i] != p2.Distance[i] {
      return false
    }
  }

  return true
}

/**
Raw train record
*/
type GPSData struct {
  Latitude float64                `bson:"latitude" json:"latitude"`
  Longitude float64               `bson:"longitude" json:"longitude"`
  Altitude float64                `bson:"altitude" json:"altitude"`
  Bearing float64                 `bson:"bearing" json:"bearing"`
  Speed float64                   `bson:"speed" json:"speed"`
  Accuracy float64                `bson:"accuracy" json:"accuracy"`
  Time int64                      `bson:"time" json:"time"`
}

type ACCData struct {
  XAcc float64                    `bson:"xacc" json:"xacc"`
  YAcc float64                    `bson:"yacc" json:"yacc"`
  ZAcc float64                    `bson:"zacc" json:"zacc"`
  Time int64                      `bson:"time" json:"time"`
}

type GYROData struct {
  XGyro float64                   `bson:"xgyro" json:"xgyro"`
  YGyro float64                   `bson:"ygyro" json:"ygyro"`
  ZGyro float64                   `bson:"zgyro" json:"zgyro"`
  Time int64                      `bson:"time" json:"time"`
}

type HeartRateData struct {
  HeartRate int                   `bson:"heartrate" json:"heartrate"`
  Time int64                      `bson:"time" json:"time"`
}

type RawTrainRecord struct {
  ObjId bson.ObjectId             `bson:"_id,omitempty" json:"_id"`
  GpsData []GPSData               `bson:"gpsdata" json:"gpsdata"`
  AccData []ACCData               `bson:"accdata" json:"accdata"`
  GyroData []GYROData             `bson:"gyrodata" json:"gyrodata"`
  HRData []HeartRateData          `bson:"hrdata" json:"hrdata"`
}

func (p1 *RawTrainRecord) Equals(p2 *RawTrainRecord) bool {
  if len(p1.GpsData) != len(p2.GpsData) ||
      len(p1.AccData) != len(p2.AccData) || len(p1.GyroData) != len(p2.GyroData) ||
      len(p1.HRData) != len(p2.HRData) {
    return false
  }

  for i, _ := range p1.GpsData {
    if p1.GpsData[i] != p2.GpsData[i] {
      return false
    }
  }

  for i, _ := range p1.AccData {
    if p1.AccData[i] != p2.AccData[i] {
      return false
    }
  }

  for i, _ := range p1.GyroData {
    if p1.GyroData[i] != p2.GyroData[i] {
      return false
    }
  }

  for i, _ := range p1.HRData {
    if p1.HRData[i] != p2.HRData[i] {
      return false
    }
  }

  return true
}

/**
query result
*/
type QueryResult struct {
  Result interface{}  `json:"result"`
  TotalNum int        `json:"total_number"`
  Before int          `json:"before"`        //-1 means this is the first page
  Current int         `json:"current"`
  Next int            `json:"next"`          //-1 means this is the last page
}

type AuthInfo struct {
  Username string     `json:"username"`
  Password string     `json:"password"`
}

func (a *AuthInfo) Valid() bool {
  return a.Username != "" && a.Password != ""
}
