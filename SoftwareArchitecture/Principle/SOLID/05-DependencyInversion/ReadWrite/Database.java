/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * Database class
 * It's concrete class of IWriter
 * @author Edward
 */
public class Database implements IWriter{
 
    @Override
    public void write(String inInput) {
        System.out.println("The data will be updated to database: " + inInput);
    }
 
}