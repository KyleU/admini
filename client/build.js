// Code generated by projectforge.dev using code from the [core] module, which is under license [CC0]
const esbuild = require('esbuild');

esbuild.build({
  entryPoints: ['src/client.ts'],
  bundle: true,
  minify: true,
  sourcemap: true,
  outfile: '../assets/client.js',
  watch: process.argv[2] === "watch" ? {
    onRebuild(error, result) {
      if (error) console.error('watch build failed:', error)
      else console.log('watch build succeeded:', result)
    }
  } : null,
  logLevel: "info"
}).catch((e) => console.error(e.message))
