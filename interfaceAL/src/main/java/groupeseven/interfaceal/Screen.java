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

    public Screen() {
        voitures = new ArrayList<Voiture>();
        feux = new ArrayList<Feu>();
    }

    public int toScreenCoordsX(double posX){
        return (int)(posX * 30);
    }

    public int toScreenCoordsY(double posY){
        return (int)(posY * 30);
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
            voitures.get(i).draw(g,this);
        }

        for (Feu f : feux) {
            f.draw(g,this);
        }

    }
}
