/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.singleresponsibility;
 
/**
 * Class desciption
 * @author Edward
 */
public class Main {
 
    /**
     * Method desciption
     * @author Edward
     * @param args
     */
    public static void main(String[] args) {
        Radio radio = new Radio();
        radio.diplayPowerState();
        radio.displayVolume();
        radio.powerOn();
        radio.volumeUp();
        radio.diplayPowerState();
        radio.displayVolume();
    }
 
}