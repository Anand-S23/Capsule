package store

type Store struct {
    UserRepo UserRepo
    ConnectionRepo ConnectionRepo
    MeetingRepo MeetingRepo
    ReminderRepo ReminderRepo
}

func NewStore(ur UserRepo, cr ConnectionRepo, mr MeetingRepo, rr ReminderRepo) *Store {
    return &Store {
        UserRepo: ur,
        ConnectionRepo: cr,
        MeetingRepo: mr,
        ReminderRepo: rr,
    }
}

