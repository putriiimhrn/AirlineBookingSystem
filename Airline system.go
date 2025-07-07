package main

import (
	"fmt"
	"os"
	"strconv"
)

const N int = 8
const KB int = 4 /*--	kursi_Bisnis  */
const M int = 22
const KE int = 6 /*--	kursi_Ekonomi */
const T int = 5

type penumpang struct {
	np                string
	nama              string
	usia              int
	tempatDuduk       kursi
	kelas_penerbangan string
	kode_reservasi    int
	np_ortu           string
	diet              string
}

type pesawat struct {
	Bisnis  [N][KB]penumpang
	Ekonomi [M][KE]penumpang
}

type kursi struct {
	kelas        string
	baris, kolom int
}

var plane pesawat
var x penumpang
var counter_res int = 1
var array_penumpang [][]penumpang
var list_kursi []kursi

func main() {
	var pilih int

	for pilih != 8 {
		fmt.Print("\nSelamat datang di Maskapai Emirat.\nSilahkan pilih opsi berikut:\n")
		fmt.Print("\n[1] Reservasi\n[2] Check-In\n[3] Data Penumpang\n[4] Data Kursi\n[5] Cancel Reservasi\n[6] Cek Kapasitas\n[7] Print Data Penumpang\n[8] Selesai\n\nPilihan: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			reservasi()
		} else if pilih == 2 {
			var kode_reservasi int
			var kelas_penerbangan string
			fmt.Print("Input Nomor Reservasi: ")
			fmt.Scan(&kode_reservasi)
			fmt.Print("Input Kelas Penerbangan (Bisnis/Ekonomi): ")
			fmt.Scan(&kelas_penerbangan)
			checkIn(kode_reservasi, kelas_penerbangan)
		} else if pilih == 3 {
			var pilih2 int
			fmt.Print("\nSilahkan input pilihan berikut:\n\n[1] Mencari Orang Tua Saya\n[2] Mencari Anak Saya\n[3] List Anak\n[4] List Vegitarian\n[5] Mencari Penumpang\n[6] Kembali\n\nPilihan: ")
			fmt.Scan(&pilih2)
			if pilih2 == 1 {
				var np string
				fmt.Print("\n")
				fmt.Print("Input Nomor Paspor Anak: ")
				fmt.Scan(&np)
				x := mencariOrangTua(np)
				fmt.Print("\n========================================================\n")
				fmt.Print("\nNama\t\t\t: " + x.nama +
					"\nNo. Reservasi\t\t: " + strconv.Itoa(x.kode_reservasi) +
					"\nNo. Paspor Orang Tua\t: " + x.np_ortu +
					"\nNo. Paspor\t\t: " + x.np +
					"\nUsia\t\t\t: " + strconv.Itoa(x.usia) +
					"\nDiet\t\t\t: " + x.diet +
					"\nBaris, No. Kursi\t: " + strconv.Itoa(x.tempatDuduk.baris) + ", " + strconv.Itoa(x.tempatDuduk.kolom) + "\n" +
					"\n========================================================\n")

			} else if pilih2 == 2 {
				var np string
				fmt.Print("\n")
				fmt.Print("Input Nomor Paspor Orang Tua: ")
				fmt.Scan(&np)
				mencariAnak(np)
			} else if pilih2 == 3 {
				urutPenumpangAnak()
			} else if pilih2 == 4 {
				penumpangVegi()
			} else if pilih2 == 5 {
				var kode_reservasi int
				var np string
				fmt.Print("Input Nomor Reservasi: ")
				fmt.Scan(&kode_reservasi)
				fmt.Print("Input Nomor Paspor: ")
				fmt.Scan(&np)
				x := get_penumpang(kode_reservasi, np)
				fmt.Print("\n========================================================\n")
				fmt.Print("\nNama\t\t\t: " + x.nama +
					"\nNo. Reservasi\t\t: " + strconv.Itoa(x.kode_reservasi) +
					"\nNo. Paspor Orang Tua\t: " + x.np_ortu +
					"\nNo. Paspor\t\t: " + x.np +
					"\nUsia\t\t\t: " + strconv.Itoa(x.usia) +
					"\nDiet\t\t\t: " + x.diet +
					"\nBaris, No. Kursi\t: " + strconv.Itoa(x.tempatDuduk.baris) + ", " + strconv.Itoa(x.tempatDuduk.kolom) + "\n" +
					"\n========================================================\n")
			}
		} else if pilih == 4 {
			fmt.Println("Kelas Bisnis")
			for i := 0; i < len(plane.Bisnis); i++ {
				for j := 0; j < len(plane.Bisnis[j]); j++ {
					if plane.Bisnis[i][j].kode_reservasi == 0 {
						fmt.Print("[]")
					} else {
						fmt.Printf("[%d]", plane.Bisnis[i][j].kode_reservasi)
					}
				}
				fmt.Println("")
			}
			fmt.Println("")

			fmt.Println("Kelas Ekonomi")
			for i := 0; i < len(plane.Ekonomi); i++ {
				for j := 0; j < len(plane.Ekonomi[j]); j++ {
					if plane.Ekonomi[i][j].kode_reservasi == 0 {
						fmt.Print("[]")
					} else {
						fmt.Printf("[%d]", plane.Ekonomi[i][j].kode_reservasi)
					}
				}
				fmt.Println("")
			}
			fmt.Println("")
			fmt.Println("Jumlah Kursi Kosong Pada Kelas Bisnis:", len(kursiKosong("Bisnis")))
			fmt.Println("Jumlah Kursi Kosong Pada Kelas Ekonomi:", len(kursiKosong("Ekonomi")))
			fmt.Println("")
		} else if pilih == 5 {
			var kode_reservasi int
			fmt.Print("Input Nomor Reservasi: ")
			fmt.Scan(&kode_reservasi)
			cancelReservasi(kode_reservasi)
		} else if pilih == 6 {
			var kelas_penerbangan string
			fmt.Print("Input Kelas Penerbangan (Bisnis/Ekonomi): ")
			fmt.Scan(&kelas_penerbangan)
			if cek_kapasitas(kelas_penerbangan) {
				fmt.Println("Kelas", kelas_penerbangan, "sudah penuh")
			} else {
				fmt.Println("Kelas", kelas_penerbangan, "belum penuh")
			}
		} else if pilih == 7 {
			printPenumpang()
		}
	}
	fmt.Print("\n")
	fmt.Print("Terima kasih telah menggunakan Maskapai Emirat.")
	fmt.Print("\n")

}

