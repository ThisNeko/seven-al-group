package groupeseven.interfaceal;

import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics;
import java.util.ArrayList;
import java.util.List;
import javax.swing.JPanel;

public class Screen extends JPanel {

    public List<Voiture> voitures;
    public List<Feu> feux;
    private int sizeCarWidht = 100;
    private int sizeCarHeight = 100;

    private int sizeFeuWidht = 30;
    private int sizeFeuHeight = 100;
    Font fonte = new Font("", Font.BOLD, 30);

    public Screen() {
        voitures = new ArrayList<Voiture>();
        feux = new ArrayList<Feu>();
    }

    public void update(Graphics g) {
        paint(g);
    }

    @Override
    public void paint(Graphics g) {
        super.paint(g);
        g.setColor(Color.WHITE);
        g.fillRect(0, 0, this.getWidth(), this.getHeight());
        for (int i = 0; i < voitures.size(); i++) {
            long timeLaps = System.currentTimeMillis() - voitures.get(i).getTimeStamp();
            if (timeLaps < 2000) {
                if (!voitures.get(i).isPanneVoiture()) {

                    if (voitures.get(i).getId() % 2 == 0)//voiture en go
                    {
                        g.setColor(Color.BLUE);
                        g.fillRect(voitures.get(i).getPositionX() % this.getWidth(), voitures.get(i).getPositionY() % this.getHeight(), getSizeCarWidht(), getSizeCarHeight());
                        g.setColor(Color.WHITE);
                        g.setFont(fonte);
                        g.drawString("" + voitures.get(i).getVitesseX(), (int) (voitures.get(i).getPositionX() + getSizeCarWidht() * 0.40), voitures.get(i).getPositionY() + getSizeCarHeight() / 2);
                    } else {//voiture en c++
                        g.setColor(Color.ORANGE);
                        g.fillRect(voitures.get(i).getPositionX() % this.getWidth(), voitures.get(i).getPositionY() % this.getHeight(), getSizeCarWidht(), getSizeCarHeight());
                        g.setColor(Color.WHITE);
                        g.setFont(fonte);
                        g.drawString("" + voitures.get(i).getVitesseX(), (int) (voitures.get(i).getPositionX() + getSizeCarWidht() * 0.40), voitures.get(i).getPositionY() + getSizeCarHeight() / 2);
                    }

                } else {
                    g.setColor(Color.BLACK);
                    g.fillRect(voitures.get(i).getPositionX() % this.getWidth(), voitures.get(i).getPositionY() % this.getHeight(), getSizeCarWidht(), getSizeCarHeight());
                    g.setColor(Color.WHITE);
                    g.setFont(fonte);
                    g.drawString("" + voitures.get(i).getVitesseX(), (int) (voitures.get(i).getPositionX() + getSizeCarWidht() * 0.40), voitures.get(i).getPositionY() + getSizeCarHeight() / 2);

                }
            }
        }

        for (Feu f : feux) {
            if (f.getCouleur() == 1) {
                g.setColor(Color.red);
            } else {
                g.setColor(Color.GREEN);
            }
            g.fillRect(f.getPositionX(), f.getPositionY(), getSizeFeuWidht(), getSizeFeuHeight());
            g.setColor(Color.BLACK);
            g.setFont(fonte);
            g.drawString("" + f.getTimer(), f.getPositionX(), f.getPositionY() + sizeFeuHeight / 2);

        }

    }

    /**
     * @return the sizeCarWidht
     */
    public int getSizeCarWidht() {
        return sizeCarWidht;
    }

    /**
     * @param sizeCarWidht the sizeCarWidht to set
     */
    public void setSizeCarWidht(int sizeCarWidht) {
        this.sizeCarWidht = sizeCarWidht;
    }

    /**
     * @return the sizeCarHeight
     */
    public int getSizeCarHeight() {
        return sizeCarHeight;
    }

    /**
     * @param sizeCarHeight the sizeCarHeight to set
     */
    public void setSizeCarHeight(int sizeCarHeight) {
        this.sizeCarHeight = sizeCarHeight;
    }

    /**
     * @return the sizeFeuWidht
     */
    public int getSizeFeuWidht() {
        return sizeFeuWidht;
    }

    /**
     * @param sizeFeuWidht the sizeFeuWidht to set
     */
    public void setSizeFeuWidht(int sizeFeuWidht) {
        this.sizeFeuWidht = sizeFeuWidht;
    }

    /**
     * @return the sizeFeuHeight
     */
    public int getSizeFeuHeight() {
        return sizeFeuHeight;
    }

    /**
     * @param sizeFeuHeight the sizeFeuHeight to set
     */
    public void setSizeFeuHeight(int sizeFeuHeight) {
        this.sizeFeuHeight = sizeFeuHeight;
    }
}
