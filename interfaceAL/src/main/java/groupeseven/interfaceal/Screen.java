package groupeseven.interfaceal;

import java.awt.Color;
import java.awt.Graphics;
import java.util.ArrayList;
import java.util.List;
import javax.swing.JPanel;

public class Screen extends JPanel {

    public List<Voiture> voitures;
    public List<Feu> feux;

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
            System.out.println(timeLaps);
            if (timeLaps < 2000) {
                if (!voitures.get(i).isPanneVoiture()) {
                    g.setColor(Color.BLACK);
                    g.fillRect(voitures.get(i).getPositionX() % this.getWidth(), voitures.get(i).getPositionY() % this.getHeight(), 100, 100);
                }else{
                    g.setColor(Color.BLACK);
                    g.fillRect(voitures.get(i).getPositionX() % this.getWidth(), voitures.get(i).getPositionY() % this.getHeight(), 100, 100);
                    g.setColor(Color.YELLOW);
                    //g.
                }
            }
        }

        for (Feu f : feux) {
            if (f.getCouleur() == 1) {
                g.setColor(Color.red);
            } else {
                g.setColor(Color.GREEN);
            }
            g.fillRect(f.getPositionX(), f.getPositionY(), 20, 100);
        }

    }
}
