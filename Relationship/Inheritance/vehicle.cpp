// Reference: https://www.geeksforgeeks.org/inheritance-in-c/
// C++ program to implement Hierarchical Inheritance
#include<iostream>
using namespace std;
 
// base class
class Vehicle
{
  public:
    Vehicle()
    {
      cout << "This is a Vehicle\n";
    }
};
 
 
// first sub class
class Car: public Vehicle
{
    public:
        Car()
        {
            cout << "This is a Car\n";
        }
};
 
// second sub class
class Bus: public Vehicle
{
    public:
        Bus()
        {
            cout << "This is a Bus\n";
        }
};

// third sub class
class Truck: public Vehicle
{
    public:
        Truck()
        {
            cout << "This is a Truck\n";
        }
};
 
// main function
int main()
{  
    // Creating object of sub class will
    // invoke the constructor of base class.
    // Run object: Vehicle -> object: Car
    Vehicle obj;
    Car obj1;
    Bus obj2;
    Truck obj3;
    return 0;
}
