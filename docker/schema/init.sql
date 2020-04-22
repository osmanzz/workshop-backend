CREATE TABLE mahasiswa (
    id serial not null,
    name varchar(300),
    password varchar(300),
    alamat text,
    primary key(id)
);

CREATE TABLE perpustakaan (
    id serial not null,
    mhs_id int not null,
    buku varchar(300),
    tanggal timestamp,
    primary key(id),
    foreign key(mhs_id) references mahasiswa(id)
);