/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * Scanner class
 * This class will work with high-level class (Copy class)
 * @author Edward
 */
public class Scanner implements IReader {
 
    @Override
    public String read() {
        return "Sample of scanning data";
    }
 
}