def clean_phone_number(number: str) -> str:
    number = number.replace("(", "")
    number = number.replace(")", "")
    number = number.replace("-", "")
    number = number.replace("+", "")
    if len(number) == 11: 
        number = number[1:]
    elif len(number) == 7: 
        number = "495" + number
    return number

numbers = list()
number = input()
numbers.append(clean_phone_number(number))
for i in range(3):
    number = input()
    cleaned_number = clean_phone_number(number)
    if cleaned_number in numbers: 
        print("YES")
    else: 
        numbers.append(clean_phone_number); print("NO")