package groupeseven.interfaceal;

import java.awt.Color;
import java.awt.Graphics;
import java.util.ArrayList;
import java.util.List;
import javax.swing.JPanel;

public class Screen extends JPanel {

    public List<Voiture> voitures;

    public Screen() {
        voitures = new ArrayList<Voiture>();
        //repaint();
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
            if (!voitures.get(i).isPanneVoiture()) {
                g.setColor(Color.BLACK);
                g.fillRect(voitures.get(i).getPositionX() % this.getWidth(), voitures.get(i).getPositionY() % this.getHeight(), 100, 100);
            }
        }

    }
}
