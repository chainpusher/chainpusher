package charge

type Repository interface {
	Find(id int64) (*Charge, error)
	Save(entity *Charge) error
}