func cek_kapasitas(kelas_penerbangan string) bool {
	/*diberikan string kelas penerbangan.
	mengembalikan true jika tidak ada data penumpang yang kosong.
	mengembalikan false jika ada data penumpang yang kosong */
	return len(kursiKosong(kelas_penerbangan)) == 0
}

func reservasi() {
	/*
		I.S.
		F.S.sistem akan membaca jenis reservasinya Bisnis atau Ekonomi
			setelah mengetahui jenis reservasi, sistem akan membaca jumlah penumpang yang reservasi,
			jika jumlah penumpang reservasi < maksimal jumlah reservasi, sistem akan mengecek apakah ada kursi kosong sejumlah penumpang reservasi.
			jika masih ada kursi kosong,maka sistem meminta user memasukkan detail data penumpang satu-persatu. kemudian sistem akan membentuk data reservasi yang berisi 1 nomor reservasi dan list data penumpang
	*/
	var x []penumpang
	var y penumpang
	var dewasa, remaja, bayi int
	var ortu, kelas string
	for kelas != "Bisnis" && kelas != "Ekonomi" {
		fmt.Print("\nSilahkan Pilih Kelas Penerbangan (Bisnis/Ekonomi): ")
		fmt.Scan(&kelas)
	}

	if kelas == "Bisnis" {
		fmt.Println("Maksimal 5 orang untuk 1 reservasi.")
		fmt.Print("Jumlah Dewasa (Umur > 18 tahun): ")
		fmt.Scan(&dewasa)
		fmt.Print("Jumlah Remaja (2-18 tahun): ")
		fmt.Scan(&remaja)
		fmt.Print("Jumlah Bayi (Umur < 2 tahun): ")
		fmt.Scan(&bayi)
		jumlah := dewasa + remaja + bayi
		jumlah_tanpa_bayi := dewasa + remaja
		if jumlah <= T {
			if len(kursiKosong(kelas)) >= jumlah_tanpa_bayi {
				for i := 0; i < jumlah; i++ {
					y.kode_reservasi = counter_res
					y.kelas_penerbangan = "Bisnis"
					fmt.Print("\n")
					fmt.Println("Data Penumpang", i+1)
					fmt.Print("\n")
					fmt.Print("Input Nama: ")
					fmt.Scan(&y.nama)
					fmt.Print("Input Nomor Paspor: ")
					fmt.Scan(&y.np)
					fmt.Print("Input Usia: ")
					fmt.Scan(&y.usia)
					for i+1 == 1 && y.usia <= 18 {
						fmt.Println("Data penumpang 1 harus orang dewasa")
						fmt.Print("Silahkan input kembali: ")
						fmt.Scan(&y.usia)
					}

					for y.diet != "Ya" && y.diet != "Tidak" {
						fmt.Print("Apakah Anda Vegetarian? (Ya/Tidak): ")
						fmt.Scan(&y.diet)
					}
					fmt.Println("Nomor reservasi Anda adalah", counter_res)
					if y.usia > 18 {
						ortu = y.np
						y.np_ortu = "-"
					} else {
						y.np_ortu = ortu
					}
					x = append(x, y)
					y.diet = "0"
				}
				array_penumpang = append(array_penumpang, x)
				counter_res++
			} else {
				fmt.Println("Mohon maaf, Kelas Bisnis saat ini penuh.")
				fmt.Println("")
			}

		} else {
			fmt.Println("Jumlah reservasi Anda melebihi maksimal.")
			fmt.Println("")
		}

	} else if kelas == "Ekonomi" {
		fmt.Println("Maksimal 5 orang untuk 1 reservasi.")
		fmt.Print("Jumlah Dewasa (Umur > 18 tahun): ")
		fmt.Scan(&dewasa)
		fmt.Print("Jumlah Remaja (2-18 tahun): ")
		fmt.Scan(&remaja)
		fmt.Print("Jumlah Bayi (Umur < 2 tahun): ")
		fmt.Scan(&bayi)
		jumlah := dewasa + remaja + bayi
		jumlah_tanpa_bayi := dewasa + remaja
		if jumlah < T {
			if len(kursiKosong(kelas)) >= jumlah_tanpa_bayi {
				for i := 0; i < jumlah; i++ {
					y.kode_reservasi = counter_res
					y.kelas_penerbangan = "Ekonomi"
					fmt.Print("\n")
					fmt.Println("Data Penumpang", i+1)
					fmt.Print("\n")
					fmt.Print("Input Nama: ")
					fmt.Scan(&y.nama)
					fmt.Print("Input Nomor Paspor: ")
					fmt.Scan(&y.np)
					fmt.Print("Input Usia: ")
					fmt.Scan(&y.usia)
					for y.diet != "Ya" && y.diet != "Tidak" {
						fmt.Print("Apakah Anda Vegetarian? (Ya/Tidak): ")
						fmt.Scan(&y.diet)
					}
					fmt.Println("Nomor reservasi Anda adalah", counter_res)
					if y.usia > 18 {
						ortu = y.np
						y.np_ortu = "-"
					} else {
						y.np_ortu = ortu
					}
					x = append(x, y)
					y.diet = "0"
				}
				array_penumpang = append(array_penumpang, x)
				counter_res++
			} else {
				fmt.Println("Mohon maaf, Kelas Ekonomi saat ini penuh.")
				fmt.Println("")
			}
		} else {
			fmt.Println("Jumlah reservasi Anda melebihi maksimal.")
			fmt.Println("")
		}
	}
}

