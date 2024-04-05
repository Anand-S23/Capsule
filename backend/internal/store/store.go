package store

type Store struct {
    UserRepo UserRepo
    ConnectionRepo ConnectionRepo
}

func NewStore(ur UserRepo, cr ConnectionRepo) *Store {
    return &Store {
        UserRepo: ur,
        ConnectionRepo: cr,
    }
}

