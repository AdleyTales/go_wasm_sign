const crypto = require('crypto')

const SALT = 'ftmsabcd@1234!'

function md5(str){
  const md5 = crypto.createHash('md5')
  const result = md5.update(SALT).update(str).digest('hex')
  return result
}

const o = {
  name: 'adley',
  age: 18,
  data: [12,32,'dasd']
}

console.log(md5(JSON.stringify(o)));
console.log(md5(JSON.stringify(o)) === 'bc265660ddecee012e7261ac19745d15');
