list=[]

for i in range(1,65):
    if i== 1:
        list.append(i)      
    else:
        list.append(2*list[-1]) 
        
total=0

print(list)

for i in list:
    total+=i

print(total)
