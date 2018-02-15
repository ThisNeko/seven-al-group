/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package groupeseven.interfaceal;

import java.awt.*;

/**
 *
 * @author jum_M
 */
public class Feu {
    private long id;
    private double positionX;
    private double positionY;
    private int couleur;
    private int timer;

    private static final double WIDTH = 1;
    private static final double HEIGHT = 1;

    private static final Font fonte = new Font("", Font.BOLD, 30);

    /**
     * @return the id
     */
    public long getId() {
        return id;
    }

    /**
     * @param id the id to set
     */
    public void setId(long id) {
        this.id = id;
    }

    /**
     * @return the positionX
     */
    public double getPositionX() {
        return positionX;
    }

    /**
     * @param positionX the positionX to set
     */
    public void setPositionX(double positionX) {
        this.positionX = positionX;
    }

    /**
     * @return the positionY
     */
    public double getPositionY() {
        return positionY;
    }

    /**
     * @param positionY the positionY to set
     */
    public void setPositionY(double positionY) {
        this.positionY = positionY;
    }

    /**
     * @return the couleur
     */
    public int getCouleur() {
        return couleur;
    }

    /**
     * @param couleur the couleur to set
     */
    public void setCouleur(int couleur) {
        this.couleur = couleur;
    }

    /**
     * @return the timer
     */
    public int getTimer() {
        return timer;
    }

    /**
     * @param timer the timer to set
     */
    public void setTimer(int timer) {
        this.timer = timer;
    }

    public void draw(Graphics g, Screen screen){
        int x = screen.toScreenCoordsX(positionX);
        int y = screen.toScreenCoordsY(positionY);
        int width = screen.toScreenCoordsX(WIDTH);
        int height = screen.toScreenCoordsY(HEIGHT);

        if (couleur == 1) {
            g.setColor(Color.RED);
        } else {
            g.setColor(Color.GREEN);
        }
        g.fillRect(x, y, width, height);
        g.setColor(Color.BLACK);
        g.setFont(fonte);
        g.drawString("" + timer, x, y + height / 2);
    }

}
