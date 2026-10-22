package domain

type FavoritesList struct {
	Cities   []string
	Presence bool
	CacheErr error
}
