package voiture

import "time"

type Materiel struct {
	ID int
	Vitesse float64
	Position Position
	Panne bool
}

func NewVoiture(ip string, materiel *Materiel, conducteur Conducteur, after func(d time.Duration) <- chan time.Time) {

	mods := NewModuleDispatcher()

	data := NewData()
	conn := NewConnection(ip)
	//stat := NewStatus(&mods,&conn)

	data.Start(&mods,&conn)
	conn.Start(&data)


	frein := NewModuleFrein(&data, conducteur)
	mods.AddModule(frein)
	feu := NewModuleFeu(&data, conducteur)
	mods.AddModule(feu)
	panne := NewModulePanne(&data,conducteur)
	mods.AddModule(panne)

	for{
		<- after(50 * time.Millisecond)
		data.UpdateStatus(*materiel)
		conn.Broadcast(NewMessage(data))
	}
}
