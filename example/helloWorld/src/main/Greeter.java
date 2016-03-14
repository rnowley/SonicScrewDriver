/**
 * This class is used to display a greeting message.
 */
public class Greeter {
	private String message;

	/**
	 * Creates a new instance of a Greeter class.
	 * @param message The message to use as a greeting.
	 */
	public Greeter(String message) {
		this.message = message;
	}

	/**
	 * Retrieves the greeting message.
	 * @return The greeting as a string.
	 */
	public String greet() {
		return message;
	}
}
