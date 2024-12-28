from _typeshed import StrOrBytesPath
from bs4 import BeautifulSoup

with open('text.html') as file:
    html_text = file.read()

soup = BeautifulSoup(html_text, 'lxml')


# .find .find_all
attr_div = soup.find('div', class_="class_name")
print(attr_div)

attr_div_all = soup.find_all('div', class_="class_name")
for div_attr in attr_div_all:
    print(div_attr.text)

# attr_div = soup.find('div', {"class_":"class_name"})
# attr_div.text


# .find_parent() .find_parents()
attr_parent = soup.find(clas_='post_text').find_parent() # Достаёт первого родителя найденого нам атрибоута

# .next_element .previous_element
next_element = soup.find(clas_='post_text').next_element # Находит следуещий элемент, но работает дотошно если будет перенос строки то найдёт его
next_func_element = soup.find(clas_='post_text').find_next() # В таких случаях используется такая ф-я

# .find_next_sibling .find_previous_sibling
next_smb = soup.find(clas_='post_text').find_next_sibling() # Находит следующий тег, который находится с найденым тегом на оодном уровне
