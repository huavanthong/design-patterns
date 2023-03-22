// Application package
package com.example.hexagonalarchitecture.application;

import com.example.hexagonalarchitecture.domain.User;
import com.example.hexagonalarchitecture.domain.UserRepository;

public class UserService {
    private final UserRepository userRepository;

    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public User save(User user) {
        return userRepository.save(user);
    }

    public User findById(Long id) {
        return userRepository.findById(id);
    }
}
