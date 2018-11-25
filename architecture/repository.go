type DBHandler interface {
	Execer
	Queryer
	Preparer
}

type UserRepository interface {
	Save(db DBHandler, userID int) error
}

