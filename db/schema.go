package db

var Schema = [...]string{
	`CREATE TABLE IF NOT EXISTS board (
		slug text UNIQUE PRIMARY KEY,
		name text,
		descr text
	)`,
	`CREATE TABLE IF NOT EXISTS post (
		no INT NOT NULL,
		resto INT NOT NULL,
		board text NOT NULL,
		com text NOT NULL,
		time TIMESTAMP,
		UNIQUE (no, board),
		FOREIGN KEY (board) REFERENCES board(slug)
	)`,
}
