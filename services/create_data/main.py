from faker import Faker

fake = Faker()

def generate_data(n:int):
    data = {
        "ID": n,
        "Nombre": f"{fake.name()} {fake.last_name()}",
        "Email": fake.email()
    }

    return data

data = [generate_data(n) for n in range(1000)]

print(data)


