/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.liskov;
 
/**
 * Main class
 * @author Edward
 */
public class Main {
 
    static final String SQUARE = "Square";
    static final String RECTANGLE = "Rectangle";
 
    /**
     * Main method
     * @author Edward
     * @param args
     */
    public static void main(String[] args) {
        Shape shape1 = getShape(SQUARE);
        shape1.setHeight(10);
        shape1.setWidth(20);
        System.out.println(SQUARE + "'s area: " + shape1.getArea());
 
        Shape shape2 = getShape(RECTANGLE);
        shape2.setHeight(10);
        shape2.setWidth(20);
        System.out.println(RECTANGLE + "'s area: " + shape2.getArea());
    }
 
    /**
     * Simple factory method
     * @author Edward
     * @param inShapeType
     * @return
     */
    static Shape getShape(String inShapeType) {
        if(inShapeType.equals(SQUARE)) {
            return new Square();
        }
        if(inShapeType.equals(RECTANGLE)) {
            return new Rectangle();
        }
        return null;
    }
 
}

/*
 * Output:
Square’s area: 400
Rectangle’s area: 200
 */