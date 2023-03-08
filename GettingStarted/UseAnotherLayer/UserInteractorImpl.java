package GettingStarted.UseAnotherLayer;

public class UserInteractorImpl implements UserInteractor {
    private final UserRepository userRepository;

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
