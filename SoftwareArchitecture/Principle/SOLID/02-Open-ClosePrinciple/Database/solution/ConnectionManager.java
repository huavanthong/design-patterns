/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.openclose;
 
/**
 * ConnectionManager will help us connect to many kind of servers we want
 * @author Edward
 */
public class ConnectionManager {
 
    /**
     * Do connect to Oracle Server
     * @author Edward
     * @param inConnection
     */
    void doConnect(Connection inConnection) {
        // Let manager class do something before Server is started
        // ...
 
        // Connect to server
        inConnection.doConnect();
 
        // The manager will do something more after Server is stopped
        // ...
    }
}