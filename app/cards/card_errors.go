package cards

type Err_CardNotFound struct{}

func (m *Err_CardNotFound) Error() string {
	return "card not found"
}
