package masterdb
//package main
import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/mattn/go-sqlite3"
)
func main(){
  DbInit()
  //InsertNode("hogaaaaa", "fugdadsa")
  GetNodeInfo()
}
type Node struct {
  Id int
  Ip  string

}
// ノード追加時にIP、ホスト名を記録
// 追加成功時　返り値 true
func InsertNode(node Node)bool{
  db, err := sql.Open("sqlite3", "./test.db")
  if err != nil {
    panic(err)
  }
  defer db.Close()
  res, err := db.Exec(
    `INSERT INTO NODES (IP) VALUES (?)`,
    node.Ip,
  )
  if err != nil {
    return false
    //panic(err)
  }

  _, err = res.LastInsertId()
  if err != nil {
    log.Fatal("InsertNode", err)
    return false
    //panic(err)
  }
  return true
}
func UpdateNode(node Node)bool{
  fmt.Println(node)
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
      panic(err)
    }
    defer db.Close()
    res, err := db.Exec(
    `UPDATE NODES SET IP=? WHERE ID=?`,
    node.Ip,
    node.Id,
  )
  if err != nil {
    //return false
    panic(err)
  }

  // 更新されたレコード数
  affect, err := res.RowsAffected()
  if err != nil {
    //return false
    panic(err)
  }

  fmt.Println("affected by update: %d\n", affect)
  return true
}
// 全ノード情報取得
func GetNodeInfo()(map[int]Node){

  db, err := sql.Open("sqlite3", "./test.db")
  if err != nil {
    panic(err)
  }
  defer db.Close()

    rows, err := db.Query(
      `SELECT * FROM NODES`,
    )
    if err != nil {
      panic(err)
    }

    // 処理が終わったらカーソルを閉じる
    nodes := make(map[int]Node)
    defer rows.Close()
    for rows.Next() {
      var n Node

      // カーソルから値を取得
      if err := rows.Scan(&n.Id, &n.Ip); err != nil {
        log.Fatal("rows.Scan()", err)
        return  make(map[int]Node)
      }
      nodes[n.Id] = n

      //fmt.Printf("id: %d, ip: %s, hostname: %s\n", id, ip, hostname)
    }
    return nodes
}
func GetNodeInfoFromId(parm int)(Node){
  fmt.Printf("id %d",parm)
  var node  Node

  db, err := sql.Open("sqlite3", "./test.db")
  if err != nil {
    panic(err)
  }
  defer db.Close()

    rows, err := db.Query(
      `SELECT * FROM NODES WHERE ID = ?`,
      parm,
    )
    if err != nil {
      panic(err)
    }

    // 処理が終わったらカーソルを閉じる
    defer rows.Close()
    for rows.Next() {
      // カーソルから値を取得
      if err := rows.Scan(&node.Id, &node.Ip); err != nil {
        log.Fatal("rows.Scan()", err)
        return node
      }

    }

    return node
}

func DbInit(){
  // データベースのコネクションを開く
  db, err := sql.Open("sqlite3", "./test.db")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // テーブル作成
  _, err = db.Exec(
    `CREATE TABLE IF NOT EXISTS "NODES" ("ID" INTEGER PRIMARY KEY, "IP" VARCHAR(255) UNIQUE)`,
  )
  if err != nil {
    panic(err)
  }

}
