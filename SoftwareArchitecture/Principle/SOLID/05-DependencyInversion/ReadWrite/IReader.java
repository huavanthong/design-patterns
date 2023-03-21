/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * IReader class
 * This class will work with high-level class (Copy class)
 * @author Edward
 */
public interface IReader {
 
    /**
     * High-level class (Copy class) will only work with this class
     * All low-level class must implement this method
     * so that high-level class can know it
     *
     * @author Edward
     * @param inInput
     */
    public String read();
}