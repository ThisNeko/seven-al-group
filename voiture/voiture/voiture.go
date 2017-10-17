package voiture

import "time"

type Materiel struct {
	ID int
	Vitesse float64
	Position Position
	Frein bool
}

func NewVoiture(ip string, materiel *Materiel) {
	frein := NewModuleFrein()
	mods := NewModuleDispatcher()
	mods.AddModule(frein)

	reg := NewRegistre(&mods)
	stat := NewStatus(&mods)

	conn := NewConnection(ip,&reg)

	for{
		<- time.After(50 * time.Millisecond)
		stat.Update(*materiel)
		conn.Broadcast(NewMessage(stat))
	}
}
