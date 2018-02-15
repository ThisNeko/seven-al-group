/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package groupeseven.interfaceal;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.Socket;
import java.util.List;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.json.JSONObject;

/**
 *
 * @author Maxime
 */
public class Connexion implements Runnable {

    private Socket socket = null;
    private static Screen screen;
    private BufferedReader read;

    public Connexion(Socket s, Screen screen) {
        this.socket = s;
        this.screen = screen;

        try {
            read = new BufferedReader(new InputStreamReader(socket.getInputStream()));
        } catch (IOException ex) {
            Logger.getLogger(Connexion.class.getName()).log(Level.SEVERE, null, ex);
        }
    }

    @Override
    public void run() {
        JSONObject jsonObject = null;
        JSONObject jsonInfo = null;
        String etat = "";
        while (true) {
            try {
                if (read.ready()) {
                    etat = read.readLine();
                    if (etat != null) {
                        jsonObject = new JSONObject(etat);
                        String typeEnum = jsonObject.getString("TypeEnum");

                        // VOITURES
                        if (typeEnum.equals("VOITURE")) {
                            jsonInfo = new JSONObject(jsonObject.getString("Info"));
                            int id = jsonInfo.getInt("ID");
                            int posX = jsonInfo.getJSONObject("Position").getInt("X");
                            int posY = jsonInfo.getJSONObject("Position").getInt("Y");
                            int vitesse = jsonInfo.getJSONObject("Vitesse").getInt("X");
                            boolean panne = jsonInfo.getBoolean("Panne");
                            Map<Integer,Voiture> voitures = screen.voitures;
                            if(voitures.containsKey(id)){
                                Voiture v = voitures.get(id);
                                v.setPositionX(posX);
                                v.setPositionY(posY);
                                v.setVitesseX(vitesse);
                                v.setTimeStamp(System.currentTimeMillis());
                                v.setPanneVoiture(panne);
                                voitures.replace(id,v);
                            } else{
                                Voiture v = new Voiture();
                                v.setId(id);
                                v.setPositionX(posX);
                                v.setPositionY(posY);
                                v.setVitesseX(vitesse);
                                v.setTimeStamp(System.currentTimeMillis());
                                v.setPanneVoiture(panne);
                                voitures.put(id,v);
                            }

                            // FEUX
                        } else if (typeEnum.equalsIgnoreCase("FEU")) {
                            jsonInfo = new JSONObject(jsonObject.getString("Info"));
                            System.out.println(jsonInfo.toString());
                            int id = jsonInfo.getInt("ID");
                            int posX = jsonInfo.getJSONObject("Position").getInt("X");
                            int posY = jsonInfo.getJSONObject("Position").getInt("Y");
                            int couleur = jsonInfo.getInt("Couleur");
                            int timer = jsonInfo.getInt("Ticker");

                            Map<Integer,Feu> feux = screen.feux;
                            if(feux.containsKey(id)){
                                Feu f = feux.get(id);
                                f.setPositionX(posX);
                                f.setPositionY(posY);
                                f.setCouleur(couleur);
                                f.setTimer(timer);
                                feux.replace(id,f);
                            } else {
                                Feu f = new Feu();
                                f.setId(id);
                                f.setPositionX(posX);
                                f.setPositionY(posY);
                                f.setCouleur(couleur);
                                f.setTimer(timer);
                            }
                        }
                    }
                } else {
                    Thread.sleep(1);
                }
            } catch (IOException | InterruptedException ex) {
                Logger.getLogger(Connexion.class.getName()).log(Level.SEVERE, null, ex);
            }
            screen.repaint();
        }
    }
}
