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
        repaint();
    }

    @Override
    public void paint(Graphics g) {
        //g.fillRect(100, 0, 100, 100);
        super.paint(g);
        for (int i = 0; i < voitures.size(); i++) {
            if(!voitures.get(i).isPanneVoiture())
            {
                g.fillRect(voitures.get(i).getPositionX()%this.getWidth(), voitures.get(i).getPositionY()%this.getHeight(), 100, 100);
            }else
            {
                g.setColor(Color.red);
                g.fillRect(voitures.get(i).getPositionX()%this.getWidth(), voitures.get(i).getPositionY()%this.getHeight(), 100, 100);
            }
                
        }

    }
}
