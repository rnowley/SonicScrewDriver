import org.restlet.Server;
import org.restlet.data.Protocol;
import restletdemo.HelloServerResource;

public class HelloServer {

    public static void main(String[] args) throws Exception {
        Server server = new Server(Protocol.HTTP, 3000, HelloServerResource.class);
        server.start();
    }

}