let x = 0

function runLoop() {
  let loopCount = 0;
  return new Promise(resolve => {
    let loop = setInterval(function() {
      if (loopCount >= 1000) {
        clearInterval(loop)
        resolve(true)
        return
      }
      loopCount = loopCount + 1;
      x = x + 1
    }, 1)

  })
}

Promise.all([
  runLoop(),
  runLoop()
]).then(function() {
  console.log(x)
})

