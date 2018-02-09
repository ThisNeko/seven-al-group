/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package groupeseven.interfaceal;

import java.util.Date;

/**
 *
 * @author user
 */
public class Voiture {

    private int positionX;
    private int positionY;

    private int vitesseX;
    private int vitesseY;

    private long id;
    public int compteurDisparition = 0;
    private boolean panneVoiture = false;
    public Date timeStamp = null;

    public Voiture() {

    }

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public int getVitesseX() {
        return vitesseX;
    }

    public void setVitesseX(int vitesseX) {
        this.vitesseX = vitesseX;
    }

    public int getVitesseY() {
        return vitesseY;
    }

    public void setVitesseY(int vitesseY) {
        this.vitesseY = vitesseY;
    }

    public int getPositionX() {
        return positionX;
    }

    public void setPositionX(int positionX) {
        this.positionX = positionX;
    }

    public int getPositionY() {
        return positionY;
    }

    public void setPositionY(int positionY) {
        this.positionY = positionY;
    }

    /**
     * @return the panneVoiture
     */
    public boolean isPanneVoiture() {
        return panneVoiture;
    }

    /**
     * @param panneVoiture the panneVoiture to set
     */
    public void setPanneVoiture(boolean panneVoiture) {
        this.panneVoiture = panneVoiture;
    }
}
