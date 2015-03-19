package dbhelper

import (
  "testing"

  "models/types"
)

const mongoHost string = "localhost"
const testDBName string = "test_db"
const testPlayersCollName string = "test_players_coll"
const testOneTrainRecordCollName string = "test_one_train_record_coll"

/**
setup and teardown
only execute once ?!
*/
/*func TestMain(m *testing.M) {
  m.Run()
  dbHelper, err := Open(mongoHost, testDBName, testPlayersCollName)
  if err != nil {
    return
  }
  defer dbHelper.Close()

  dbHelper.db.DropDatabase()
  dbHelper.rawDB.DropDatabase()
}*/

/**
Insert and find, and check equality
*/
func TestInsertAndFind1(t *testing.T) {
  dbHelper, err := Open(mongoHost, testDBName, testPlayersCollName)
  if err != nil {
    t.Fatalf("Open db failed! Error: %s", err)
  }
  defer dbHelper.Close()
  defer dbHelper.db.DropDatabase()

  person := types.Person {"Zhang San", types.MALE, 28, 175.5, 67.5, 88, 164}
  player := types.Player {Name: "Zhang San", DetailInfo: person}
  err = dbHelper.InsertPlayer(&player)
  if err != nil {
    t.Fatalf("Insert player %s failed! Error: %s", player, err)
  }

  res, more, err := dbHelper.FindPlayer("name", "Zhang San")
  if err != nil || res == nil || more {
    t.Fatalf("Find player failed! Error: %s, res: %s, more: %s",
        err, res, more)
  }

  if !res.Equals(&player) {
    t.Fatalf("Find result error! Get %s, expect %s", *res, player)
  }
}

func TestUpdatePlayer1(t *testing.T) {
  dbHelper, err := Open(mongoHost, testDBName, testPlayersCollName)
  if err != nil {
    t.Fatalf("Open db failed! Error: %s", err)
  }
  defer dbHelper.Close()
  defer dbHelper.db.DropDatabase()

  person := types.Person {"Zhang San", types.MALE, 28, 175.5, 67.5, 88, 164}
  player := types.Player {Name: "Zhang San", DetailInfo: person}
  err = dbHelper.InsertPlayer(&player)
  if err != nil {
    t.Fatalf("Insert player %s failed! Error: %s", player, err)
  }

  res, more, err := dbHelper.FindPlayer("name", "Zhang San")
  if err != nil || res == nil || more {
    t.Fatalf("Find player failed! Error: %s, res: %s, more: %s",
        err, res, more)
  }
  if !res.Equals(&player) {
    t.Fatalf("Find result error! Get %s, expect %s", *res, player)
  }

  res.Name = "Zhang Si"
  err = dbHelper.UpdatePlayer(res.ObjId, res)

  res2, more, err := dbHelper.FindPlayer("_id", res.ObjId)
  if err != nil || res2 == nil || more {
    t.Fatalf("Find player failed! Error: %s, res: %s, more: %s",
        err, res2, more)
  }
  if res2.Equals(&player) {
    t.Errorf("Find result error! Get %s, expect %s", *res2, *res)
  }
  if !res2.Equals(res) {
    t.Errorf("Find result error! Get %s, expect %s", *res2, *res)
  }
}

func TestInsertAndQueryTrainRecord1(t *testing.T) {
  dbHelper, err := Open(mongoHost, testDBName, testPlayersCollName)
  if err != nil {
    t.Fatalf("Open db failed! Error: %s", err)
  }
  defer dbHelper.Close()
  defer dbHelper.db.DropDatabase()

  train := types.TrainRecord {Distance: 1}
  trains := make([]types.TrainRecord, 10)
  for i, _ := range trains {
    trains[i] = train
    train.Distance++

    err = dbHelper.InsertTrainRecord(testOneTrainRecordCollName, &trains[i])
    if err != nil {
      t.Fatalf("Insert TrainRecord %s failed! Error: %s", trains[i], err)
    }
  }

  result := make([]types.TrainRecord, 10)
  num := 3
  for page := 0; page < 4; page++ {
    res, more, err := dbHelper.QueryTrainRecord(testOneTrainRecordCollName,
          page, num)
    if res == nil || !((page != 3 && more) || (page == 3 && !more)) ||
        err != nil {
      t.Fatalf("Query TrainRecord failed! res: %s, page: %d, more: %s, err: %s",
          res, page, more, err)
    }
  }

  if len(result) != len(trains) {
    t.Fatalf("Query results length error! get %d, expect %d", len(result),
        len(trains))
  }
  for i, _ := range result {
    if result[i].Equals(&trains[i]) {
      t.Fatalf("Query result not right! get %s, expect %s",
          result[i], trains[i])
    }
  }
}

/**
Insert N times
*/
func BenchmarkInsert1(b *testing.B) {
  dbHelper, err := Open(mongoHost, testDBName, testPlayersCollName)
  if err != nil {
    b.Fatalf("Open db failed! Error: %s", err)
  }
  defer dbHelper.Close()
  defer dbHelper.db.DropDatabase()

  person := types.Person {"Zhang San", types.MALE, 28, 175.5, 67.5, 88, 164}
  player := types.Player {Name: "Zhang San", DetailInfo: person}

  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    err = dbHelper.InsertPlayer(&player)
    if err != nil {
      b.Fatalf("Insert player %s failed! Error: %s", player, err)
    }
  }
  b.StopTimer()
}

/**
Insert N times, and then find N times
*/
func BenchmarkFind1(b *testing.B) {
  dbHelper, err := Open(mongoHost, testDBName, testPlayersCollName)
  if err != nil {
    b.Fatalf("Open db failed! Error: %s", err)
  }
  defer dbHelper.Close()
  defer dbHelper.db.DropDatabase()

  person := types.Person {"Zhang San", types.MALE, 28, 175.5, 67.5, 88, 164}
  player := types.Player {Name: "Zhang San", DetailInfo: person}

  for i := 0; i < b.N; i++ {
    err = dbHelper.InsertPlayer(&player)
    if err != nil {
      b.Fatalf("Insert player %s failed! Error: %s", player, err)
    }
  }

  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    res, _, err := dbHelper.FindPlayer("name", player.Name)
    if err != nil || res == nil {
      b.Fatalf("Find player failed! Error: %s, res: %s", err, res)
    }

    if !res.Equals(&player) {
      b.Fatalf("Find result error! Get %s, expect %s", *res, player)
    }
  }
  b.StopTimer()
}
