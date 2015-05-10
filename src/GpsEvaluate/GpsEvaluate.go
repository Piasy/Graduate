package main

import (
  "fmt"
  "os"
  "bufio"
  "encoding/json"
  "math"
)

func gps2dist(lon1, lat1, lon2, lat2 float64) float64 {
  R := 6378.137 // Radius of earth in KM
  dLat := math.Abs(lat2 - lat1) * math.Pi / 180
  dLon := math.Abs(lon2 - lon1) * math.Pi / 180
  a := math.Sin(dLat/2) * math.Sin(dLat/2) +
          math.Cos(lat1 * math.Pi / 180) * math.Cos(lat2 * math.Pi / 180) *
                  math.Sin(dLon/2) * math.Sin(dLon/2)
  c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
  d := R * c
  return d * 1000
}

type GpsRecord struct {
  Time      int64     `json:"time"`
  Longitude float64   `json:"longitude"`
  Latitude  float64   `json:"latitude"`
  Altitude  float32   `json:"altitude"`
  Accuracy  float32   `json:"accuracy"`
  Speed     float64   `json:"speed"`
  Bearing   float32   `json:"bearing"`
}

func loadData(filename string) (data []GpsRecord) {
  file, err := os.Open(filename)
  if err != nil {
    return
  }
  defer file.Close()

  reader := bufio.NewReader(file)
  for {
    line, err := reader.ReadBytes('\n')
    if err != nil {
      return
    }
    record := GpsRecord{}
    err = json.Unmarshal(line, &record)
    if err == nil {
      data = append(data, record)
    }
  }
}

func calcDistence(data []GpsRecord) (dist float64) {
  for i := 0; i < len(data) - 1; i++ {
    dist += gps2dist(data[i].Longitude, data[i].Latitude, data[i + 1].Longitude, data[i + 1].Latitude)
  }
  return
}

func verifySpeedCalc(data []GpsRecord) {
  total, min, max := 0.0, 100.0, 0.0
  for i := 0; i < len(data) - 1; i++ {
    dist := gps2dist(data[i].Longitude, data[i].Latitude, data[i + 1].Longitude, data[i + 1].Latitude)
    calc := dist / (float64(data[i + 1].Time - data[i].Time) / 1000)
    diff := math.Abs((calc - data[i + 1].Speed) / data[i + 1].Speed)
    total += diff
    if diff < min {
      min = diff
    }
    if diff > max {
      max = diff
    }
  }
  fmt.Println(total / float64(len(data) - 1), min, max)
}

func main() {
  captures := []string {"location-capture-1427773201110.txt", "location-capture-1427773646529.txt", "location-capture-1427774683895.txt", "location-capture-1427774277595.txt"}

  data1 := loadData(captures[0])
  data2 := loadData(captures[1])
  data3 := loadData(captures[2])
  data4 := loadData(captures[3])
  fmt.Println(calcDistence(data1), calcDistence(data2), calcDistence(data3), calcDistence(data4))

  for _, record := range data4 {
    fmt.Println(record.Speed)
  }
}
