public interface Animal {
    void fly();
    void run();
    void bark();
}
public class Bird implements Animal {
    public void bark() { /* do nothing */ }
    public void run() {
        // write code about running of the bird
    }
    public void fly() {
        // write code about flying of the bird
    }
}
public class Cat implements Animal {
    public void fly() { throw new Exception("Undefined cat property"); }
    public void bark() { throw new Exception("Undefined cat property"); }
    public void run() {
        // write code about running of the cat
    }
}
public class Dog implements Animal {
    public void fly() { }
    public void bark() {
        // write code about barking of the dog
    }
    public void run() {
        // write code about running of the dog
    }
}