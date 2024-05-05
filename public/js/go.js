// create go instance
const go = new Go();

// start webassembly instance
WebAssembly.instantiateStreaming(fetch('./public/js/main.wasm'), go.importObject)
  .then(
    result => {
      // start go wasm instance
      go.run(result.instance)
      // see the result in browser console
      console.log("Summition happening in Golang")
    }
  )
