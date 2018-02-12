#include "traffic_lights_analyzer.hpp"
#include "io/driver_interface.hpp"

TrafficLightStatus* SelectTrafficLight(std::map <int, TrafficLightStatus> &trafficLightsRegistry, const CarStatus &carStatus)
{
    TrafficLightStatus *selected = nullptr;

    for (auto it = trafficLightsRegistry.begin(); it != trafficLightsRegistry.end(); ++it)
    {
        selected = &(it->second);        
    }

    if (selected == nullptr)
    {
        PrintToDriver("> TrafficLightsAnalyzer: No traffic light has been found.");
    }

    return selected;
}


double calculVitesse(TrafficLightStatus* tl, const CarStatus &carStatus){

    Vitesse vitesseVoiture = carStatus.vitesse;
    int tempsEcoule = tl->ticker;
    int tempsFeuTotal = tl->timer;
    int X1 = tl->pos.X;
    int Y1 = tl->pos.Y;
    int X2 = carStatus.position.X;
    int Y2 = carStatus.position.Y;

    double distanceFeuVoiture =  sqrt(pow(X2-X1,2)+pow(Y2-Y1,2));

    double tempsVoitureArriveFeu = (distanceFeuVoiture)/(vitesseVoiture.X/3.6);

    int modulo = fmod(tempsVoitureArriveFeu+double(tempsEcoule),2*tempsFeuTotal);

    if(tl->couleur==RED){
        if(modulo<=tempsFeuTotal){
            int nouvelleVitesse = int(distanceFeuVoiture)/int((tempsVoitureArriveFeu + (tempsFeuTotal - modulo)));
            //carStatus.
            return (double)nouvelleVitesse*3.6;
        } else {

        }
    } else {
        if(modulo<=tempsFeuTotal){

        } else {
            int nouvelleVitesse = int(distanceFeuVoiture)/int(tempsVoitureArriveFeu+ (2*tempsFeuTotal-(modulo+1)));
            return (double)nouvelleVitesse*3.6;
        }
    }




}