func checkIn(kode_reservasi int, kelas_penerbangan string) {
	/*
		I.S diberikan nomor reservasi dan kelas penerbangan
		F.S sistem akan membaca nomor reservasi dan membaca data penumpang,kemudian mengupdate posisi tempat duduk penumpang reservasi
			tempat duduk akan dipilihkan yang berdekatan asumsi data orang tua selalu pertama dengan urutan kedekatan adalah nomor berurutan dan bayi memiliki tempatDuduk yang sama dengan orang tua
	*/
	var pilihan string
	var kursi_ortu kursi

	if kelas_penerbangan == "Bisnis" {
		listKursi := kursiKosong(kelas_penerbangan)
		fmt.Print("\n========================================================\n")
		for i := 0; i < len(array_penumpang[kode_reservasi-1]); i++ {
			x := array_penumpang[kode_reservasi-1][i]
			if x.usia >= 2 {
				x.tempatDuduk = kursi{baris: listKursi[i].baris, kolom: listKursi[i].kolom}
				plane.Bisnis[listKursi[i].baris][listKursi[i].kolom] = x
				if x.usia > 18 {
					kursi_ortu = x.tempatDuduk
				}
			} else {
				x.tempatDuduk = kursi_ortu

			}
			fmt.Print("\nNama\t\t\t: " + x.nama +
				"\nNo. Reservasi\t\t: " + strconv.Itoa(x.kode_reservasi) +
				"\nNo. Paspor Orang Tua\t: " + x.np_ortu +
				"\nNo. Paspor\t\t: " + x.np +
				"\nUsia\t\t\t: " + strconv.Itoa(x.usia) +
				"\nDiet\t\t\t: " + x.diet +
				"\nBaris, No. Kursi\t: " + strconv.Itoa(x.tempatDuduk.baris) + ", " + strconv.Itoa(x.tempatDuduk.kolom) + "\n" +
				"\n========================================================\n")

		}
	}
	if kelas_penerbangan == "Bisnis" {
		listKursi := kursiKosong(kelas_penerbangan)
		fmt.Print("\n========================================================\n")
		for i := 0; i < len(array_penumpang[kode_reservasi-1]); i++ {
			x := array_penumpang[kode_reservasi-1][i]
			if x.usia >= 2 {
				x.tempatDuduk = kursi{baris: listKursi[i].baris, kolom: listKursi[i].kolom}
				plane.Ekonomi[listKursi[i].baris][listKursi[i].kolom] = x
				if x.usia > 18 {
					kursi_ortu = x.tempatDuduk
				}
			} else {
				x.tempatDuduk = kursi_ortu

			}
			fmt.Print("\nNama\t\t\t: " + x.nama +
				"\nNo. Reservasi\t\t: " + strconv.Itoa(x.kode_reservasi) +
				"\nNo. Paspor Orang Tua\t: " + x.np_ortu +
				"\nNo. Paspor\t\t: " + x.np +
				"\nUsia\t\t\t: " + strconv.Itoa(x.usia) +
				"\nDiet\t\t\t: " + x.diet +
				"\nBaris, No. Kursi\t: " + strconv.Itoa(x.tempatDuduk.baris) + ", " + strconv.Itoa(x.tempatDuduk.kolom) + "\n" +
				"\n========================================================\n")

		}
	}

	fmt.Print("\n")
	fmt.Print("Ingin mengubah kursi dan/atau diet? (Ya/Tidak): ")
	fmt.Scan(&pilihan)
	if pilihan == "Ya" {
		ubah_diet_tempat_duduk()
	}
}

