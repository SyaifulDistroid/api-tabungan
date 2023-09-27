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

	UpdateSaldoQuery = `
		UPDATE tb_rekening SET
			saldo = $2
		WHERE id = $1;
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

	CreateHistoryQuery = `
		INSERT INTO tb_history(
			no_rekening, 
			kode_transaksi, 
			nominal, 
			saldo, 
			created_by)
		VALUES ($1, $2, $3, $4, $5)
	`

	GetHistoryQuery = `
		SELECT id,
			no_rekening, 
			kode_transaksi, 
			nominal, 
			saldo,
			created_at
		 FROM tb_history
		 WHERE no_rekening = $1
		 ORDER BY created_at DESC;
	`
)
