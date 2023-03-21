/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * Main class.
 * It's used for testing
 * @author Edward
 */
public class Main {
 
    /**
     * Main method
     * @author Edward
     * @param args
     */
    public static void main(String[] args) {
        // In this example, we will read a string from user's input
        // then write to database
        // You can try with other kinds of Reader and Writer you have
        IReader reader = new Keyboard();
        IWriter writer = new Database();
 
        Copy copy = new Copy(reader, writer);
 
        copy.doWork();
    }
 
}