package connection

import "github.com/theGOURL/warning"

func CloseConnError(err error)error{
  warning.PRINT_DEFAULT_ERRORS(err,"cannot close database connection");
  	return nil;
}
