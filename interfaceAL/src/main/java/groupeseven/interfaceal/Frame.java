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
    public static Thread t2;

    public Frame() {
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setTitle("Interface");
        setSize(800, 800);
        setVisible(true);
        setLocationRelativeTo(null);
        setLayout(new GridLayout(1, 1, 0, 0));
        s = new Screen();
        setContentPane(s);
        
        try {
            socket = new Socket("127.0.0.1", 1234);
        } catch (IOException ex) {
            Logger.getLogger(Frame.class.getName()).log(Level.SEVERE, null, ex);
        }

        System.out.println("Connexion établie avec le wifi");
        t1 = new Thread(new Connexion(socket, s));
        t2 = new Thread(new Ping(socket));
        t1.start();
        t2.start();
    }
    
    public static void main(String[] args) {
        new Frame();
    }
}
