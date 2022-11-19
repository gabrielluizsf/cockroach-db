package connection

func CloseConnError(err error)error{
	if err != nil {
		return err
	}

	return nil
}