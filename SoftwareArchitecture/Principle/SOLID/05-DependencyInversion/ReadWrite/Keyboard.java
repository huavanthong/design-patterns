/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * Keyboard class
 * It's concrete class of IReader
 * @author Edward
 */
public class Keyboard implements IReader {
 
    @Override
    public String read() {
        java.util.Scanner in = new java.util.Scanner(System.in);
        System.out.print("Input a string: ");
        return in.nextLine();
    }
 
}
