type DBHandler interface {
	Execer
	Querier
	Preparer
}

type UserRepository interface {
	Save(db DBHandler, userID int) error
}

