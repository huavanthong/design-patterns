package main;

// Step 1: Create a NutritionFacts class with members
public class NutritionFacts {
    private final int servingSize;
    private final int servings;
    private final int calories;
    private final int sodium;
    private final int carbohydrate;

    public class Builder {
        // Step 2: Require paramter
        // ===> Nghĩa là user cần phải input parameters.
        private final servingSize; 
        private final servings;

        // Step 3: Optional parameters, and initialize to default value.
        // ===> Nghĩa là user có quyền input hoặc không input các input parameters
        private int sodium;
        private int carbohydrate;

        // Step 4: Contructor
        // ===> Nghĩa là nếu cần object NutritioFacts, user phải sử dụng Constructor Builder này
        public Builder(int servingSize, int servings) {
            this.servingSize = servingSize;
            this.servings = servings;
        }

        // Step 5: Create methods with data type is Buidler class
        // ===> Nghĩa là user có thể sử dụng methods này để set parameters
        public Builder calories(int calories) {
            this.calories = calories;
            return this
        }

        // Step 5: 
        public Builder sodium(int sodium) {
            this.sodium = sodium;
            return this
        }

        // Step 5:
        public Builder carbohydrate(int carbohydrate) {
            this.carbohydrate = carbohydrate;
            return this
        }

        // Step 6: Create method with NutritionFacts class, with a input is a builder object
        public NutritionFacts build() {
            return new NutritionFacts(this);
        }
    }

    // Step 7: At main class, receive all input parameteres from Builder(0)
    private NutritionFacts(Buidler builder) {
        servingSize = builder.servingSize;
        servings = builder.servings;
        calories = builder.calories;
        sodium = builder.sodium;
        carbohydrate = builder.carbohydrate;
    }

    // Step 8: 
}