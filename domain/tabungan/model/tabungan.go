package model

import "time"

type Rekening struct {
	ID          int64  `db:"id" json:"id"`
	Nama        string `db:"nama" json:"nama"`
	NIK         string `db:"nik" json:"nik"`
	NoHandphone string `db:"no_handphone" json:"no_handphone"`
	NoRekening  int64  `db:"no_rekening" json:"no_rekening"`
	Saldo       int64  `db:"saldo" json:"saldo"`
	CreatedBy   string `db:"created_by" json:"createted_by"`
}

type CreateRekeningRequest struct {
	Nama        string `json:"nama"`
	NIK         string `json:"nik"`
	NoHandphone string `json:"no_handphone"`
}

type SavingRekeningRequest struct {
	NoRekening int64 `json:"no_rekening"`
	Nominal    int64 `json:"nominal"`
}

type WitdrawalRekeningRequest struct {
	NoRekening int64 `json:"no_rekening"`
	Nominal    int64 `json:"nominal"`
}

type NasabahJSON struct {
	NoRekening int64 `json:"no_rekening"`
}

type SaldoJSON struct {
	NoRekening int64 `json:"no_rekening"`
	Saldo      int64 `json:"saldo"`
}

type Mutasi struct {
	ID            int64     `db:"id" json:"id"`
	NoRekening    int64     `db:"no_rekening" json:"no_rekening"`
	KodeTransaksi string    `db:"kode_transaksi" json:"kode_transaksi"`
	Nominal       int64     `db:"nominal" json:"nominal"`
	Saldo         int64     `db:"saldo" json:"saldo"`
	CreatedBy     string    `db:"created_by" json:"created_by"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type MutasiJSON struct {
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       int64     `json:"nominal"`
	Saldo         int64     `json:"saldo"`
	Waktu         time.Time `json:"waktu"`
}
