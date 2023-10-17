const mysql = require ('mysql2');


// Konfigurasi koneksi ke database sumber (Santri)
    const sourceDB = mysql.createConnection({
        host: 'localhost',
        user: 'root',
        password: 'Pastibisa',
        database: 'Santri',
    });

    // Konfigurasi koneksi ke database tujuan (pzn)
    const targetDB = mysql.createConnection({
        host: 'localhost',
        user: 'root',
        password: 'Pastibisa',
        database: 'goSantri',
    });

    // Query data dari database sumber
    sourceDB.query('SELECT * FROM Santri', (err, results) => {
        if (err) {
            console.error('Error querying source database:', err);
            return;
        }

        // Memasukkan data ke database tujuan
        const dataToInsert = results.map((row) => [row.nama, row.user_id, row.hp, row.email, row.gender, row.alamat, row.angkatan, row.jurusan, row.minat, row.status]);
        targetDB.query('INSERT INTO Santri (nama, user_id, hp, email, gender, alamat, angkatan, jurusan, minat, status) VALUES ?', [dataToInsert], (err) => {
            if (err) {
                console.error('Error inserting data into target database:', err);
            } else {
                console.log('Data successfully transferred to target database.');
            }
            // Tutup koneksi ke kedua database
            sourceDB.end();
            targetDB.end();
        });
    });
