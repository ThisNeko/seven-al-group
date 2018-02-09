/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package groupeseven.interfaceal;

import java.io.BufferedReader;
import java.io.DataInputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.ObjectInputStream;
import java.net.Socket;
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.json.JSONObject;

/**
 *
 * @author jj128194
 */
public class Connexion implements Runnable {

    private Socket socket = null;
    private ObjectInputStream ois;
    private DataInputStream is;
    private static Screen screen;
    BufferedReader read;
    StringBuilder sb = new StringBuilder();

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
        while (true) {
            try {
                    String etat = read.readLine();
                    if (etat != null) {
                        
                        JSONObject jsonObject = new JSONObject(etat);
                        JSONObject jsonInfo = new JSONObject(jsonObject.getString("Info"));
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
                                    voitures.get(j).compteurDisparition = 0;
                                }
                                else
                                {
                                    voitures.get(j).compteurDisparition++;
                                }
                            }

                            if (voitureDansListe == false) {
                                Voiture v = new Voiture();
                                v.setId(id);
                                v.setPositionX(posX);
                                v.setPositionY(posY);
                                screen.voitures.add(v);
                            }
                            
                        } else {
                            Voiture v = new Voiture();
                            v.setId(id);
                            v.setPositionX(posX);
                            v.setPositionY(posY);
                            screen.voitures.add(v);
                        }
                    }
                screen.repaint();
            } catch (IOException ex) {
                Logger.getLogger(Connexion.class.getName()).log(Level.SEVERE, null, ex);
            }

        }
    }
}
