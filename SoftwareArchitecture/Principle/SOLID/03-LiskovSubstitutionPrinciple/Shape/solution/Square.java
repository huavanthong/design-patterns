/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.liskov;
 
/**
 * Class desciption
 * @author Edward
 */
public class Square extends Shape {
 
    @Override
    public int getWidth() {
        return mWidth;
    }
 
    @Override
    public void setWidth(int inWidth) {
        SetWidthAndHeight(inWidth);
    }
 
    @Override
    public int getHeight() {
        return mHeight;
    }
 
    @Override
    public void setHeight(int inHeight) {
        SetWidthAndHeight(inHeight);
    }
 
    /**
     * Set both width and height with the same value
     * @author Edward
     * @param inValue
     */
    private void SetWidthAndHeight(int inValue) {
        mHeight = inValue;
        mWidth = inValue;
    }
 
}