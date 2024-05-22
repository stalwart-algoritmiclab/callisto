package chain

// LastBlock - describes an interface for working with database models.
type LastBlock interface {
	Get() (uint64, error)
	Update(id uint64) error
}
