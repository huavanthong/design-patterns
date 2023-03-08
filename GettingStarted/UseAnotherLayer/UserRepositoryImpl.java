package GettingStarted.UseAnotherLayer;

public class UserRepositoryImpl implements UserRepository {
    private List<User> users = new ArrayList<>();

    public User findById(long id) {
        return users.stream().filter(u -> u.getId() == id).findFirst().orElse(null);
    }

    public List<User> findAll() {
        return users;
    }

    public void save(User user) {
        users.add(user);
    }

    public void update(User user) {
        User oldUser = findById(user.getId());
        if (oldUser != null) {
            users.remove(oldUser);
            users.add(user);
        }
    }

    public void delete(User user) {
        users.remove(user);
    }
}
