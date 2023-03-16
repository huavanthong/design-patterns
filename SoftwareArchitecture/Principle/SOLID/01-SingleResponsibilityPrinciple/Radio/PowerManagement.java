/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.singleresponsibility;
 
/**
 * This class will manage the power state of radio
 * @author Edward
 */
public class PowerManagement {
 
    /**
     * We have two state of Radio: ON or OFF
     * @author Edward
     */
    public enum PowerState {
        ON,
        OFF
    }
 
    /** mRadio */
    private Radio mRadio;
 
    /**
     * @param inRadio
     */
    public PowerManagement(Radio inRadio) {
        mRadio = inRadio;
    }
 
    /**
     * Power off the radio
     * @author Edward
     */
    public void off() {
        if(mRadio != null &amp;amp;amp;&amp;amp;amp; mRadio.getPowerState() == PowerState.ON) {
            mRadio.setPowerState(PowerState.OFF);
        }
    }
 
    /**
     * Power on the radio
     * @author Edward
     */
    public void on() {
        if(mRadio != null &amp;amp;amp;&amp;amp;amp; mRadio.getPowerState() == PowerState.OFF) {
            mRadio.setPowerState(PowerState.ON);
        }
    }
}