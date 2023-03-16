/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.liskov;
 
/**
 * Rectangle class
 * @author Edward
 */
public class Rectangle extends Shape
{
    @Override
    public int getWidth() {
        return mWidth;
    }
 
    @Override
    public int getHeight() {
        return mHeight;
    }
 
    @Override
    public void setWidth(int inWidth) {
        mWidth = inWidth;
    }
 
    @Override
    public void setHeight(int inHeight) {
        mHeight = inHeight;
    }
 
}