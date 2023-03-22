// Infrastructure package
package com.example.hexagonalarchitecture.infrastructure;

import com.example.hexagonalarchitecture.domain.User;
import com.example.hexagonalarchitecture.domain.UserRepository;

public class InMemoryUserRepository implements UserRepository {
    @Override
    public User save(User user) {
        // Logic to save user to database
        return user;
    }

    @Override
    public User findById(Long id) {
        // Logic to find user by id from database
        return new User(id, "John Doe", "john.doe@example.com");
    }
}
