def main():
    print(mortalRabbitFib(97, 16))

def mortalRabbitFib(n_months: int, lifespan: int) -> int:
    rabbits = [1, 1]
    month = 2

    while month < n_months:
        if month < lifespan:
            rabbits.append(rabbits[-1] + rabbits[-2])
        elif month == lifespan or month == lifespan + 1:
            rabbits.append(rabbits[-1] + rabbits[-2] - 1)
        else: 
            rabbits.append(rabbits[-1] + rabbits[-2] - rabbits[-(lifespan+1)])
        month += 1
    
    return rabbits[-1]

if __name__ == '__main__':
    main()