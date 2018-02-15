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
                        if (typeEnum.equals("VOITURE")) {
                            jsonInfo = new JSONObject(jsonObject.getString("Info"));
                            long id = jsonInfo.getLong("ID");
                            int posX = jsonInfo.getJSONObject("Position").getInt("X");
                            int posY = jsonInfo.getJSONObject("Position").getInt("Y");
                            int vitesse = jsonInfo.getJSONObject("Vitesse").getInt("X");
                            List<Voiture> voitures = screen.voitures;
                            if (!voitures.isEmpty()) {
                                boolean voitureDansListe = false;
                                for (int j = 0; j < voitures.size(); j++) {
                                    if (voitures.get(j).getId() == id) {
                                        voitures.get(j).setPositionX(posX);
                                        voitures.get(j).setPositionY(posY);
                                        voitures.get(j).setVitesseX(vitesse);
                                        voitures.get(j).setTimeStamp(System.currentTimeMillis());
                                        voitureDansListe = true;
                                    }
                                }

                                if (voitureDansListe == false) {
                                    Voiture v = new Voiture();
                                    v.setId(id);
                                    v.setPositionX(posX);
                                    v.setPositionY(posY);
                                    v.setVitesseX(vitesse);
                                    v.setTimeStamp(System.currentTimeMillis());
                                    screen.voitures.add(v);
                                }

                            } else {
                                Voiture v = new Voiture();
                                v.setId(id);
                                v.setPositionX(posX);
                                v.setPositionY(posY);
                                v.setVitesseX(vitesse);
                                v.setTimeStamp(System.currentTimeMillis());
                                screen.voitures.add(v);
                            }
                        } else if (typeEnum.equalsIgnoreCase("FEU")) {
                            jsonInfo = new JSONObject(jsonObject.getString("Info"));
                            System.out.println(jsonInfo.toString());
                            long id = jsonInfo.getLong("ID");
                            int posX = jsonInfo.getJSONObject("Position").getInt("X");
                            int posY = jsonInfo.getJSONObject("Position").getInt("Y");
                            int couleur = jsonInfo.getInt("Couleur");
                            int timer = jsonInfo.getInt("Ticker");

                            List<Feu> feux = screen.feux;
                            if (!feux.isEmpty()) {
                                boolean feuDansListe = false;
                                for (int j = 0; j < feux.size(); j++) {
                                    if (feux.get(j).getId() == id) {
                                        feux.get(j).setCouleur(couleur);
                                        feux.get(j).setTimer(timer);
                                        feuDansListe = true;
                                    }
                                }

                                if (feuDansListe == false) {
                                    Feu f = new Feu();
                                    f.setId(id);
                                    f.setPositionX(posX);
                                    f.setPositionY(posY);
                                    f.setCouleur(couleur);
                                    f.setTimer(timer);
                                    screen.feux.add(f);
                                }

                            } else {
                                Feu f = new Feu();
                                f.setId(id);
                                f.setPositionX(posX);
                                f.setPositionY(posY);
                                f.setCouleur(couleur);
                                f.setTimer(timer);
                                screen.feux.add(f);

                            }
                        }
                    }
                } else {
                    Thread.sleep(1);
                }
            } catch (IOException ex) {
                Logger.getLogger(Connexion.class.getName()).log(Level.SEVERE, null, ex);
            } catch (InterruptedException ex) {
                Logger.getLogger(Connexion.class.getName()).log(Level.SEVERE, null, ex);
            }
            screen.repaint();
        }
    }
}
