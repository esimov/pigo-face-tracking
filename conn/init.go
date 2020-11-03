package conn

var httpConn = &HttpParams{
	Address: "localhost:5000",
	Prefix:  "/",
	Root:    ".",
}

func GetParams() *HttpParams {
	return httpConn
}
