list=[]

for i in range(1,65):
    if i== 1:
        list.append(i)
        
    elif i ==2:
        list.append(i)
        
    elif i ==3:
        list.append(4)
        
    else:
        list.append(2*list[-1]) 
        
total=0

print(list)

for i in list:
    total+=i

print(total)
