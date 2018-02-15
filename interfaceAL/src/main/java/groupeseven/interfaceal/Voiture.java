/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package groupeseven.interfaceal;

import java.awt.*;

/**
 *
 * @author user
 */
public class Voiture {

    private static final int WIDTH = 3;
    private static final int HEIGHT = 1;

    private int positionX;
    private int positionY;

    private int vitesseX;

    private long id;
    public int compteurDisparition = 0;
    private boolean panneVoiture = false;
    private long timeStamp;

    private static final Font fonte = new Font("", Font.BOLD, 30);

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

    /**
     * @return the timeStamp
     */
    public long getTimeStamp() {
        return timeStamp;
    }
    
    /**
     * @param timeStamp the timeStamp to set
     */
    public void setTimeStamp(long timeStamp) {
        this.timeStamp = timeStamp;
    }

    public void draw(Graphics g, Screen screen){
        long timeLaps = System.currentTimeMillis() - timeStamp;
        if (timeLaps < 2000) {
            int x = screen.toScreenCoordsX(positionX);
            int y = screen.toScreenCoordsY(positionY);
            int width = screen.toScreenCoordsX(WIDTH);
            int height = screen.toScreenCoordsY(HEIGHT);

            if (panneVoiture) {
                g.setColor(Color.BLACK);
            } else if (id % 2 == 0) {
                g.setColor(Color.BLUE);
            } else {
                g.setColor(Color.ORANGE);
            }

            g.fillRect(x, y, width, height);
            g.setColor(Color.WHITE);
            g.setFont(fonte);
            g.drawString("" + (int) vitesseX, x, y + height);
        }
    }
}
