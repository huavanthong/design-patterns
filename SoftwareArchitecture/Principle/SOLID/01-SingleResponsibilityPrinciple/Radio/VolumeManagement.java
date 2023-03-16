/**
 *      OO Design Principle Tutorial
 *      View more at: https://edwardthienhoang.wordpress.com/
 */
package edward.tutorial.designprinciple.singleresponsibility;
 
/**
 * This class will manage the volume state of radio
 * @author Edward
 */
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