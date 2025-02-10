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
     comment TEXT NOT NULL,
     topic TEXT NOT NULL
);