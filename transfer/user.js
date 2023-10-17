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
sourceDB.query('SELECT * FROM users', (err, results) => {
    if (err) {
        console.error('Error querying source database:', err);
        return;
    }

    // Memasukkan data ke database tujuan
    const dataToInsert = results.map((row) => [row.username, row.password, row.as_id]);
    targetDB.query('INSERT INTO UserSantri (username, password, as_id) VALUES ?', [dataToInsert], (err) => {
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
