package controllers

import (
	"bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

	"tugas04/models"
)

func DisplayMenu() int {
	var choice int
	fmt.Println("\n=== Dashboard Iftar GDCoC ===")
	fmt.Println("1. Pilih Opsi Kendaraan")
	fmt.Println("2. Input Barang yang Dibawa")
	fmt.Println("3. Tambah Rekomendasi")
	fmt.Println("4. Data Teman yang Ikut")
	fmt.Println("5. Lihat Semua Data")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu (1-6): ")
	
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Input tidak valid, coba lagi.")
		var discard string
		fmt.Scanln(&discard)
		return DisplayMenu()
	}
	
	if choice < 1 || choice > 6 {
		fmt.Println("Pilihan tidak valid, coba lagi.")
		return DisplayMenu()
	}
	
	return choice
}

func ManageTransportation(dash *models.Dashboard) {
	for {
		fmt.Println("\n=== Pilih Kendaraan ===")
		fmt.Println("Kendaraan yang tersedia:")
		fmt.Println("1. Kendaraan Pribadi")
		fmt.Println("2. Bus Keliling")
		fmt.Println("3. Nebeng")
		fmt.Println("4. Kembali ke Menu Utama")
		
		var choice int
		fmt.Print("Pilih opsi (1-4): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Input tidak valid, coba lagi.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		
		if choice == 4 {
			break
		}
		
		var transType string
		switch choice {
		case 1:
			transType = "Kendaraan Pribadi"
		case 2:
			transType = "Bus Keliling"
		case 3:
			transType = "Nebeng"
		case 4: 
			transType = "Travel"
		}
		
		if len(dash.Transportation) >= 1 {
			fmt.Println("Anda sudah memilih 1 kendaraan. Hapus salah satu untuk menambah yang baru.")
			fmt.Println("Kendaraan saat ini:")
			for i, t := range dash.Transportation {
				fmt.Printf("%d. %s\n", i+1, t.Type)
			}
			
			var removeChoice int
			fmt.Print("Pilih kendaraan untuk dihapus (1-" + strconv.Itoa(len(dash.Transportation)) + ") atau 0 untuk batal: ")
			fmt.Scanln(&removeChoice)
			
			if removeChoice > 0 && removeChoice <= len(dash.Transportation) {
				dash.Transportation = append(dash.Transportation[:removeChoice-1], dash.Transportation[removeChoice:]...)
				dash.Transportation = append(dash.Transportation, models.Transportation{Type: transType})
				fmt.Println("Kendaraan berhasil diperbarui.")
			} else if removeChoice != 0 {
				fmt.Println("Pilihan tidak valid.")
			}
		} else {
			dash.Transportation = append(dash.Transportation, models.Transportation{Type: transType})
			fmt.Println("Kendaraan berhasil ditambahkan.")
		}
	}
}

func ManageItems(dash *models.Dashboard) {
	for {
		fmt.Println("\n=== Kelola Barang ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Lihat Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Kembali ke Menu Utama")
		
		var choice int
		fmt.Print("Pilih opsi (1-4): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Input tidak valid, coba lagi.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		
		switch choice {
		case 1:
			var itemName string
			fmt.Print("Nama barang: ")
			scanner := bufio.NewReader(os.Stdin)
			itemName, _ = scanner.ReadString('\n')
			itemName = strings.TrimSpace(itemName)
			
			if itemName != "" {
				dash.Items = append(dash.Items, models.Item{Name: itemName})
				fmt.Println("Barang berhasil ditambahkan.")
			} else {
				fmt.Println("Nama barang tidak boleh kosong.")
			}
		case 2:
			if len(dash.Items) == 0 {
				fmt.Println("Belum ada barang yang ditambahkan.")
			} else {
				fmt.Println("Daftar Barang:")
				for i, item := range dash.Items {
					fmt.Printf("%d. %s\n", i+1, item.Name)
				}
			}
		case 3:
			if len(dash.Items) == 0 {
				fmt.Println("Belum ada barang yang ditambahkan.")
				continue
			}
			
			fmt.Println("Daftar Barang:")
			for i, item := range dash.Items {
				fmt.Printf("%d. %s\n", i+1, item.Name)
			}
			
			var removeChoice int
			fmt.Print("Pilih barang untuk dihapus (1-" + strconv.Itoa(len(dash.Items)) + ") atau 0 untuk batal: ")
			fmt.Scanln(&removeChoice)
			
			if removeChoice > 0 && removeChoice <= len(dash.Items) {
				dash.Items = append(dash.Items[:removeChoice-1], dash.Items[removeChoice:]...)
				fmt.Println("Barang berhasil dihapus.")
			} else if removeChoice != 0 {
				fmt.Println("Pilihan tidak valid.")
			}
		case 4:
			return
		}
	}
}

func ManageRecommendations(dash *models.Dashboard) {
	for {
		fmt.Println("\n=== Kelola Rekomendasi ===")
		fmt.Println("1. Tambah Rekomendasi")
		fmt.Println("2. Lihat Rekomendasi")
		fmt.Println("3. Hapus Rekomendasi")
		fmt.Println("4. Kembali ke Menu Utama")
		
		var choice int
		fmt.Print("Pilih opsi (1-4): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Input tidak valid, coba lagi.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		
		switch choice {
		case 1:
			var category, content string
			scanner := bufio.NewReader(os.Stdin)
			
			fmt.Print("Kategori (contoh: Film, Makanan): ")
			category, _ = scanner.ReadString('\n')
			category = strings.TrimSpace(category)
			
			fmt.Print("Isi (contoh: Avatar, Nasi Goreng): ")
			content, _ = scanner.ReadString('\n')
			content = strings.TrimSpace(content)
			
			if category != "" && content != "" {
				dash.Recommendations = append(dash.Recommendations, models.Recommendation{
					Category: category,
					Content:  content,
				})
				fmt.Println("Rekomendasi berhasil ditambahkan.")
			} else {
				fmt.Println("Kategori dan isi tidak boleh kosong.")
			}
		case 2:
			if len(dash.Recommendations) == 0 {
				fmt.Println("Belum ada rekomendasi yang ditambahkan.")
			} else {
				fmt.Println("Daftar Rekomendasi:")
				for i, rec := range dash.Recommendations {
					fmt.Printf("%d. %s: %s\n", i+1, rec.Category, rec.Content)
				}
			}
		case 3:
			if len(dash.Recommendations) == 0 {
				fmt.Println("Belum ada rekomendasi yang ditambahkan.")
				continue
			}
			
			fmt.Println("Daftar Rekomendasi:")
			for i, rec := range dash.Recommendations {
				fmt.Printf("%d. %s: %s\n", i+1, rec.Category, rec.Content)
			}
			
			var removeChoice int
			fmt.Print("Pilih rekomendasi untuk dihapus (1-" + strconv.Itoa(len(dash.Recommendations)) + ") atau 0 untuk batal: ")
			fmt.Scanln(&removeChoice)
			
			if removeChoice > 0 && removeChoice <= len(dash.Recommendations) {
				dash.Recommendations = append(dash.Recommendations[:removeChoice-1], dash.Recommendations[removeChoice:]...)
				fmt.Println("Rekomendasi berhasil dihapus.")
			} else if removeChoice != 0 {
				fmt.Println("Pilihan tidak valid.")
			}
		case 4:
			return
		}
	}
}

func ManageFriends(dash *models.Dashboard) {
	for {
		fmt.Println("\n=== Kelola Teman ===")
		fmt.Println("1. Tambah Teman")
		fmt.Println("2. Lihat Teman")
		fmt.Println("3. Hapus Teman")
		fmt.Println("4. Kembali ke Menu Utama")
		
		var choice int
		fmt.Print("Pilih opsi (1-4): ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Input tidak valid, coba lagi.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		
		switch choice {
		case 1:
			var name, division string
			scanner := bufio.NewReader(os.Stdin)
			
			fmt.Print("Nama teman: ")
			name, _ = scanner.ReadString('\n')
			name = strings.TrimSpace(name)
			
			fmt.Print("Divisi: ")
			division, _ = scanner.ReadString('\n')
			division = strings.TrimSpace(division)
			
			if name != "" && division != "" {
				dash.Friends = append(dash.Friends, models.Friend{
					Name:     name,
					Division: division,
				})
				fmt.Println("Teman berhasil ditambahkan.")
			} else {
				fmt.Println("Nama dan divisi tidak boleh kosong.")
			}
		case 2:
			if len(dash.Friends) == 0 {
				fmt.Println("Belum ada teman yang ditambahkan.")
			} else {
				fmt.Println("Daftar Teman:")
				for i, friend := range dash.Friends {
					fmt.Printf("%d. %s (Divisi: %s)\n", i+1, friend.Name, friend.Division)
				}
			}
		case 3:
			if len(dash.Friends) == 0 {
				fmt.Println("Belum ada teman yang ditambahkan.")
				continue
			}
			
			fmt.Println("Daftar Teman:")
			for i, friend := range dash.Friends {
				fmt.Printf("%d. %s (Divisi: %s)\n", i+1, friend.Name, friend.Division)
			}
			
			var removeChoice int
			fmt.Print("Pilih teman untuk dihapus (1-" + strconv.Itoa(len(dash.Friends)) + ") atau 0 untuk batal: ")
			fmt.Scanln(&removeChoice)
			
			if removeChoice > 0 && removeChoice <= len(dash.Friends) {
				dash.Friends = append(dash.Friends[:removeChoice-1], dash.Friends[removeChoice:]...)
				fmt.Println("Teman berhasil dihapus.")
			} else if removeChoice != 0 {
				fmt.Println("Pilihan tidak valid.")
			}
		case 4:
			return
		}
	}
}

func DisplayAllData(dash *models.Dashboard) {
	fmt.Println("\n=== Data Iftar GDCoC ===")
	
	fmt.Println("\nOpsi Kendaraan:")
	if len(dash.Transportation) == 0 {
		fmt.Println("Belum ada kendaraan yang dipilih.")
	} else {
		for i, trans := range dash.Transportation {
			fmt.Printf("%d. %s\n", i+1, trans.Type)
		}
	}
	
	fmt.Println("\nBarang yang Dibawa:")
	if len(dash.Items) == 0 {
		fmt.Println("Belum ada barang yang ditambahkan.")
	} else {
		for i, item := range dash.Items {
			fmt.Printf("%d. %s\n", i+1, item.Name)
		}
	}
	
	fmt.Println("\nRekomendasi:")
	if len(dash.Recommendations) == 0 {
		fmt.Println("Belum ada rekomendasi yang ditambahkan.")
	} else {
		for i, rec := range dash.Recommendations {
			fmt.Printf("%d. %s: %s\n", i+1, rec.Category, rec.Content)
		}
	}
	
	fmt.Println("\nTeman yang Ikut:")
	if len(dash.Friends) == 0 {
		fmt.Println("Belum ada teman yang ditambahkan.")
	} else {
		for i, friend := range dash.Friends {
			fmt.Printf("%d. %s (Divisi: %s)\n", i+1, friend.Name, friend.Division)
		}
	}
	
	fmt.Println("\nTekan Enter untuk kembali ke menu...")
	var discard string
	fmt.Scanln(&discard)
}