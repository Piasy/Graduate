package processor

import (
  "math"

  "models/types"
)

type NaiveProcessor struct {
}

func (processor NaiveProcessor) RawData2Record(raw *types.RawTrainRecord) *types.TrainRecord {
  record := types.TrainRecord {}

  record.HeartRate = raw.HeartRateData

  record.Speed = make([]float64, len(raw.GpsData))
  for i, _ := range raw.GpsData {
    record.Speed[i] = raw.GpsData[i].Speed
  }

  record.Distance = make([]float64, len(raw.GpsData))
  record.Distance[0] = 0
  for i := 1; i < len(record.Distance); i++ {
    record.Distance[i] = record.Distance[i - 1] +
      gps2dist(raw.GpsData[i].Latitude, raw.GpsData[i].Longitude,
        raw.GpsData[i - 1].Latitude, raw.GpsData[i - 1].Longitude)
  }
  return &record
}

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
