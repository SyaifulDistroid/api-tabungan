package repository

const (
	CreateRekeningQuery = `
		INSERT INTO tb_rekening(
			nama, 
			nik, 
			no_handphone, 
			no_rekening, 
			created_by)
		VALUES ($1, $2, $3, $4, $5)
	`

	GetRekeningQuery = `
		SELECT id,
			nama, 
			nik, 
			no_handphone, 
			no_rekening,
			saldo
		 FROM tb_rekening
	`
)
