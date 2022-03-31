#*********************************************************************************************
# 
#  Reference: https://sourcemaking.com/design_patterns/factory_method/python/1
# 
#*********************************************************************************************
import abc

# Define một abstract các payment methods
#
class PaymentMethod(metaclass=abc.ABCMeta):
    """
    Declare the Factory Method
    """
    # Step 1: Create constructor for this class
    def __init__(self):
        # Assign a product variable without knowing
        self.product = self._Pay()

    # Step 2: Provide abstract method, 
    # other class must override this method
    @abc.abstractmethod
    def _Pay(self):
        pass

    # Step 2: Provide some parent method, 
    # other class can inherit
    def some_operation(self):
        self.product.interface()



class ConcreteCashPM(PaymentMethod):
    """
    Override the factory method to return an instance of a
    CashPM
    """

    def _Pay(self):
        return ConcretePayCash()


class ConcreteDebitCardPM(PaymentMethod):
    """
    Override the factory method to return an instance of a
    DebitCardPM.
    """

    def _Pay(self):
        return ConcreteDebitCard()


#===== Define a abstract class for payment. ====#
# Class này rất quan trọng, để cung cấp các abstract methods 
# cho các method Payment với các điểm lưu ý như sau:
# 1. Pay phải là một metaclass
# 2. Phải cung cấp các astract methods cho các Concrete class còn lại.
# 3. Phải sử dụng self như là một input parameters cho các methods
class Pay(metaclass=abc.ABCMeta):
    """
    Define the interface of objects the factory method creates.
    """
    # Step 1: 
    @abc.abstractmethod
    def interface(self):
        pass

    # Step 2: Nếu developer muốn 
    @abc.abstractmethod
    def payMoney(self, value):
        pass
    

#===== Concrete a abstract class =====#
# Class này sẽ implement các abstract class từ Pay
# Class này nhận input là một abstract Pay
class ConcretePayCash(Pay):
    """
    Implement the PayCash interface.
    """

    def interface(self):
        print("Pay by CashPM method")
        pass

    def payMoney(self, value):
        print("You pay: " + str(value))



class ConcreteDebitCard(Pay):
    """
    Implement the PayDebitCard interface.
    """

    def interface(self):
        print("Pay by DebitCardPM method")
        pass

    def payMoney(self, value):
        print("You pay: " + str(value))


def main():
    # Question 1: Đây là initialize a object từ ConcreteCashPM(),
    #            Nó giống gì bên Java?
    CashPM = ConcreteCashPM()
    # Question 2: Tại sao chúng ta cần phải có .product
    #           Chúng ta bỏ .product đi được không?
    CashPM.product.interface()
    CashPM.product.payMoney(1000)
    # Question 3: Với methods này thì cần làm gì? Nó thường implement ở đâu
    CashPM.some_operation()

    DebitCard = ConcreteDebitCardPM()
    DebitCard.product.interface()
    DebitCard.product.payMoney(2000)
    DebitCard.some_operation()



if __name__ == "__main__":
    main()
