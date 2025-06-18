import sqlite3

conn = sqlite3.connect('admin.db')
c = conn.cursor()

# Create admin table
c.execute('''
CREATE TABLE IF NOT EXISTS admin (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
)
''')

# Insert default admin (you can change username & password)
c.execute("INSERT INTO admin (username, password) VALUES (?, ?)", ('Pritesh', 'Admin#_0812'))

conn.commit()
conn.close()
print("Database created and admin user added.")
