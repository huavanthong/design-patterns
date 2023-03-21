/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.openclose;
 
/**
 * OracleConnection
 * @author Edward
 */
public class OracleConnection extends Connection {
 
    /**
     * Connection process
     * @author Edward
     */
    public void doConnect() {
        System.out.println("Connect to Oracle Server");
    }
}