package main;

public class Patient {
    private String name;

    public Patient(String name) {
        this.name = name;
    }
}

public class Doctor {
    private String name;
    private List patients;

    public Doctor(String name) {
        this.name = name;
    }

    public void addPatient(Patient p) {
        patients.add(p);
    }
}

public class Main {
    public static void main(String [] args) {
        Patient dave = new Patient("Dave");
        Patient frank = new Patient("Frank");
        Patient bestsy = new Patient("Bestsy");

        Doctor james = new Doctor("James");
        Doctor scott = new Doctor("Scott")

        james.addPatient(dave);

        scott.addPatient(frank);
        scott.addPatient(bestsy);
    }
}