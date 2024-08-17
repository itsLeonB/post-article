package apperror

import "errors"

var ErrTxConversion = errors.New("fail to convert to sql.Tx")
