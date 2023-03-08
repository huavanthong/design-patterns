# Question
Context:  
    * Giả sử ta có 2 layer interactor và repostory.
    * Trong mỗi layer đều có define một list các phương thức dưới dạng interface.
    * Mỗi phương thức đều hoạt động đúng mục đích của nó.

Question:  
    Để interactor layer sử dụng các phương thức của repository layer chúng ta cân phải làm như thế nào?

# Ideas
Ta cần phải tự hiểu được kiến thức xây dựng cho kiến trúc sau.
1. Đầu tiên, phải xây dụng được interface cho layer repos đó.
2. Sau đó, phải implement code cho các interface tại layer repos đó.
3. Ở các class implementation ở layer khác, ta sẽ phải khai báo một đối tượng phụ thuộc và tiêm vào nó.
4. Ở App, ta sẽ tạo đối tượng repos, khởi tạo đối tượng interactor với parameter input là repos.
# Getting Started
**Step 1:** Tạo một interface cho repository layer, định nghĩa các phương thức cần thiết để thao tác với dữ liệu. Ví dụ:
```java
public interface UserRepository {
    User findById(long id);
    List<User> findAll();
    void save(User user);
    void update(User user);
    void delete(User user);
}

```
**Step 2:** Triển khai interface này trong repository layer, cung cấp logic thực hiện các phương thức đã định nghĩa. Ví dụ:
```java
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

```
**Step 3:** Trong interactor layer, khai báo một đối tượng của repository layer bằng cách sử dụng dependency injection. Ví dụ:
```java
public class UserInteractorImpl implements UserInteractor {
    private final UserRepository userRepository; // dependency injection

    public UserInteractorImpl(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public UserDTO getUser(long id) {
        User user = userRepository.findById(id);
        if (user != null) {
            return new UserDTO(user.getId(), user.getName(), user.getEmail());
        }
        return null;
    }

    public List<UserDTO> getAllUsers() {
        List<User> users = userRepository.findAll();
        return users.stream()
                .map(user -> new UserDTO(user.getId(), user.getName(), user.getEmail()))
                .collect(Collectors.toList());
    }

    public void createUser(UserDTO userDTO) {
        User user = new User(userDTO.getName(), userDTO.getEmail());
        userRepository.save(user);
    }

    public void updateUser(UserDTO userDTO) {
        User user = userRepository.findById(userDTO.getId());
        if (user != null) {
            user.setName(userDTO.getName());
            user.setEmail(userDTO.getEmail());
            userRepository.update(user);
        }
    }

    public void deleteUser(long id) {
        User user = userRepository.findById(id);
        if (user != null) {
            userRepository.delete(user);
        }
    }
}

```
**Step 4:** Sử dụng các phương thức của repository layer trong interactor layer bằng cách gọi các phương thức đã định nghĩa trong interface UserRepository. Ví dụ:
```java
UserRepository userRepository = new UserRepositoryImpl();
UserInteractor userInteractor = new UserInteractorImpl(userRepository);

UserDTO userDTO = userInteractor.getUser(1);
List<UserDTO> users = userInteractor.getAllUsers();

UserDTO newUser = new UserDTO("John", "john@example.com");
userInteractor.createUser(newUser);

UserDTO userToUpdate = userInteractor.getUser(1);
userToUpdate.setName("John Doe");
userToUpdate.setEmail("johndoe@example.com");
userInteractor.updateUser(userToUpdate);

userInteractor.deleteUser(1);
```
