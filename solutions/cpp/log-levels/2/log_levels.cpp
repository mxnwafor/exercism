#include <string>

namespace log_line {
std::string message(std::string line) {
    // return the message
    std::size_t start = line.find(": ");
    if (start == std::string::npos) {
        return "";
    }
    return line.substr(start + 2);
}

std::string log_level(std::string line) {
    // return the log level
    std::size_t start = line.find('[');
    std::size_t end = line.find(']');

    if (start != std::string::npos && end != std::string::npos && end > start) {
        return line.substr(start + 1, end - start - 1);
    }
    return "";
        
}

std::string reformat(std::string line) {
    // return the reformatted message
    return message(line) + " (" + log_level(line) + ")";
}
}   // namespace log_line
