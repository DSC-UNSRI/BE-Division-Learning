package main

import "fmt"


type Pertandingan struct {
    timA, timB int
}

func simulasiPertandingan(saatIni int, skorSementara []int, daftarPertandingan []Pertandingan, totalPertandingan int, skorAkhir []int, jumlahTim int, hasil *bool) {
    if saatIni >= totalPertandingan {
        cocok := true
        for i := 0; i < jumlahTim; i++ {
            if skorSementara[i] != skorAkhir[i] {
                cocok = false
                break
            }
        }
        if cocok {
            *hasil = true
        }
        return
    }

    A, B := daftarPertandingan[saatIni].timA, daftarPertandingan[saatIni].timB

    skorBaru := make([]int, jumlahTim)
    copy(skorBaru, skorSementara)
    skorBaru[A] += 3
    simulasiPertandingan(saatIni+1, skorBaru, daftarPertandingan, totalPertandingan, skorAkhir, jumlahTim, hasil)
    if *hasil {
        return
    }

    copy(skorBaru, skorSementara)
    skorBaru[B] += 3
    simulasiPertandingan(saatIni+1, skorBaru, daftarPertandingan, totalPertandingan, skorAkhir, jumlahTim, hasil)
    if *hasil {
        return
    }

    copy(skorBaru, skorSementara)
    skorBaru[A]++
    skorBaru[B]++
    simulasiPertandingan(saatIni+1, skorBaru, daftarPertandingan, totalPertandingan, skorAkhir, jumlahTim, hasil)
    if *hasil {
        return
    }
}

func main() {
    var jumlahUji int
    fmt.Scan(&jumlahUji)

    for t := 0; t < jumlahUji; t++ {
        var jumlahTim int
        fmt.Scan(&jumlahTim)

        skorAkhir := make([]int, jumlahTim)
        for i := 0; i < jumlahTim; i++ {
            fmt.Scan(&skorAkhir[i])
        }

        totalPertandingan := (jumlahTim * (jumlahTim - 1)) / 2
        daftarPertandingan := make([]Pertandingan, 0, totalPertandingan)

        for i := 0; i < jumlahTim; i++ {
            for j := i + 1; j < jumlahTim; j++ {
                daftarPertandingan = append(daftarPertandingan, Pertandingan{i, j})
            }
        }

        skorSaatIni := make([]int, jumlahTim)
        kemungkinan := false
        simulasiPertandingan(0, skorSaatIni, daftarPertandingan, totalPertandingan, skorAkhir, jumlahTim, &kemungkinan)

        if kemungkinan {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
