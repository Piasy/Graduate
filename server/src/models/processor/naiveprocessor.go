package processor

import (
  "math"

  "github.com/astaxie/beego"

  "models/types"
)

const SPEED_PIVOT_NORM float64 = 3.5
const SPEED_PIVOT_FAST float64 = 5
//const SPEED_PIVOT_HIGH_INTENSITY float64 =

type NaiveProcessor struct {
}

//TODO improve it
func (processor NaiveProcessor) RawData2Record(raw *types.RawTrainRecord) *types.TrainRecord {
  record := types.TrainRecord {}

  //TODO time sync better
  record.HeartRate = make([]int, len(raw.GpsData))
  for gpsIndex, hrIndex := 0, 0; gpsIndex < len(raw.GpsData) && hrIndex < len(raw.HRData); gpsIndex++ {
    for ; hrIndex < len(raw.HRData) && raw.HRData[hrIndex].Time < raw.GpsData[gpsIndex].Time; hrIndex++ {
    }
    if hrIndex < len(raw.HRData) {
      record.HeartRate[gpsIndex] = raw.HRData[hrIndex].HeartRate
    } else {
      record.HeartRate[gpsIndex] = raw.HRData[hrIndex - 1].HeartRate
    }
  }

  record.TimeStamp = make([]int64, len(raw.GpsData))
  record.Speed = make([]float64, len(raw.GpsData))
  for i, _ := range raw.GpsData {
    record.Speed[i] = raw.GpsData[i].Speed
    record.TimeStamp[i] = raw.GpsData[i].Time
  }

  record.Distance = make([]float64, len(raw.GpsData))
  record.Distance[0] = 0
  for i := 1; i < len(record.Distance); i++ {
    record.Distance[i] = record.Distance[i - 1] +
      gps2dist(raw.GpsData[i].Latitude, raw.GpsData[i].Longitude,
        raw.GpsData[i - 1].Latitude, raw.GpsData[i - 1].Longitude)
  }

  record.CurSpeed = record.Speed[len(record.Speed) - 1]
  record.CurDistance = record.Distance[len(record.Distance) - 1]
  record.CurHeartRate = record.HeartRate[len(record.HeartRate) - 1]

  record.DistWithSpeed = make([]float64, 3)
  for i := 1; i < len(record.Distance); i++ {
    if record.Speed[i] < SPEED_PIVOT_NORM {
      record.DistWithSpeed[0] += record.Distance[i] - record.Distance[i - 1]
    } else if record.Speed[i] < SPEED_PIVOT_FAST {
      record.DistWithSpeed[1] += record.Distance[i] - record.Distance[i - 1]
    } else {
      record.DistWithSpeed[2] += record.Distance[i] - record.Distance[i - 1]
    }
  }

  hiRun := false
  startIndex := 0
  for i, _ := range record.Speed {
    if record.Speed[i] >= SPEED_PIVOT_FAST {
      if !hiRun {
        record.HIRun.Times++
        startIndex = i
      }
      hiRun = true
    } else if hiRun {
      record.HIRun.Interval = append(record.HIRun.Interval,
        float32(raw.GpsData[i].Time - raw.GpsData[startIndex].Time) / 1000)
      hiRun = false
    }
  }
  //filter
  startIndex = 0
  for i, _ := range record.HIRun.Interval {
    if record.HIRun.Interval[i] > 5 {
      record.HIRun.Interval[startIndex] = record.HIRun.Interval[i]

      startIndex++
      record.HIRun.AveInterval += record.HIRun.Interval[i]
    } else {
      record.HIRun.Times--
    }
  }
  beego.Info(record.HIRun.Times, len(record.HIRun.Interval))
  if record.HIRun.Times > 0 {
    record.HIRun.AveInterval /= float32(record.HIRun.Times)
    if record.HIRun.Times <= len(record.HIRun.Interval) {
      record.HIRun.Interval = record.HIRun.Interval[: record.HIRun.Times]
    }
  } else {
    record.HIRun.Interval = record.HIRun.Interval[: 0]
  }

  //heart rate range time elapse
  //[0, 100), [100, 120), [120, 140), [140, 160), [160, 180), [180, 190), [190, 200), [200, )
  record.HeartRateElapse = make([]float32, 8)
  for i := 1; i < len(record.HeartRate); i++ {
    delta := float32(record.TimeStamp[i] - record.TimeStamp[i - 1]) / 1000
    if record.HeartRate[i] < 100 {
      record.HeartRateElapse[0] += delta
    } else if record.HeartRate[i] < 120 {
      record.HeartRateElapse[1] += delta
    } else if record.HeartRate[i] < 140 {
      record.HeartRateElapse[2] += delta
    } else if record.HeartRate[i] < 160 {
      record.HeartRateElapse[3] += delta
    } else if record.HeartRate[i] < 180 {
      record.HeartRateElapse[4] += delta
    } else if record.HeartRate[i] < 190 {
      record.HeartRateElapse[5] += delta
    } else if record.HeartRate[i] < 200 {
      record.HeartRateElapse[6] += delta
    } else {
      record.HeartRateElapse[7] += delta
    }
  }

  record.HRWithSpeed = make([]int, 3)
  count := make([]int, 3)
  for i := 1; i < len(record.HeartRate); i++ {
    if record.HeartRate[i] > 0 {
      if record.Speed[i] < SPEED_PIVOT_NORM {
        record.HRWithSpeed[0] += record.HeartRate[i]
        count[0]++
      } else if record.Speed[i] < SPEED_PIVOT_FAST {
        record.HRWithSpeed[1] += record.HeartRate[i]
        count[1]++
      } else {
        record.HRWithSpeed[2] += record.HeartRate[i]
        count[2]++
      }
    }
  }
  for i, _ := range count {
    if count[i] != 0 {
      record.HRWithSpeed[i] /= count[i]
    }
  }
  record.SetHRWithSpeedCount(count)

  record.DistWithHR = make([]float64, 3)
  for i := 1; i < len(record.Distance); i++ {
    if record.HeartRate[i] < 120 {
      record.DistWithHR[0] += record.Distance[i] - record.Distance[i - 1]
    } else if record.HeartRate[i] < 150 {
      record.DistWithHR[1] += record.Distance[i] - record.Distance[i - 1]
    } else {
      record.DistWithHR[2] += record.Distance[i] - record.Distance[i - 1]
    }
  }

  record.Position = raw.GpsData

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
