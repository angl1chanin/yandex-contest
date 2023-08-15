troom, tcond = map(int, input().split())
mode = input()

match mode:
    case "freeze":
        if troom <= tcond:
            print(troom)
        else:
            print(tcond)
    case "heat":
        if troom < tcond:
            print(tcond)
        else:
            print(troom)
    case "auto":
        print(tcond)
    case "fan":
        print(troom)