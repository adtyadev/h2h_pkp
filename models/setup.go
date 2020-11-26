package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Connection() {

	user := goDotEnvVariable("DB_USERNAME")
	pass := goDotEnvVariable("DB_PASSWORD")
	host := goDotEnvVariable("DB_HOST")
	port := goDotEnvVariable("DB_PORT")
	dbName := goDotEnvVariable("DB_NAME")

	setupDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)
	//fmt.Println(setupDB)
	//dsn := "root:@tcp(127.0.0.1:3306)/h2h_pkp_2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := setupDB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func Migrate() {
	// Migrate the schema
	DB.AutoMigrate(&Request{}, &Agunan{}, &Throttle{}, &Response{})
}

func Seeder() {
	var seedRequest = Request{
		SecretKey:          "BA0F93E302BB99C0336AEE0D5FD1EF62845BC87F",
		KodeBranch:         "B123",
		NamaBranch:         "BNI Syariah cabang Depok",
		Cif:                "8798798712379",
		NamaDebitur:        "Atang Sanjaya",
		NoIdentitas:        "321298782810007",
		JenisKelamin:       "1",
		TglLahir:           "1978-03-29",
		Alamat:             "Jl Raya Depok",
		NoTelepon:          "0218789879879",
		KodePos:            "78901",
		Kk:                 "099881739131728",
		Npwp:               "1928198273127",
		Profesi:            "3",
		SektorUsaha:        "98080980",
		StatusPekerjaan:    "2",
		JenisAkad:          "MRB",
		NoAkad:             "AKAD/KUR/01",
		TglAkad:            "2020-08-04",
		TglPencairan:       "2020-08-05",
		NoLoan:             "LD99099898",
		PlafondPembiayaan:  40000000,
		NilaiJaminan:       40000000,
		NoRekening:         "7009899891",
		TglAwalPembiayaan:  "2020-08-05",
		TglAkhirPembiayaan: "2022-08-05",
		TglAwalPenjaminan:  "2020-08-05",
		TglAkhirPenjaminan: "2022-08-05",
		JangkaWaktu:        24,
		NoPermohonan:       "89888",
		TglPermohonan:      "2020-08-01",
		TglPembayaran:      "2020-08-05",
		BiayaMaterai:       0,
		BiayaAdmin:         0,
		NilaiIjk:           1000000,
		JumlahBayar:        1000000,
		ReferensiBayar:     "9923423423232",
	}
	DB.Create(&seedRequest)
	//c.JSON(http.StatusOK, gin.H{"data": seedRequest})
}
