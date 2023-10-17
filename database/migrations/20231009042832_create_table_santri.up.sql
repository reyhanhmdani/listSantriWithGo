CREATE TABLE Santri (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NULL,
    nama VARCHAR(255),
    hp BIGINT,
    email VARCHAR(255),
    gender VARCHAR(255),
    alamat VARCHAR(255),
    angkatan INT,
    jurusan INT,
    minat INT,
    status INT,
    FOREIGN KEY (jurusan) REFERENCES Jurusan(id),
    FOREIGN KEY (minat) REFERENCES Minat(id),
    FOREIGN KEY (status) REFERENCES Status(id),
    FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
    CONSTRAINT FK_Status FOREIGN KEY (status) REFERENCES Status (id) ON DELETE SET NULL,
    CONSTRAINT FK_Jurusan FOREIGN KEY (jurusan) REFERENCES Jurusan(id) ON DELETE SET NULL,
    CONSTRAINT FK_Minat FOREIGN KEY (minat) REFERENCES Minat(id) ON DELETE SET NULL
);