// Domain (Core) package
package com.example.hexagonalarchitecture.domain;

public interface UserRepository {
    User save(User user);
    User findById(Long id);
}
