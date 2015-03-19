package main

import (
  "fmt"

  "dbhelper"
  "types"
)

func main() {
  person := types.Person {"Zhang San", types.MALE, 28, 175.5, 67.5, 88, 164}
  player := types.Player {ID: 1, Name: "Zhang San", DetailInfo: person}
  dbHelper, err := dbhelper.Open("localhost", "GraduateTestDB", "Players")
  if err != nil {
    panic(err)
  }
  defer dbHelper.Close()

  dbHelper.InsertPlayer(&player)
  res, _, err := dbHelper.FindPlayer("name", "Zhang San")
  if err != nil {
    panic(err)
  }

  if res != nil {
    fmt.Println(*res)
  }

  dbHelper.QueryTrainRecord("ss", 0, 10)
}
