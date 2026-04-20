pub fn is_armstrong_number(num: u32) -> bool {
    let mut sum = 0;
    let digit_count = num.to_string().len();

    for x in num.to_string().chars() {
        let digit = x.to_digit(10).unwrap();
        sum += digit.pow(digit_count as u32);
    }

    sum == num
}
