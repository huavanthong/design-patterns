# Reference: https://refactoring.guru/design-patterns/builder/python/example
from __future__ import annotations
from abc import ABC, abstractmethod
import string
from typing import Any


class Builder(ABC):
    """
    The Builder interface specifies methods for creating the different parts of
    the Product objects.
    """

    @property
    @abstractmethod
    def product(self) -> None:
        pass

    @abstractmethod
    def servings(self) -> None:
        pass

    @abstractmethod
    def calories(self) -> None:
        pass

    @abstractmethod
    def sodium(self) -> None:
        pass

    @abstractmethod
    def carbohydrate(self) -> None:
        pass


class ConcreteBuilder1(Builder):
    """
    The Concrete Builder classes follow the Builder interface and provide
    specific implementations of the building steps. Your program may have
    several variations of Builders, implemented differently.
    """

    def __init__(self) -> None:
        """
        A fresh builder instance should contain a blank product object, which is
        used in further assembly.
        """
        self.reset()

    def reset(self) -> None:
        self._product = NutritionFacts()

    @property
    def product(self) -> NutritionFacts:
        """
        Concrete Builders are supposed to provide their own methods for
        retrieving results. That's because various types of builders may create
        entirely different products that don't follow the same interface.
        Therefore, such methods cannot be declared in the base Builder interface
        (at least in a statically typed programming language).

        Usually, after returning the end result to the client, a builder
        instance is expected to be ready to start producing another product.
        That's why it's a usual practice to call the reset method at the end of
        the `getProduct` method body. However, this behavior is not mandatory,
        and you can make your builders wait for an explicit reset call from the
        client code before disposing of the previous result.
        """
        product = self._product
        self.reset()
        return product

    def servings(self, servingSize, servings) -> None:
        self._product.add(servingSize)
        self._product.add(servings)


    def calories(self, calories) -> None:
        self._product.add(calories)

    def sodium(self, sodium) -> None:
        self._product.add(sodium)

    def carbohydrate(self, carbohydrate) -> None:
        self._product.add(carbohydrate)

# Step 1: Implement class NutritionFacts
class NutritionFacts():

    def __init__(self) -> None:
        self.parts = []

    def add(self, part: Any) -> None:
        self.parts.append(part)

    def list_parts(self) -> None:
        print("Nutrition Facts: ")
        print(self.parts)

# Step 2: Implement Director will set builder
class Director:

    def __init__(self) -> None:
        self._builder = None

    # Step 2.1: Set builder to Builder class
    @property
    def builder(self) -> Builder:
        return self._builder

    @builder.setter
    def builder(self, builder: Builder) -> None:
        """
        The Director works with any builder instance that the client code passes
        to it. This way, the client code may alter the final type of the newly
        assembled product.
        """
        self._builder = builder

    """
    The Director can construct several product variations using the same
    building steps.
    """
    # The below block code will implement methods to initialize object
    def build_minimal_viable_product(self) -> None:
        self.builder.servings(270, 80)


    def build_full_featured_product(self) -> None:
        self.builder.servings(270, 80)
        self.builder.calories(100)
        self.builder.sodium(35)
        self.builder.sodium(27)


if __name__ == "__main__":
    """
    The client code creates a builder object, passes it to the director and then
    initiates the construction process. The end result is retrieved from the
    builder object.
    """

    director = Director()
    builder = ConcreteBuilder1()
    director.builder = builder

    print("Standard basic product: ")
    director.build_minimal_viable_product()
    builder.product.list_parts()

    print("\n")

    print("Standard full featured product: ")
    director.build_full_featured_product()
    builder.product.list_parts()

    print("\n")

    # Remember, the Builder pattern can be used without a Director class.
    print("Custom product: ")
    builder.servings(270, 80)
    builder.calories(333)
    builder.product.list_parts()