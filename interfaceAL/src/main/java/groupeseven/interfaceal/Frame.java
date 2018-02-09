/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package groupeseven.interfaceal;

import java.awt.GridLayout;
import java.io.IOException;
import java.net.Socket;
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.swing.JFrame;

/**
 *
 * @author user
 */
public class Frame extends JFrame {

    Screen s;
    public static Socket socket = null;
    public static Thread t1;

    public Frame() {
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setSize(800, 800);
        setResizable(false);
        setTitle("Interface");
        init();
        try {
            socket = new Socket("127.0.0.1", 1234);
        } catch (IOException ex) {
            Logger.getLogger(Frame.class.getName()).log(Level.SEVERE, null, ex);
        }

        System.out.println("Connexion établie avec le wifi");
        t1 = new Thread(new Connexion(socket, s));
        t1.start();
    }

    public void init() {
        setLocationRelativeTo(null);
        setLayout(new GridLayout(1, 1, 0, 0));
        s = new Screen();
        add(s);
        setVisible(true);
    }

    public static void main(String[] args) {
        new Frame();
    }
}
