package models

import "fmt"

type Friend struct {
    Name   string
    Divisi string
}

var friends []Friend

func AddFriend() {
    var name, division string
    fmt.Print("Masukkan nama teman: ")
    fmt.Scan(&name)
    fmt.Print("Masukkan divisi teman: ")
    fmt.Scan(&division)

    friends = append(friends, Friend{Name: name, Divisi: division})
    fmt.Println("Teman berhasil ditambahkan.")
}
