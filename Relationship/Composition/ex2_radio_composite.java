/*
 * Reference: https://edwardthienhoang.wordpress.com/2013/11/09/single-responsibility-principle-srp/
 */
public class ex2_radio_composite {
    /** mPowerState */
    private PowerState mPowerState = PowerState.OFF;
 
    /** mVolume */
    private int mVolume = VolumeManagement.DEFAULT_VOLUME;
 
    private PowerManagement mPowerManagement;
    private VolumeManagement mVolumeManagement;
 
    public Radio() {
        mPowerManagement = new PowerManagement(this);
        mVolumeManagement = new VolumeManagement(this);
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
}


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

public class VolumeManagement {
 
    public static final int DEFAULT_VOLUME = 50;
    public static final int MAX_VOLUME = 100;
    public static final int VOLUME_SEEK = 1;
 
    /** mRadio */
    private Radio mRadio;
 
    /**
     * @param inRadio
     */
    public VolumeManagement(Radio inRadio) {
        mRadio = inRadio;
    }
 
    /**
     * Volume up the radio
     * @author Edward
     */
    public void up() {
        if(mRadio != null) {
            int volume = mRadio.getVolume();
            if(volume &amp;amp;lt; MAX_VOLUME) {                mRadio.setVolume(volume + VOLUME_SEEK);             }       }   }   /**      * Volume down the radio     * @author Edward    */     public void down() {        if(mRadio != null) {            int volume = mRadio.getVolume();            if(volume &amp;amp;gt; 0) {
                mRadio.setVolume(volume - VOLUME_SEEK);
            }
        }
    }
}