package controllers

import (
	"h2h_pkp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRequest(c *gin.Context) {
	var input models.Request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	request := models.Request{
		SecretKey:          input.SecretKey,
		KodeBranch:         input.KodeBranch,
		NamaBranch:         input.NamaBranch,
		Cif:                input.Cif,
		NamaDebitur:        input.NamaDebitur,
		NoIdentitas:        input.NoIdentitas,
		JenisKelamin:       input.JenisKelamin,
		TglLahir:           input.TglLahir,
		Alamat:             input.Alamat,
		NoTelepon:          input.NoTelepon,
		KodePos:            input.KodePos,
		Kk:                 input.Kk,
		Npwp:               input.Npwp,
		Profesi:            input.Profesi,
		SektorUsaha:        input.SektorUsaha,
		StatusPekerjaan:    input.StatusPekerjaan,
		JenisAkad:          input.JenisAkad,
		NoAkad:             input.NoAkad,
		TglAkad:            input.TglAkad,
		TglPencairan:       input.TglPencairan,
		NoLoan:             input.NoLoan,
		PlafondPembiayaan:  input.PlafondPembiayaan,
		NilaiJaminan:       input.NilaiJaminan,
		NoRekening:         input.NoRekening,
		TglAwalPembiayaan:  input.TglAwalPembiayaan,
		TglAkhirPembiayaan: input.TglAkhirPembiayaan,
		TglAwalPenjaminan:  input.TglAwalPenjaminan,
		TglAkhirPenjaminan: input.TglAkhirPenjaminan,
		JangkaWaktu:        input.JangkaWaktu,
		NoPermohonan:       input.NoPermohonan,
		TglPermohonan:      input.TglPermohonan,
		TglPembayaran:      input.TglPembayaran,
		BiayaMaterai:       input.BiayaMaterai,
		BiayaAdmin:         input.BiayaAdmin,
		NilaiIjk:           input.NilaiIjk,
		JumlahBayar:        input.JumlahBayar,
		ReferensiBayar:     input.ReferensiBayar,
	}

	models.DB.Create(&request)

	for _, value := range input.ListAgunan {
		var listAgunan models.Agunan
		listAgunan.IdRequest = request.ID
		listAgunan.JenisAgunan = value.JenisAgunan
		listAgunan.NilaiTaksasiKredit = value.NilaiTaksasiKredit
		listAgunan.NamaKepemilikan = value.NamaKepemilikan
		listAgunan.CaraPengikatanAgunan = value.CaraPengikatanAgunan
		listAgunan.AlamatAgunan = value.AlamatAgunan
		models.DB.Create(&listAgunan)

	}

	c.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func TestThrottle(c *gin.Context) {

	x_api_key := c.GetHeader("x_api_key")
	check := "123456789"
	if check != x_api_key {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Token x_api_key"})
		return
	}

	var input models.Throttle
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	throttle := models.Throttle{
		DataInteger:  input.DataInteger,
		DataInteger8: input.DataInteger8,
		DataFloat:    input.DataFloat,
		DataString:   input.DataString,
		DataDate:     input.DataDate,
		DataCount:    input.DataCount,
	}

	models.DB.Create(&throttle)
	c.JSON(http.StatusOK, gin.H{"status": "Success"})

}

func AddAgunan(c *gin.Context) {
	// var input models.DataAgunan
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error()})
	// 	return
	// }
	// for key, element := range input.ListAgunan {
	// 	fmt.Println("Key:", key, "=>", "Element:", element)
	// }
	// c.JSON(http.StatusOK, gin.H{"success": "True",
	// 	"data": input.Data, "agunana": input.ListAgunan})

}
