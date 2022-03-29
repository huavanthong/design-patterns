package main;

interface PaymentMethod {
    String Pay(int amount);
}

enum PaymentType { 
	Cash,
	DebitCard;
}


class CashPM implements PaymentMethod {
    private String methodString;

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
public class FactoryMethodDemo {
    public static void main(String [] args) {
        PaymentMethod payment = null;
        payment = new CashPM(PaymentType.Cash);
        payment.Pay(1000);
        payment = new DebitCardPM(PaymentType.DebitCard);
        payment.Pay(2000);
    }
}

