package main;

// Declare Factory Methods
interface PaymentMethod {
    String Pay(int amount);
}

// Declare Payment Type
enum PaymentType { 
	Cash,
	DebitCard;
}

// Make a method implement from common method.
class CashPM implements PaymentMethod {
    private String methodString;

    // To select exactly payment type we want to pay, 
    // we need to use PaymentType as input Parameters
    public CashPM(PaymentType type) {

        if(type.equals(PaymentType.Cash))
            this.methodString = "Pay by CashPM method";
    }

    @Override 
    public String Pay(int amount) {
        System.out.println(this.methodString);
        System.out.println(amount);
        return "OK";
    }
}


class DebitCardPM implements PaymentMethod {
    private String methodString;

    public DebitCardPM(PaymentType type) {
	    
    	if(type.equals(PaymentType.DebitCard))
            this.methodString = "Pay by DebitCardPM method";
    }
    @Override 
    public String Pay(int amount) {
        System.out.println(this.methodString);
        System.out.println(amount);
        return "OK";
    }
}

// With Factory Methods Design, we can select the type that we want to pay money.
public class FactoryMethodDemo {
    public static void main(String [] args) {
        PaymentMethod payment = null;
        payment = new CashPM(PaymentType.Cash);
        payment.Pay(1000);
        payment = new DebitCardPM(PaymentType.DebitCard);
        payment.Pay(2000);
    }
}

