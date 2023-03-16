/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.singleresponsibility;
 
import edward.tutorial.designprinciple.singleresponsibility.PowerManagement.PowerState;
 
/**
 * This class will display the information of radio
 * @author Edward
 */
public class DisplayManagement {
 
    /** mRadio */
    private Radio mRadio;
 
    /**
     * @param inRadio
     */
    public DisplayManagement(Radio inRadio) {
        mRadio = inRadio;
    }
 
    /**
     * Power off the radio
     * @author Edward
     */
    public void diplayPowerState() {
        if(mRadio != null) {
            if(mRadio.getPowerState() == PowerState.OFF) {
                System.out.println("Radio has been powered off");
            }
            else if(mRadio.getPowerState() == PowerState.ON) {
                System.out.println("Radio has been powered on");
            }
        }
    }
 
    /**
     * Power on the radio
     * @author Edward
     */
    public void displayVolume() {
        if(mRadio != null &amp;amp;amp;&amp;amp;amp; mRadio.getPowerState() == PowerState.ON) {
            System.out.println("Volume: " + mRadio.getVolume());
        }
    }
}