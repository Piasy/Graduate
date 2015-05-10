package dbhelper

import (
  "crypto/sha1"
  "fmt"
  "bytes"
  "encoding/gob"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

  "models/types"
  "utils"
)

const rawSuffix string = "_raw"

type DBHelper struct {
  session *mgo.Session
  db *mgo.Database
  players *mgo.Collection
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
  data.ObjId = bson.NewObjectId()
  w := new(bytes.Buffer)
  encoder := gob.NewEncoder(w)
  err := encoder.Encode(data)
  if err != nil {
    return err
  }
  data.History = fmt.Sprintf("%x", sha1.Sum(w.Bytes()))
  return db.players.Insert(data)
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

/**
no result: nil, nil
error happened: nil, err
partially result(has error): result, err
success: result, nil
*/
func (db *DBHelper) QueryAllPlayer() ([]types.Player, error) {
  query := findAll(db.players)
  count, err := query.Count()
  if err != nil {
    return nil, err
  }

  if count == 0 {
    return nil, nil
  }

  result := make([]types.Player, count)
  iter := query.Iter()
  index := 0
  for {
    next := iter.Next(&result[index])
    if next {
      index++
      if index == count {
        break
      }
    } else if iter.Err() != nil {
      return result, err
    } else {
      break
    }
  }

  return result, nil
}

func (db *DBHelper) UpdatePlayer(id bson.ObjectId,
      player *types.Player) error {
  return db.players.UpdateId(id, player)
}

func (db *DBHelper) InsertTrainRecord(coll string,
      data *types.TrainRecord) error {
  data.ObjId = bson.NewObjectId()
  return db.db.C(coll).Insert(data)
}

func (db *DBHelper) QueryTrainRecord(coll string, page, num int) (
      []types.TrainRecord, int, error) {
  query := findAll(db.db.C(coll))
  count, err := query.Count()
  if err != nil {
    return nil, 0, err
  }

  if count <= page * num {
    return nil, count, nil
  }
  query = query.Skip(page * num)

  length := utils.Min(count - page * num, num)
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
      return result, count, err
    } else {
      break
    }
  }

  return result, count, nil
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

  length := utils.Min(count - page * num, num)
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
  data.ObjId = bson.NewObjectId()
  return db.db.C(coll + rawSuffix).Insert(data)
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
