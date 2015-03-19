package models

import (
  "io"
  "time"
  "strconv"
  "math/rand"
  "crypto/md5"
  "encoding/hex"

  "github.com/astaxie/beego"
  "github.com/astaxie/beego/config"

  "models/types"
  "models/dbhelper"
)

const dbHost string = "localhost"
const dbName string = "GraduateDesignDB"
const playersCollName string = "PlayersColl"
var players []types.Player
var dbHelper *dbhelper.DBHelper
var username, password, token string
var rander *rand.Rand

func init() {
  var err error
  dbHelper, err = dbhelper.Open(dbHost, dbName, playersCollName)
  if err != nil {
    panic(err)
  }

  players, err = dbHelper.QueryAllPlayer()
  if err != nil {
    if players == nil {
      panic(err)
    } //else
    beego.Warn("Error happened when query all player: ", err)
  }

  if players == nil {
    players = make([]types.Player, 0)
  }

  appConf, err := config.NewConfig("ini", "conf/app.conf")
  if err != nil {
    panic(err)
  }

  username = appConf.String("admin::username")
  password = appConf.String("admin::password")
  rander = rand.New(rand.NewSource(time.Now().UnixNano()))
  ResetToken()
}

func VerifyAdmin(user, pass string) bool {
  return user == username && pass == password
}

func GetToken() string {
  return token
}

func ResetToken() {
  h := md5.New()
  io.WriteString(h, strconv.Itoa(rander.Int()))
  token = hex.EncodeToString(h.Sum(nil))
  beego.Info("new token: ", token)
}

func QueryPlayer(page, num int) *types.QueryResult {
  var next int
  var result []types.Player
  if 0 <= page * num && page * num < len(players) {
    if (page + 1) * num <= len(players) {
      next = page + 1
      result = players[page * num : (page + 1) * num]
    } else {
      next = -1
      result = players[page * num :]
    }
  } else {
    result = make([]types.Player, 0)
  }

  return &types.QueryResult {result, len(players), page - 1, page, next}
}

func InsertPlayer(p *types.Player) error {
  err := dbHelper.InsertPlayer(p)
  if err != nil {
    return err
  }

  players = append(players, *p)
  return nil
}

func UpdatePlayer(p *types.Player) error {
  err := dbHelper.UpdatePlayer(p.ObjId, p)
  if err != nil {
    return err
  }

  for i, _ := range players {
    if players[i].ObjId == p.ObjId {
      players[i] = *p
      break
    }
  }
  return nil
}

func CloseDB() {
  if dbHelper != nil {
    dbHelper.Close()
  } else {
    beego.Info("DBHelper is nil!")
  }
}