func get_penumpang(kode_reservasi int, np string) penumpang {
	/* diberikan nomor reservasi dan nomor paspor untuk mencari penumpang */
	var x penumpang
	for j := 0; j < len(array_penumpang); j++ {
		for i := 0; i < len(array_penumpang[j]); i++ {
			if array_penumpang[j][i].kode_reservasi == kode_reservasi && array_penumpang[j][i].np == np {
				return array_penumpang[j][i]
			}
		}
	}
	return x
}
func ubah_diet_tempat_duduk() {
	/*
		I.S
		F.S diet penumpang berubah sesuai nilai d, tempat duduk penumpang berubah sesuai nilai nd
	*/
	var kode int
	var np, d string
	var nd kursi
	var y int
	var z penumpang
	fmt.Println("Silahkan input data sebelumnya jika tidak ada perubahan!")
	fmt.Print("Nomor Reservasi: ")
	fmt.Scan(&kode)
	fmt.Print("Nomor Paspor: ")
	fmt.Scan(&np)
	fmt.Print("Apakah Anda Vegetarian? (Ya/Tidak): ")
	fmt.Scan(&d)
	fmt.Print("\n")

	for i := 0; i < len(array_penumpang[kode-1]); i++ {
		if array_penumpang[kode-1][i].np == np {
			y = i
		}
	}
	x := array_penumpang[kode-1][y]

	if x.kelas_penerbangan == "Bisnis" {
		fmt.Print("\nKursi Kosong Pada Kelas Bisnis:\n")
		b := kursiKosong("Bisnis")
		for i := 0; i < len(b); i++ {
			if i%3 == 0 {
				fmt.Print("\n")
			}
			fmt.Printf("[%d %d] ", b[i].baris, b[i].kolom)
		}
		fmt.Print("\n")
		fmt.Print("Baris kursi yang diinginkan: ")
		fmt.Scan(&nd.baris)
		fmt.Print("Nomor kursi yang diinginkan: ")
		fmt.Scan(&nd.kolom)
		plane.Bisnis[x.tempatDuduk.baris][x.tempatDuduk.kolom] = z
		x.tempatDuduk = nd
		if x.usia > 18 {
			for i := 0; i < len(array_penumpang); i++ {
				if array_penumpang[kode-1][i].np_ortu == np && array_penumpang[kode-1][i].usia < 2 {
					array_penumpang[kode-1][i].tempatDuduk = nd
				}
			}
		}
		x.diet = d
		plane.Bisnis[nd.baris][nd.kolom] = x
	} else {
		e := kursiKosong("Ekonomi")
		fmt.Print("\nKursi Kosong Pada Kelas Ekonomi:\n")
		for i := 0; i < len(e); i++ {
			if i%3 == 0 {
				fmt.Print("\n")
			}
			fmt.Printf("[%d %d] ", e[i].baris, e[i].kolom)
		}
		fmt.Print("\n")
		fmt.Print("Baris kursi yang diinginkan: ")
		fmt.Scan(&nd.baris)
		fmt.Print("Nomor kursi yang diinginkan: ")
		fmt.Scan(&nd.kolom)
		plane.Ekonomi[x.tempatDuduk.baris][x.tempatDuduk.kolom] = z
		x.tempatDuduk = nd
		if x.usia > 18 {
			for i := 0; i < len(array_penumpang); i++ {
				if array_penumpang[kode-1][i].np_ortu == np && array_penumpang[kode-1][i].usia < 2 {
					array_penumpang[kode-1][i].tempatDuduk = nd
				}
			}
		}
		x.diet = d
		plane.Ekonomi[nd.baris][nd.kolom] = x
	}

}
func mencariOrangTua(np string) penumpang {
	/* diberikan nomor passpor orang tua dari anak untuk mencari orang tua */
	var x penumpang
	for j := 0; j < len(array_penumpang); j++ {
		for i := 0; i < len(array_penumpang[j]); i++ {
			if array_penumpang[j][i].np == np {
				return array_penumpang[j][i]
			}
		}
	}
	return x
}

