package controller

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"tugas4/models"
)

func TemanController(dashboard *models.Dashboard){
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') 

	var divisi string;
	fmt.Print("Masukkan divisi teman :")
	divisi, _ = reader.ReadString('\n')
	divisi = strings.TrimSpace(divisi)


	var nama string
	fmt.Print("Masukkan nama teman :")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	if dashboard.Teman == nil {
		dashboard.Teman = make(map[string][]string)
	}

	dashboard.Teman[divisi] = append(dashboard.Teman[divisi], nama)
	Dashboard(dashboard)
}