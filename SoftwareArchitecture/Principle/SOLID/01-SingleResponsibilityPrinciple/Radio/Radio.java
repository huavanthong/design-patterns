/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.singleresponsibility;
 
import edward.tutorial.designprinciple.singleresponsibility.PowerManagement.PowerState;
 
/**
 * Radio class
 * @author Edward
 */
public class Radio {
 
    /** mPowerState */
    private PowerState mPowerState = PowerState.OFF;
 
    /** mVolume */
    private int mVolume = VolumeManagement.DEFAULT_VOLUME;
 
    private PowerManagement mPowerManagement;
    private VolumeManagement mVolumeManagement;
    private DisplayManagement mDisplayManagement;
 
    public Radio() {
        mPowerManagement = new PowerManagement(this);
        mVolumeManagement = new VolumeManagement(this);
        mDisplayManagement = new DisplayManagement(this);
    }
 
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
        mPowerManagement.off();
    }
 
    /**
     * Power on the radio
     * @author Edward
     */
    public void powerOn() {
        mPowerManagement.on();
    }
 
    /**
     * Volume up the radio
     * @author Edward
     */
    public void volumeUp() {
        mVolumeManagement.up();
    }
 
    /**
     * Volume down the radio
     * @author Edward
     */
    public void volumeDown() {
        mVolumeManagement.down();
    }
 
    /**
     * Display power state of radio
     * @author Edward
     */
    public void diplayPowerState() {
        mDisplayManagement.diplayPowerState();
    }
 
    /**
     * Display volume state of radio
     * @author Edward
     */
    public void displayVolume() {
        mDisplayManagement.displayVolume();
    }
}