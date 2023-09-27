CREATE TABLE IF NOT EXISTS tb_rekening(
    id bigserial NOT NULL primary key,
    nama varchar NOT NULL,
	nik varchar NOT NULL,
	no_handphone varchar NOT NULL,
	no_rekening bigint NOT NULL,
	saldo bigint DEFAULT 0,
	created_by varchar
);

CREATE TABLE IF NOT EXISTS tb_history(
	id bigserial NOT NULL primary key,
	no_rekening bigint,
	kode_transaksi varchar,
	nominal bigint,
	saldo bigint,
	created_by varchar,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)