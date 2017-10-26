package voiture

import "time"

type Materiel struct {
	ID int
	Vitesse float64
	Position Position
	Frein bool
}

func NewVoiture(ip string, materiel *Materiel) {

	mods := NewModuleDispatcher()

	reg := NewRegistre(&mods)
	conn := NewConnection(ip,&reg)
	stat := NewStatus(&mods,&conn)


	frein := NewModuleFrein(&reg,&stat)
	mods.AddModule(frein)

	for{
		<- time.After(50 * time.Millisecond)
		stat.Update(*materiel)
		conn.Broadcast(NewMessage(stat))
	}
}
