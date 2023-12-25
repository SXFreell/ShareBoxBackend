package dao

var (
	user string = `
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uid TEXT NOT NULL,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			email TEXT,
			create_time TEXT NOT NULL,
			update_time TEXT NOT NULL
		)
	`

	text string = `
		CREATE TABLE IF NOT EXISTS text (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uid TEXT NOT NULL,
			code TEXT NOT NULL,
			content TEXT NOT NULL,
			expires INTEGER NOT NULL,
			pickup_count INTEGER NOT NULL,
			create_time TEXT NOT NULL,
			update_time TEXT NOT NULL
		)
	`

	file string = `
		CREATE TABLE IF NOT EXISTS file (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uid TEXT NOT NULL,
			code TEXT NOT NULL,
			name TEXT NOT NULL,
			path TEXT NOT NULL,
			expires INTEGER NOT NULL,
			pickup_count INTEGER NOT NULL,
			create_time TEXT NOT NULL,
			update_time TEXT NOT NULL
		)
	`
)