func penumpangVegi() {
	/*diberikan data pesawat. menghasilkan data penumpang dengan diet vegetarian(data penumpang Bisnis yg vegetarian + data penumpang Ekonomi yg vegetarian)*/
	fmt.Print("\n========================================================\n")
	fmt.Print("\n")
	fmt.Println("List Penumpang yang vegetarian:")
	for j := 0; j < len(array_penumpang); j++ {
		for i := 0; i < len(array_penumpang[j]); i++ {
			if array_penumpang[j][i].diet == "Ya" {
				fmt.Println("- ", array_penumpang[j][i].nama)
			}
		}
	}
	fmt.Print("\n========================================================\n")

}

func urutPenumpangAnak() {
	/*
		I.S
		F.S menghasilkan list penumpang anak yang terurut berdasarkan usia beserta orang tuanya*/
	var list_anak []penumpang
	for j := 0; j < len(array_penumpang); j++ {
		for i := 0; i < len(array_penumpang[j]); i++ {
			if array_penumpang[j][i].usia <= 18 {
				list_anak = append(list_anak, array_penumpang[j][i])
			}
		}
	}
	var pass, idx, inc int
	var temp penumpang
	for pass < len(list_anak)-1 {
		idx = pass
		inc = pass + 1
		for inc < len(list_anak) {
			if list_anak[idx].usia > list_anak[inc].usia {
				idx = inc
			}
			inc++
		}
		temp = list_anak[idx]
		list_anak[idx] = list_anak[pass]
		list_anak[pass] = temp
		pass++
	}
	fmt.Print("\n========================================================\n")
	fmt.Print("\n")
	for i := 0; i < len(list_anak); i++ {
		y := mencariOrangTua(list_anak[i].np_ortu)
		fmt.Println("Anak:", list_anak[i].nama, "\tUsia:", list_anak[i].usia, "\tOrang tua:", y.nama)
	}
	fmt.Print("\n========================================================\n")
}

