package domain

type Repository interface {
	GetTweet(id int) (Tweet, bool, error)
	PostTweet(tweet Tweet) (Tweet, Transaction, error)
	SearchTweets() ([]Tweet, error)
}

type Transaction interface {
	Abort()
	Commit()
}
