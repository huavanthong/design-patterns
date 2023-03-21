/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.openclose;
 
/**
 * Class desciption
 * @author Edward
 */
public class Main {
 
    /**
     * Main class
     * @author Edward
     * @param args
     */
    public static void main(String[] args) {
        ConnectionManager connectionPool = new ConnectionManager();
 
        connectionPool.doConnect(new MySqlConnection());
        connectionPool.doConnect(new OracleConnection());
        connectionPool.doConnect(new SqlServerConnection());
    }
 
}

/* 
 * Kết quả sẽ ra:
    Connect to MySQL Server
    Connect to Oracle Server
    Connect to SQL Server
 */