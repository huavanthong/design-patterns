package GettingStarted.UseAnotherLayer;

public interface IUserRepository {
    User findById(long id);
    List<User> findAll();
    void save(User user);
    void update(User user);
    void delete(User user);
}
