package groupeseven.interfaceal;

import java.io.IOException;
import java.io.PrintWriter;
import java.net.Socket;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author jum_M
 */
public class Ping implements Runnable {

    private Socket socket = null;
    private PrintWriter out;

    public Ping(Socket s) {
        this.socket = s;
        try {
            out = new PrintWriter(s.getOutputStream());
        } catch (IOException ex) {
            Logger.getLogger(Ping.class.getName()).log(Level.SEVERE, null, ex);
        }
    }

    @Override
    public void run() {
        while (true) {
            try {
                out.println("ping\n");
                out.flush();
                Thread.sleep(2000);
            } catch (InterruptedException ex) {
                Logger.getLogger(Ping.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
    }
}
