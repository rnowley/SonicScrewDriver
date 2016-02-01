public class Greeter {
    private String message;

    public Greeter(String message) {
        this.message = message;
    }

    public void greet() {
        System.out.println(message);
    }
}