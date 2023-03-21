/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.openclose;
 
/**
 * SqlServerConnection
 * @author Edward
 */
public class SqlServerConnection extends Connection {
 
    /**
     * Connection process
     * @author Edward
     */
    public void doConnect() {
        System.out.println("Connect to SQL Server");
    }
}