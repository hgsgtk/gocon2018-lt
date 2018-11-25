type SimpleServiceImpl struct {
	DB         repository.DBConnector
	User       repository.UserRepository
}

func (s *SimpleServiceImpl) Run(userID int) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	if err = s.User.Save(tx, userID); err != nil { // 引数に*Txを渡す
		wrapRollback(tx)
		return errors.Wrap(err, "failed to save user")
	}
	if err = wrapCommit(tx); err != nil {
		wrapRollback(tx)
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
