public class Greeter {
    private string message = string.Empty;

    public Greeter(string message) {
        this.message = message;
    }

    public string SayHello() {
        return message;
    }
}