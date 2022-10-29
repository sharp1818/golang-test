stack = []
package = []
needopen = []
needclose= []
open = {"{": "}", "[": "]", "(": ")"}
close = {"}": "{", "]": "[", ")": "("}

def result(s):
    array = list(s)
    for s in array:
        try:
            stack.append(open[s])
        except KeyError:
            if len(stack) == 0:
                needopen.append(s)
            elif s == stack[-1]:
                stack.pop()
            elif s in stack:
                stack.remove(s)
                package.append(close[s])
            else:
                needopen.append(s)
    for e in stack:
        needclose.append(close[e])
    if(len(array) == 0):
        print("Empty String")
    elif(len(needclose)==0 and len(package)==0 and len(needopen)==0):
        print("Valid String")
    elif(len(needclose)>0 and len(package)>0):
        print("Invalid String", package[0], needclose[0])
    elif(len(needclose)>0):
        print("Invalid String, grouping sign need close", needclose[0])
    elif(len(needopen)>0):
        print("Invalid String, grouping sign need open", needopen[0])
        
test = "[()]{}{[()()]()}"
result(test)

