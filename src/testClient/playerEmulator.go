package main

import (
  "net/http"
  "time"
  "math/rand"
  "bytes"
  "encoding/json"
  "io/ioutil"
  "fmt"

  "gopkg.in/mgo.v2/bson"
)

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

type PostData struct {
  Op string                 `json:"op"`
  Player string             `json:"player"`
  Data RawTrainRecord       `json:"data"`
}
type Success struct {
  Succ bool                 `json:"success"`
}

func timemilis() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}

func randNextGps(last *GPSData, rander *rand.Rand) *GPSData {
  last.Time += 500 + int64(rander.Float32() * 1000)
  last.Latitude += (rander.Float64() * 2 - 1) * 0.0001
  last.Longitude += (rander.Float64() * 2 - 1) * 0.0001
  last.Speed += (rander.Float64() * 2 - 1) * 2
  return last
}

func randNextAcc(last *ACCData, rander *rand.Rand) *ACCData {
  return last
}

func randNextGyro(last *GYROData, rander *rand.Rand) *GYROData {
  return last
}

func randNextHR(last *HeartRateData, rander *rand.Rand) *HeartRateData {
  last.Time += 500 + int64(rander.Float32() * 1000)
  last.HeartRate += int((rander.Float32() * 2 - 1) * 10)
  return last
}

func emulate(player string, ch chan int) {
  lastGps := GPSData{40.00797556, 116.32371427, 30, 96.9000015258789, 5.2744019031524658, 3, 1427774768000}
  lastAcc := ACCData{1.4322, 2.322, 9.6372, 1427774768000}
  lastGyro := GYROData{1, 2, 3, 1427774768000}
  lastHR := HeartRateData{88, 1427774768000}
  rander := rand.New(rand.NewSource(time.Now().UnixNano()))
  for i := 0; i < 270; i++ {
    data := PostData{}
    data.Op = "append"
    data.Player = player
    data.Data = RawTrainRecord{}
    data.Data.GpsData = make([]GPSData, 10)
    data.Data.AccData = make([]ACCData, 10)
    data.Data.GyroData = make([]GYROData, 10)
    data.Data.HRData = make([]HeartRateData, 10)

    lastGps.Time = timemilis()
    lastAcc.Time = timemilis()
    lastGyro.Time = timemilis()
    lastHR.Time = timemilis()
    for j := 0; j < 10; j++ {
      data.Data.GpsData[j] = *randNextGps(&lastGps, rander)
      data.Data.AccData[j] = *randNextAcc(&lastAcc, rander)
      data.Data.GyroData[j] = *randNextGyro(&lastGyro, rander)
      data.Data.HRData[j] = *randNextHR(&lastHR, rander)
    }
    bb, err := json.Marshal(&data)
    if err == nil {
      buf := bytes.NewReader(bb)
      resp, err := http.Post("http://127.0.0.1:8080/api/rawtrainrecord", "application/json; charset=utf-8", buf)
      if err == nil {
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Post raw data for " + player + " failed! read resp error", err)
        } else {
          suc := Success{}
          err = json.Unmarshal(body, &suc)
          if err != nil || !suc.Succ {
            fmt.Println("Post raw data for " + player + " failed! parse succ error", err, suc)
          }
        }
      }
    }
    time.Sleep(10 * 1000 * time.Millisecond)
  }
  ch <- 1
}

func main() {
  // players := []string {"b564ab4a431ce4448969106d60e0c4e0ea85e76a",
  //                      "d25db944c93f7757264be315b43d64b0dd3ef3c5",
  //                      "18f6718c7a50cb45e4ce7989f48fd4c2641255f9",
  //                      "de33647785c3788a882e7006af395f5e9e1b9f84",
  //                      "64e61495e028e93faf6134b10d4bdf3a1ce77755",
  //                      "afc45b5c7b4cf463b6dbcf4378c31366aad69ed9",
  //                      "5a7f3fbe5a921ce6ebb936ca455016c306ddf060",
  //                      "7cd84e8b75e476261f418f64c5d3da2fb48ce242",
  //                      "ff1a83c49d65548d76ae0e9a2ccfc9848c699270",
  //                      "b8d893136c4b54463bd470318a342cefc8947704"}


  players := []string {"b564ab4a431ce4448969106d60e0c4e0ea85e76a",
                       "d25db944c93f7757264be315b43d64b0dd3ef3c5",
                       "18f6718c7a50cb45e4ce7989f48fd4c2641255f9"}
  ch := make(chan int)
  for _, p := range players {
    go emulate(p, ch)
    time.Sleep(2 * 1000 * time.Millisecond)
  }

  for range players {
    <- ch
  }
}
