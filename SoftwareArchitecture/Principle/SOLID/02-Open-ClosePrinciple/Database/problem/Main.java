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
 
        connectionPool.doMySqlConnect(new MySqlConnection());
        connectionPool.doSqlServerConnect(new SqlServerConnection());
    }
 
}

/* 
    Chạy hàm Main ta được:
    
    Connect to MySQL Server
    Connect to SQL Server
 */