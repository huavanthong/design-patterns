package main;

import java.util.Enumeration;
import java.util.Vector;

//Step 1: Create a interface Alarm Listener
interface AlarmListener {
    void alarm();
}

//Step 2: Implement a Sensor sytem, where register senor
class SensorSystem {
    // Step 2.1: Create a vector mapping, we store all sensor
    private Vector listeners = new Vector();

    public void register(AlarmListener alarmListener) {
        listeners.addElement(alarmListener);
    }

    public void soundTheAlarm() {
        for (Enumeration e = listeners.elements(); e.hasMoreElements();) {
            ((AlarmListener) e.nextElement()).alarm();
        }
    }
}

//Step 3: Implement interface for functions 
class Lighting implements AlarmListener {
    public void alarm() {
        System.out.println("lights up");
    }
}

class Gates implements AlarmListener {
    public void alarm() {
        System.out.println("gates close");
    }
}

//Step 3: Create a class contain manys functions/feature operate on sensor
class CheckList {
 // Template Method design pattern
    public void byTheNumbers() {
        localize();
        isolate();
        identify();
    }

    protected void localize() {
        System.out.println("   establish a perimeter");
    }

    protected void isolate() {
        System.out.println("   isolate the grid");
    }

    protected void identify() {
        System.out.println("   identify the source");
    }
}

//class inherit.
//type inheritance
class Surveillance extends CheckList implements AlarmListener {
    public void alarm() {
        System.out.println("Surveillance - by the numbers:");
        byTheNumbers();
    }

    protected void isolate() {
        System.out.println("   train the cameras");
    }
}

public class ex2_sensor {
 public static void main( String[] args ) {
     SensorSystem sensorSystem = new SensorSystem();
     sensorSystem.register(new Gates());
     sensorSystem.register(new Lighting());
     sensorSystem.register(new Surveillance());
     sensorSystem.soundTheAlarm();
 }
}
