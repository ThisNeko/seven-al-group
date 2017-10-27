package voiture

import "time"

type Materiel struct {
	ID int
	Vitesse float64
	Position Position
	Frein bool
}

func NewVoiture(ip string, materiel *Materiel, conducteur Conducteur, after func(d time.Duration) <- chan time.Time) {

	mods := NewModuleDispatcher()

	reg := NewRegistre(&mods)
	conn := NewConnection(ip,&reg)
	stat := NewStatus(&mods,&conn)


	frein := NewModuleFrein(&reg,&stat, conducteur)
	mods.AddModule(frein)

	for{
		<- after(50 * time.Millisecond)
		stat.Update(*materiel)
		conn.Broadcast(NewMessage(stat))
	}
}
