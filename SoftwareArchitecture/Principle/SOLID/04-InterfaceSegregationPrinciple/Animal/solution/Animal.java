public interface Flyable {
    void fly();
}
public interface Runnable {
    void run();
}
public interface Barkable {
    void bark();
}
public class Bird implements Flyable, Runnable {
    public void run() {
        // write code about running of the bird
    }
    public void fly() {
        // write code about flying of the bird
    }
}
public class Cat implements Runnable{
    public void run() {
        // write code about running of the cat
    }
}
public class Dog implements Runnable, Barkable {
    public void bark() {
        // write code about barking of the dog
    }
    public void run() {
        // write code about running of the dog
    }
}