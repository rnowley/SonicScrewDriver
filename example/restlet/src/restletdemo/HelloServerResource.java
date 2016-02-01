package restletdemo;

import org.restlet.resource.Get;
import org.restlet.resource.ServerResource;

public class HelloServerResource extends ServerResource {

    @Get
    public String greet() {
        return "Hello from Restlet!";
    }

}