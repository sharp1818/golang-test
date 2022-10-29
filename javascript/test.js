let stack = []
let packet = []
let needclose = []
let needopen =[]
const open = (value) => {
    switch(value) {
        case "[":
            return "]"
        case "(":
            return ")"
        case "{":
            return "}"
        default:
            return null
    }
}
const close = (value) => {
    switch(value) {
        case "]":
            return "["
        case ")":
            return "("
        case "}":
            return "{"
        default:
            return null
    }
}
const result = (s) => {
    const array = s.split('')
    for (let i = 0; i < array.length; i++) {
        const element = open(array[i]);
        if(element != null){
            stack.push(element)
        }
        if(stack[stack.length-1]===array[i]){
            stack.pop()
        } 
        else if (stack.find(e => e===array[i])){
            const index = stack.indexOf(array[i])
            packet.push(array[i])
            stack.splice(index, 1)
        }
        else if (stack[stack.length-1]!==array[i] && element=== null){
            needopen.push(array[i])
        } 
    }
    for ( let i = 0; i < stack.length; i++) {
        needclose[i] = close(stack[i])
    }
    if(array.length== 0){
        console.log("Empty String")
    }else if(needclose.length == 0 && packet.length == 0 && needopen.length == 0) {
        console.log("Valid String")
    }else if(needclose.length>0 && packet.length>0) {
        console.log("Invalid String", packet[0], needclose[0])
    }else if(needclose.length>0){
        console.log("Invalid String, grouping sign need close", needclose[0])
    }else if(needopen.length>0) {
        console.log("Invalid String, grouping sign need open", needopen[0])
    }
}

let test = "[()]{}{[()()]()}"
result(test)
