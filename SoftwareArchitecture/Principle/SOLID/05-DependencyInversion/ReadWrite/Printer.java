/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * Printer class
 * It's concrete class of IWriter
 * @author Edward
 */
public class Printer implements IWriter {
 
    @Override
    public void write(String inInput) {
        System.out.println("The text will be printed to paper: " + inInput);
    }
 
}