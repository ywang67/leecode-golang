function deepCopy(obj, res) {
    for (key in obj) {

        if (obj[key] == null || obj[key] == undefined || typeof obj[key] != 'object') {
            res[key] = obj[key]
            console.log('222tester: ', res)
            return
        }
    
        res[key] = {}
        deepCopy(obj[key], res[key])
    }

    return res
}

const copy = () => {
    console.log('copy starting...')
    const obj = {
        a: {
            a1: {
                a2: 111
            }
        },
        b: undefined,
    }
    const res = {}
    deepCopy(obj, res)
    console.log("000tester: ", res)
}

copy()
