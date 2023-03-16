/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.singleresponsibility;
 
/**
 * Radio class
 * @author Edward
 */
public class Radio {
 
    /**
     * We have two state of Radio: ON or OFF
     * @author Edward
     */
    public enum PowerState {
        ON,
        OFF
    }
 
    public static final int DEFAULT_VOLUME = 50;
    public static final int MAX_VOLUME = 100;
    public static final int VOLUME_SEEK = 1;
 
    /** mPowerState */
    private PowerState mPowerState = PowerState.OFF;
 
    /** mVolume */
    private int mVolume = DEFAULT_VOLUME;
 
    /**
     * @return the mPowerState
     */
    public PowerState getPowerState() {
        return mPowerState;
    }
 
    /**
     * @param mPowerState the mPowerState to set
     */
    public void setPowerState(PowerState mPowerState) {
        this.mPowerState = mPowerState;
    }
 
    /**
     * @return the mVolume
     */
    public int getVolume() {
        return mVolume;
    }
 
    /**
     * @param mVolume the mVolume to set
     */
    public void setVolume(int mVolume) {
        this.mVolume = mVolume;
    }
 
    /**
     * Power off the radio
     * @author Edward
     */
    public void powerOff() {
        if(mPowerState == PowerState.ON) {
            mPowerState = PowerState.OFF;
        }
    }
 
    /**
     * Power on the radio
     * @author Edward
     */
    public void powerOn() {
        if(mPowerState == PowerState.OFF) {
            mPowerState = PowerState.ON;
        }
    }
 
    /**
     * Volume up the radio
     * @author Edward
     */
    public void volumeUp() {
        if(mVolume &amp;amp;lt; MAX_VOLUME) {           mVolume += VOLUME_SEEK;         }   }   /**      * Volume down the radio     * @author Edward    */     public void volumeDown() {      if(mVolume &amp;amp;gt; 0) {
            mVolume -= VOLUME_SEEK;
        }
    }
 
    /**
     * Display power state of radio
     * @author Edward
     */
    public void diplayPowerState() {
        if(mPowerState == PowerState.OFF) {
            System.out.println("Radio has been powered off");
        }
        else if(mPowerState == PowerState.ON) {
            System.out.println("Radio has been powered on");
        }
    }
 
    /**
     * Display volume state of radio
     * @author Edward
     */
    public void displayVolume() {
        if(mPowerState == PowerState.ON) {
            System.out.println("Volume: " + mVolume);
        }
    }
}