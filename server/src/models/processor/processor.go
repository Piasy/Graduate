package processor

import (
  "models/types"
)

type Processor interface {
  RawData2Record(raw *types.RawTrainRecord) *types.TrainRecord
}
