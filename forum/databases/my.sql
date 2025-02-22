-- Create Users table
CREATE TABLE IF NOT EXISTS users  (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    email TEXT UNIQUE,
    password TEXT
);

-- Create Posts table
CREATE TABLE IF NOT EXISTS posts (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     title TEXT NOT NULL,
     descriptions TEXT NOT NULL,
     time TEXT,
     topic TEXT NOT NULL,
     likes INTEGER DEFAULT 0,
     dislikes INTEGER DEFAULT 0
);