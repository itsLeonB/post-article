package appcontext

type ctxKey string

var KeyTx = ctxKey("tx")
var KeyLimit = ctxKey("limit")
var KeyOffset = ctxKey("offset")
