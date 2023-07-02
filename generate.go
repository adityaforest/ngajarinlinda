package ngajarinlinda

//go:generate swagger generate server --exclude-main -t gen -f ./swagger.yml
//go:generate swagger -q generate client -f ./swagger.yml -c pkg/client -m gen/models
