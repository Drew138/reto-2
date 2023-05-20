# Codigos en Python

```python
#!/usr/bin/env python
# coding: utf-8

# # Ejercicio 1

# In[2]:


from mrjob.job import MRJob
from statistics import mean

class AverageSalaryBySector(MRJob):

    def mapper(self, _, line):
        _, sector, salary, _ = line.split()
        if idemp != "idemp":
            yield sector, salary

    def reducer(self, key, values):
        yield key, mean(values)

if __name__ == '__main__':
    AverageSalaryBySector.run()


# In[ ]:


from mrjob.job import MRJob
from statistics import mean

class AverageSalaryByEmployee(MRJob):

    def mapper(self, _, line):
        idemp, _, salary, _ = line.split()
        if idemp != "idemp":
            yield idemp, salary_

    def reducer(self, key, values):
        yield key, mean(values)

if __name__ == '__main__':
    AverageSalaryByEmployee.run()


# In[ ]:


from mrjob.job import MRJob


class EconomicSectorByEmployee(MRJob):

    def mapper(self, _, line):
        idemp, sector, * = line.split()
        if idemp != "idemp":
            yield idemp, sector

    def reducer(self, key, values):
        yield key, len(set(values))

if __name__ == '__main__':
    EconomicSectorByEmployee.run()


# # Ejercicio 2

# In[ ]:


from mrjob.job import MRJob


class HighestLowestByStock(MRJob):

    def mapper(self, _, line):
        company, price, date = line.split()
        if idemp != "company":
            yield company, (price, date)

    def reducer(self, key, values):
        yield key, min(values, key=lambda x: x[0])[1]

if __name__ == '__main__':
    HighestLowestByStock.run()


# In[ ]:


from mrjob.job import MRJob


class StableStocks(MRJob):

    def mapper(self, _, line):
        company, price, date = line.split()
        if idemp != "company":
            yield company, (price, date)

    def reducer(self, key, values):
        values.sort(key=lambda x:[1])
        is_stable = True
        prev = 0
        for price, _ in values:
            if price < prev:
                is_stable = False
                break
            prev = value
        if is_stable:
            yield key, min(values, key=lambda x: x[0])[1]

if __name__ == '__main__':
    StableStocks.run()


# In[ ]:


# TODO
from mrjob.job import MRJob


class BlackDay(MRJob):

    def mapper(self, _, line):
        company, price, date = line.split()
        if idemp != "company":
            yield company, (price, date)

    def reducer(self, key, values):
        values.sort(key=lambda x:[1])
        is_stable = True
        prev = 0
        for price, _ in values:
            if price < prev:
                is_stable = False
                break
            prev = value
        if is_stable:
            yield key, min(values, key=lambda x: x[0])[1]

if __name__ == '__main__':
    BlackDay.run()


# # Ejercicio 3

# In[ ]:


from mrjob.job import MRJob
from statistics import mean

class MoviesByUser(MRJob):

    def mapper(self, _, line):
        user,movie,rating,genre,date = line.split()
        if idemp != "User":
            yield user, rating

    def reducer(self, key, values):
        yield key, (sum(values), mean(values))

if __name__ == '__main__':
    AverageSalaryBySector.run()


# In[ ]:


from mrjob.job import MRJob
from statistics import mean

class MostSeenDay(MRJob):

    def mapper(self, _, line):
        user,movie,rating,genre,date = line.split()
        if idemp != "User":
            yield None, date

    def reducer(self, key, values):
        date_map = {}
        for val in values:
            date_map[val] += 1
        
        yield key, max(date_map.values(), key=lambda x: x[1])[0]

if __name__ == '__main__':
    MostSeenDay.run()


# In[ ]:


from mrjob.job import MRJob
from statistics import mean

class LeastSeenDay(MRJob):

    def mapper(self, _, line):
        user,movie,rating,genre,date = line.split()
        if idemp != "User":
            yield movie, (user, rating)

    def reducer(self, key, values):
        yield key, (len(values), mean(values))

if __name__ == '__main__':
    LeastSeenDay.run()


# In[ ]:


from mrjob.job import MRJob
from statistics import mean

class MoviesSeenByUsers(MRJob):

    def mapper(self, _, line):
        user,movie,rating,genre,date = line.split()
        if idemp != "User":
            yield movie, (user, rating)

    def reducer(self, key, values):
        yield key, (len(values), mean(values))

if __name__ == '__main__':
    MoviesSeenByUsers.run()
```
