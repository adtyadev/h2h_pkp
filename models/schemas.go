package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	SecretKey          string       `json:"secretKey" binding:"required"`
	KodeBranch         string       `json:"kodeBranch" binding:"required"`
	NamaBranch         string       `json:"namaBranch" binding:"required"`
	Cif                string       `json:"cif" binding:"required"`
	NamaDebitur        string       `json:"namaDebitur" binding:"required"`
	NoIdentitas        string       `json:"noIdentitas" binding:"required"`
	JenisKelamin       string       `json:"jenisKelamin" binding:"required"`
	TglLahir           string       `json:"tglLahir" binding:"required"`
	Alamat             string       `json:"alamat" binding:"required"`
	NoTelepon          string       `json:"noTelepon" binding:"required"`
	KodePos            string       `json:"kodePos" binding:"required"`
	Kk                 string       `json:"kk" binding:"required"`
	Npwp               string       `json:"npwp" binding:"required"`
	Profesi            string       `json:"profesi" binding:"required"`
	SektorUsaha        string       `json:"sektorUsaha" binding:"required"`
	StatusPekerjaan    string       `json:"statusPekerjaan" binding:"required"`
	JenisAkad          string       `json:"jenisAkad" binding:"required"`
	NoAkad             string       `json:"noAkad" binding:"required"`
	TglAkad            string       `json:"tglAkad" binding:"required"`
	TglPencairan       string       `json:"tglPencairan" binding:"required"`
	NoLoan             string       `json:"noLoan" binding:"required"`
	PlafondPembiayaan  int          `json:"plafondPembiayaan" binding:"required"`
	NilaiJaminan       int          `json:"nilaiJaminan" binding:"required"`
	NoRekening         string       `json:"noRekening" binding:"required"`
	TglAwalPembiayaan  string       `json:"tglAwalPembiayaan" binding:"required"`
	TglAkhirPembiayaan string       `json:"tglAkhirPembiayaan" binding:"required"`
	TglAwalPenjaminan  string       `json:"tglAwalPenjaminan" binding:"required"`
	TglAkhirPenjaminan string       `json:"tglAkhirPenjaminan" binding:"required"`
	JangkaWaktu        int          `json:"jangkaWaktu" binding:"required"`
	ListAgunan         []DataAgunan `json:"listAgunan" gorm:"-" binding:"required,dive"` //`gorm:"-"`  // ignore this field when write and read
	NoPermohonan       string       `json:"noPermohonan" binding:"required"`
	TglPermohonan      string       `json:"tglPermohonan" binding:"required"`
	TglPembayaran      string       `json:"tglPembayaran" binding:"required"`
	BiayaMaterai       int          `json:"biayaMaterai" binding:"required"`
	BiayaAdmin         int          `json:"biayaAdmin" binding:"required"`
	NilaiIjk           int          `json:"nilaiIjk" binding:"required"`
	JumlahBayar        int          `json:"jumlahBayar" binding:"required"`
	ReferensiBayar     string       `json:"referensiBayar" binding:"required"`
}

type Agunan struct {
	gorm.Model
	IdRequest            uint    `json:"IdRequest" binding:"required"`
	JenisAgunan          int     `json:"JenisAgunan" binding:"required"`
	NilaiTaksasiKredit   float64 `json:"NilaiTaksasiKredit" binding:"required"`
	NamaKepemilikan      string  `json:"NamaKepemilikan" binding:"required"`
	CaraPengikatanAgunan int     `json:"CaraPengikatanAgunan" binding:"required"`
	AlamatAgunan         string  `json:"AlamatAgunan" binding:"required"`
}

type Response struct {
	gorm.Model
	NomorPolis   string    `json:"NomorPolis" binding:"required"`
	TanggalPolis time.Time `json:"TanggalPolis" binding:"required" gorm:"type:date" time_format:"2006-01-02"`
	NoPermohonan string    `json:"NoPermohonan" binding:"required"`
	Ijk          int       `json:"Ijk" binding:"required"`
	ResponseCode string    `json:"ResponseCode" binding:"required"`
	DescRescCode string    `json:"DescRescCode" binding:"required"`
}

type Throttle struct {
	gorm.Model
	DataInteger  int       `json:"DataInteger" binding:"required"`
	DataInteger8 int8      `json:"DataInteger8" binding:"required"`
	DataFloat    float64   `json:"DataFloat" binding:"required"`
	DataString   string    `json:"DataString" binding:"required" gorm:"size:50;not null"`
	DataDate     time.Time `json:"DataDate" binding:"required" gorm:"type:date" time_format:"2006-01-02"`
	DataCount    string    `json:"DataCount" binding:"required"`
}

type DataAgunan struct {
	JenisAgunan          int     `json:"JenisAgunan" binding:"required"`
	NilaiTaksasiKredit   float64 `json:"NilaiTaksasiKredit" binding:"required"`
	NamaKepemilikan      string  `json:"NamaKepemilikan" binding:"required"`
	CaraPengikatanAgunan int     `json:"CaraPengikatanAgunan" binding:"required"`
	AlamatAgunan         string  `json:"AlamatAgunan" binding:"required"`
}

// type DataAgunan struct {
// 	Data        string `binding:"required"`
// 	ListAgunan  map[string]string
// 	Information string
// 	Email       string `json:"email" binding:"required,email"`
// }
