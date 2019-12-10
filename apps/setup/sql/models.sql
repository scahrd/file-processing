CREATE TABLE IF NOT EXISTS report
(
    id serial NOT NULL,
    cpf varchar(14) NOT NULL,
    privado boolean NOT NULL,
    incompleto boolean NOT NULL,
    data_ultima_compra date,
    ticket_medio numeric(10,2),
    ticket_ultima_compra numeric(10,2),
    loja_frequente varchar(18),
    loja_ultima_compra varchar(18),
    PRIMARY KEY (id)
);