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
}

func (p *Person) Valid() bool {
  return p.Name != "" && (p.Gender == MALE || p.Gender == FEMALE) &&
    p.Age > 0 && p.Height > 0 && p.Weight > 0 && p.HeartRate > 0 &&
    p.MaxHeartRate > 0
}

type Remark struct {
  Speed float32                 `bson:"speed" json:"speed"`
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
}

func (p *Player) Valid() bool {
  return p.Name != "" && p.DetailInfo.Valid() && p.History != ""
}

func (p1 *Player) Equals(p2 *Player) bool {
  return p1.Name == p2.Name &&
    p1.DetailInfo == p2.DetailInfo && p1.OverallRemark == p2.OverallRemark &&
    p1.History == p2.History
}

type TrainRecord struct {
  ObjId bson.ObjectId             `bson:"_id,omitempty" json:"_id"`
  Speed []float32                 `bson:"speed" json:"speed"`
  Distance []float32              `bson:"distance" json:"distance"`
  /**
  more display criterion
  */
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
  Latitude float32                `bson:"latitude" json:"latitude"`
  Longitude float32               `bson:"longitude" json:"longitude"`
  Height float32                  `bson:"height" json:"height"`
}

type ACCData struct {
  XAcc float32                    `bson:"xacc" json:"xacc"`
  YAcc float32                    `bson:"yacc" json:"yacc"`
  ZAcc float32                    `bson:"zacc" json:"zacc"`
}

type GYROData struct {
  XGyro float32                   `bson:"xgyro" json:"xgyro"`
  YGyro float32                   `bson:"ygyro" json:"ygyro"`
  ZGyro float32                   `bson:"zgyro" json:"zgyro"`
}

type RawTrainRecord struct {
  ObjId bson.ObjectId             `bson:"_id,omitempty" json:"_id"`
  GpsData []GPSData               `bson:"gpsdata" json:"gpsdata"`
  AccData []ACCData               `bson:"accdata" json:"accdata"`
  GyroData []GYROData             `bson:"gyrodata" json:"gyrodata"`
  HeartRateData []int             `bson:"heartratedata" json:"heartratedata"`
}

func (p1 *RawTrainRecord) Equals(p2 *RawTrainRecord) bool {
  if len(p1.GpsData) != len(p2.GpsData) ||
      len(p1.AccData) != len(p2.AccData) || len(p1.GyroData) != len(p2.GyroData) ||
      len(p1.HeartRateData) != len(p2.HeartRateData) {
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

  for i, _ := range p1.HeartRateData {
    if p1.HeartRateData[i] != p2.HeartRateData[i] {
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
