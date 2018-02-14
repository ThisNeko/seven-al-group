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
        while (true) {
            try {

                if (read.ready()) {
                    String etat = read.readLine();
                    if (etat != null) {

                        jsonObject = new JSONObject(etat);
                        jsonInfo = new JSONObject(jsonObject.getString("Info"));
                        System.out.println(jsonInfo.toString());
                        long id = jsonInfo.getLong("ID");
                        int posX = jsonInfo.getJSONObject("Position").getInt("X");
                        int posY = jsonInfo.getJSONObject("Position").getInt("Y");

                        List<Voiture> voitures = screen.voitures;
                        if (!voitures.isEmpty()) {
                            boolean voitureDansListe = false;
                            for (int j = 0; j < voitures.size(); j++) {
                                if (voitures.get(j).getId() == id) {
                                    voitures.get(j).setPositionX(posX);
                                    voitures.get(j).setPositionY(posY);
                                    voitureDansListe = true;
                                    screen.repaint();
                                }
                            }

                            if (voitureDansListe == false) {
                                Voiture v = new Voiture();
                                v.setId(id);
                                v.setPositionX(posX);
                                v.setPositionY(posY);
                                screen.voitures.add(v);
                                screen.repaint();

                            }

                        } else {
                            Voiture v = new Voiture();
                            v.setId(id);
                            v.setPositionX(posX);
                            v.setPositionY(posY);
                            screen.voitures.add(v);
                            screen.repaint();
                        }
                    }
                }
            } catch (IOException ex) {
                Logger.getLogger(Connexion.class.getName()).log(Level.SEVERE, null, ex);
            }
            
        }
    }
}
