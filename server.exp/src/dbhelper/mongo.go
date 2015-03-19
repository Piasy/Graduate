package dbhelper

import (
  "fmt"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "types"
)

const rawSuffix string = "_raw"

type DBHelper struct {
  session *mgo.Session
  db *mgo.Database
  players *mgo.Collection
  rawPlayers *mgo.Collection
}

func Open(url, dbName, playersCollName string) (*DBHelper, error) {
  dbHelper := &DBHelper{}
  var err error
  dbHelper.session, err = mgo.Dial(url)
  if err != nil {
    return nil, err
  }

  dbHelper.db = dbHelper.session.DB(dbName)
  dbHelper.players = dbHelper.db.C(playersCollName)
  dbHelper.rawPlayers = dbHelper.db.C(playersCollName + rawSuffix)
  return dbHelper, nil
}

func (db *DBHelper) Close() {
  defer db.session.Close()
}

/**
if use different database, the performance becomes very bad!!
243 / 1
*/
func (db *DBHelper) InsertPlayer(data *types.Player) error {
  //2000	    978292 ns/op
  //1000	   1128452 ns/op
  res, _, _ := db.FindPlayer("id", data.ID)
  if res != nil {
    return fmt.Errorf("Player %s exist!", *data)
  }

  //5000	    261272 ns/op
  //5000	    311218 ns/op
  err := db.players.Insert(data)
  if err != nil {
    return err
  }

  //1000	   1168428 ns/op
  //1000	   1123258 ns/op
  res, _, _ = db.findRawPlayer("id", data.ID)
  if res != nil {
    return fmt.Errorf("Player %s exist!", *data)
  }

  //3	 396848376 ns/op
  //1000	   1502766 ns/op
  err = db.rawPlayers.Insert(data)
  if err != nil {
    return err
  }

  return nil
}

func (db *DBHelper) FindPlayer(key string, value interface {}) (*types.Player,
          bool, error) {
  query := db.players.Find(bson.M{key: value})
  num, err := query.Count()
  if err != nil {
    return nil, false, err
  }

  if num < 1 {
    return nil, false, nil
  }

  result := &types.Player {}
  err = query.One(result)
  if err != nil {
    return nil, false, err
  }

  return result, num > 1, nil
}

func (db *DBHelper) UpdatePlayer(id bson.ObjectId, player *types.Player) error {
  return db.players.UpdateId(id, player)
}

func (db *DBHelper) InsertTrainRecord(coll string, data *types.TrainRecord) error {
  return db.db.C(coll).Insert(data)
}

func (db *DBHelper) QueryTrainRecord(coll string, page, num int) (
      []types.TrainRecord, bool, error) {
  query := findAll(db.db.C(coll))
  count, err := query.Count()
  if err != nil {
    return nil, false, err
  }

  if count <= page * num {
    return nil, false, nil
  }
  query = query.Skip(page * num)

  length := min(count - page * num, num)
  result := make([]types.TrainRecord, length)
  iter := query.Iter()
  index := 0

  for {
    next := iter.Next(&result[index])
    if next {
      index++
      if length <= index {
        break
      }
    } else if iter.Err() != nil {
      return result, false, err
    } else {
      break
    }
  }

  return result, num < count - page * num, nil
}

func (db *DBHelper) QueryRawTrainRecord(coll string, page, num int) (
      []types.RawTrainRecord, bool, error) {
  query := findAll(db.db.C(coll + rawSuffix))
  count, err := query.Count()
  if err != nil {
    return nil, false, err
  }

  if count <= page * num {
    return nil, false, nil
  }
  query = query.Skip(page * num)

  length := min(count - page * num, num)
  result := make([]types.RawTrainRecord, length)
  iter := query.Iter()
  index := 0

  for {
    next := iter.Next(&result[index])
    if next {
      index++
      if length <= index {
        break
      }
    } else if iter.Err() != nil {
      return result, false, err
    } else {
      break
    }
  }

  return result, num < count - page * num, nil
}

func (db *DBHelper) InsertRawTrainRecord(coll string,
        data *types.RawTrainRecord) error {
  return db.db.C(coll + rawSuffix).Insert(data)
}

func (db *DBHelper) findRawPlayer(key string, value interface {}) (*types.Player,
          bool, error) {
  query := db.rawPlayers.Find(bson.M{key: value})
  num, err := query.Count()
  if err != nil {
    return nil, false, err
  }

  if num < 1 {
    return nil, false, nil
  }

  result := &types.Player {}
  err = query.One(result)
  if err != nil {
    return nil, false, err
  }

  return result, num > 1, nil
}

func  hasColl(db *mgo.Database, coll string) (bool, error) {
  colls, err := db.CollectionNames()
  if err != nil {
    return false, err
  }

  for _, name := range colls {
    if name == coll {
      return true, nil
    }
  }

  return false, nil
}

func findAll(coll *mgo.Collection) *mgo.Query {
  return coll.Find(bson.M{})
}

func min(a, b int) int {
  if a < b {
    return a
  }

  return b
}
