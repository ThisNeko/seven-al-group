package groupeseven.interfaceal;

import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import javax.swing.JPanel;

public class Screen extends JPanel {

    public Map<Integer,Voiture> voitures;
    public Map<Integer,Feu> feux;

    public Screen() {
        voitures = new ConcurrentHashMap<>();
        feux = new ConcurrentHashMap<>();
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
        for (Map.Entry<Integer, Voiture> v : voitures.entrySet()) {
            v.getValue().draw(g,this);
        }

        for (Map.Entry<Integer, Feu> f : feux.entrySet()) {
            f.getValue().draw(g,this);
        }

    }
}
