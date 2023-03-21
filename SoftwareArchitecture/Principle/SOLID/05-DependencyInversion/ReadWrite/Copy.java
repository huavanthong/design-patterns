/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.dependencyinversion;
 
/**
 * Copy class
 * I assume it is high-level class in this example
 * It only know 2 interfaces: IReader and IWriter
 * @author Edward
 */
public class Copy {
 
    /** mReader */
    private IReader mReader;
 
    /** mWriter */
    private IWriter mWriter;
 
    /**
     * Constructor
     * @param inReader
     * @param inWriter
     */
    public Copy(IReader inReader, IWriter inWriter) {
        if(inReader == null || inWriter == null) {
            throw new NullPointerException();
        }
        mReader = inReader;
        mWriter = inWriter;
    }
 
    /**
     * It will take an input by using a Reader
     * then write to an ouptut by using a Writer
     * @author Edward
     */
    public void doWork() {
        mWriter.write(mReader.read());
    }
}