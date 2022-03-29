package main;

public class BankAccount {
    private final String ownerName;
    private final int identificationNumber;
    private final String branch;
    private final int balance;

    public class Builder {
        private final String ownerName;
        private final int identificationNumber;
        private final String branch;
        private final int balance;

        public Builder() {

        }

        public Builder withOwnerName(String ownerName) {
            this.ownerName = ownerName;
            return this;
        }
        
        public Buidler WithOwnerIdentity(int identificationNumber) {
            this.identificationNumber = identificationNumber;
            return this;
        }

        public Buidler AtBranch(String branch) {
            this.branch = branch;
            return this;
        }

        public Builder OpenBalancing(int balance) {
            this.balance = balance;
            return this;
        }

        public BankAccount build() {
            return new BankAccount(this);
        }
    }

    public BankAccount(Buidler builder) {
        this.ownerName = builder.ownerName;
        this.identificationNumber = builder.identificationNumber;
        this.branch = builder.branch;
        this.balance = builder.balance;
    }

    public BankAccount Deposit(int deposit) {
        this.balance = this.balance + deposit;
        return this;
    }

    public BankAccount WithDraw(int withDraw) {
        this.balance = this.balance - withDraw;
        return this;
    }
}

public class BuilderBankDemo {
    public static void main(String [] args) {
        account = new BankAccount()
                            .WithOwnerName("Tuan")
                            .WithOwnerIdentity(123456789)
                            .AtBranch("Sai Gon")
                            .OpenBalancing(1000)
                            .build();

        Systems.out.println(account);

        account.Deposit(10000);
        Systems.out.println(account);
        account.WithDraw(50000);
        Systems.out.println(account);

    }
}