package db

type Option func(*DB)

// DBName is the name of the database to connect to.
func DBName(DBName string) Option {
	return func(db *DB) {
		db.Name = DBName
	}
}

// ConnectStyle is the style of the database to connect to.
func ConnectStyle(ConnectStyle string) Option {
	return func(db *DB) {
		db.ConnectStyle = ConnectStyle
	}
}

// LogLevel is the log level to use when connecting to the database.
func LogLevel(LogLevel string) Option {
	return func(db *DB) {
		db.LogLevel = LogLevel
	}
}

// PassWord is the password to use when connecting to the database.
func PassWord(PassWord string) Option {
	return func(db *DB) {
		db.PassWord = PassWord
	}
}

// Port is the port to use when connecting to the database.
func Port(Port string) Option {
	return func(db *DB) {
		db.Port = Port
	}
}

// UserName is the user to use when connecting to the database.
func UserName(UserName string) Option {
	return func(db *DB) {
		db.UserName = UserName
	}
}

// Host is the host to use when connecting to the database.
func Host(Host string) Option {
	return func(db *DB) {
		db.Host = Host
	}
}