func kursiKosong(kelas_penerbangan string) []kursi {
	/*diberikan data pesawat dan list kursi. menghasilkan list kursi kosong(nomor kursi tidak ada di data penumpang pesawat)*/
	var x penumpang
	list_kursi = nil
	if kelas_penerbangan == "Bisnis" {
		for j := 0; j < len(plane.Bisnis); j++ {
			for i := 0; i < len(plane.Bisnis[j]); i++ {
				if plane.Bisnis[j][i] == x {
					list_kursi = append(list_kursi, kursi{kelas: "Bisnis", baris: j, kolom: i})
				}
			}
		}
		return list_kursi
	} else {
		for j := 0; j < len(plane.Ekonomi); j++ {
			for i := 0; i < len(plane.Ekonomi[j]); i++ {
				if plane.Ekonomi[j][i] == x {
					list_kursi = append(list_kursi, kursi{kelas: "Ekonomi", baris: j, kolom: i})
				}
			}
		}
		return list_kursi
	}

}

func printPenumpang() {
	var data [200][200]string
	for j := 0; j < len(array_penumpang); j++ {
		for i := 0; i < len(array_penumpang[j]); j++ {
			x := array_penumpang[j][i]
			data[j][i] = string("\n========================================================\n" + "\n" +
				"\nNama\t\t\t: " + x.nama +
				"\nNo. Reservasi\t\t: " + strconv.Itoa(x.kode_reservasi) +
				"\nNo. Paspor Orang Tua\t: " + x.np_ortu +
				"\nNo. Paspor\t\t: " + x.np +
				"\nUsia\t\t\t: " + strconv.Itoa(x.usia) +
				"\nDiet\t\t\t: " + x.diet +
				"\nBaris, No. Kursi\t: " + strconv.Itoa(x.tempatDuduk.baris) + ", " + strconv.Itoa(x.tempatDuduk.kolom) + "\n" +
				"\n========================================================\n")
		}
	}

	f, err := os.Create("DataPenumpang.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	for _, v := range data {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
==================================================================================================================================================================================
2 Fungsionalitas Tambahan
==================================================================================================================================================================================
*/

func mencariAnak(np string) {
	/*diberikan nomor passpor orang tua dan akan di print nama anak dari orang tua tersebut */

	fmt.Print("\n========================================================\n")
	for j := 0; j < len(array_penumpang); j++ {
		for i := 0; i < len(array_penumpang[j]); i++ {
			if array_penumpang[j][i].np == np {
				fmt.Printf("\nAnak dari orang tua %s adalah\n", array_penumpang[j][i].nama)
			}
			if array_penumpang[j][i].np_ortu == np {
				fmt.Println("- ", array_penumpang[j][i].nama)
			}
		}
	}
	fmt.Print("\n========================================================\n")
}

func cancelReservasi(kode_reservasi int) {
	/*
		I.S data penumpang dengan nomor reservasi nomorRes ada di list penumpang
		F.S data penumpang dengan nomor reservasi nomorRes tidak ada di list penumpang
	*/
	var x penumpang
	for i := 0; i < len(array_penumpang); i++ {
		if array_penumpang[kode_reservasi][i].kode_reservasi == kode_reservasi {
			array_penumpang[kode_reservasi][i] = x
		}
	}

}
