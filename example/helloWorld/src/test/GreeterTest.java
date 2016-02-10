import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class GreeterTest {

    @Test
    public void testGreetMethod() {
        Greeter greeterToTest = new Greeter("Hello test");
        String greeting = greeterToTest.greet();
        assertEquals("Hello test", greeting);
    }
}
