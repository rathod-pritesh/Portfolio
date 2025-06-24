import sqlite3
import os
from dotenv import load_dotenv

load_dotenv()

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

username = os.getenv('ADMIN_USERNAME')
password = os.getenv('ADMIN_PASSWORD')
c.execute("INSERT INTO admin (username, password) VALUES (?, ?)",  (username, password))

conn.commit()
conn.close()
print("Database created and admin user added.")
