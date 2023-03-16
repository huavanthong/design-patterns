/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.liskov;
 
/**
 * Class desciption
 * @author Edward
 */
public abstract class Shape {
 
    /** mHeight */
    protected int mHeight;
 
    /** mWidth */
    protected int mWidth;
 
    /**
     * Get shape's width
     * @author Edward
     * @return
     */
    public abstract int getWidth();
 
    /**
     * Set shape's width
     * @author Edward
     */
    public abstract void setWidth(int inWidth);
 
    /**
     * Get shape's height
     * @author Edward
     * @return
     */
    public abstract int getHeight();
 
    /**
     * Set shape's height
     * @author Edward
     */
    public abstract void setHeight(int inHeight);
 
    /**Return area of shape
     * @author Edward
     * @return
     */
    public int getArea() {
        return mHeight * mWidth;
    }
}