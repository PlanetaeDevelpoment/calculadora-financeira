package calculadorarecisao

import "time"

type AvisoPrevio int

const (
	Trabalhado  = 0
	Indenizado  = 1
	Dispensado  = 2
	NãoCumprido = 3
)

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
