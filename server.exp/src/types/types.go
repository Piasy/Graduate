package types

import (
  "gopkg.in/mgo.v2/bson"
)

const (
  MALE = iota
  FEMALE
)

type Person struct {
  Name string
  Gender int
  Age int
  Height, Weight float32
  HeartRate, MaxHeartRate int
}

type Remark struct {
  Speed float32
  HeartRate int
  /**
  more remark criterion
  */
}

type Player struct {
  ObjId bson.ObjectId `bson:"_id,omitempty"`
  ID int
  Name string
  DetailInfo Person
  OverallRemark Remark
  History string
}

func (p1 *Player) Equals(p2 *Player) bool {
  return p1.ID == p2.ID && p1.Name == p2.Name &&
      p1.DetailInfo == p2.DetailInfo && p1.OverallRemark == p2.OverallRemark &&
      p1.History == p2.History
}

type TrainRecord struct {
  ObjId bson.ObjectId `bson:"_id,omitempty"`
  TrainID int
  Speed []float32
  Distance float32
  /**
  more display criterion
  */
}

func (p1 *TrainRecord) Equals(p2 *TrainRecord) bool {
  if p1.TrainID != p2.TrainID || len(p1.Speed) != len(p2.Speed) ||
      p1.Distance != p2.Distance {
    return false
  }

  for i, _ := range p1.Speed {
    if p1.Speed[i] != p2.Speed[i] {
      return false
    }
  }

  return true
}

/**
Raw train record
*/
type GPSData struct {
  Latitude, Longitude, Height float32
}

type ACCData struct {
  XAcc, YAcc, ZAcc float32
}

type GYROData struct {
  XGyro, YGyro, ZGyro float32
}

type RawTrainRecord struct {
  ObjId bson.ObjectId `bson:"_id,omitempty"`
  TrainID int
  GpsData []GPSData
  AccData []ACCData
  GyroData []GYROData
  HeartRateData []int
}

func (p1 *RawTrainRecord) Equals(p2 *RawTrainRecord) bool {
  if p1.TrainID != p2.TrainID || len(p1.GpsData) != len(p2.GpsData) ||
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
