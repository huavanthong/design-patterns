/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * IWriter
 * This class will work with high-level class (Copy class)
 * @author Edward
 */
public interface IWriter {
 
    /**
     * High-level class (Copy class) will only work with this class
     * All low-level class must implement this method
     * so that high-level class can know it
     *
     * @author Edward
     * @param inInput
     */
    public void write(String inInput);
}