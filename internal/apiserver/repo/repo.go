package repo

var client Repository

type Repository interface {
	Assets() AssetRepository
	Close() error
}

func Client() Repository {
	return client
}

func SetClient(c Repository) {
	client = c
}
