package models

import (
  "math"

  "models/types"
)

func RawData2Remark(raw *types.RawTrainRecord) *types.Remark {

}

func omega2deltaPhi(oemga1, omega2 float32, deltaT int64) float32 {
  return (omega1 + omega2) / 2 * deltaT
}
