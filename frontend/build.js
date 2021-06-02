const fs = require('fs');
const esbuild = require('esbuild');

console.log('Compiling app...');
try {
  const { errors, warnings } = esbuild.buildSync({
    entryPoints: ['src/app.jsx'],
    bundle: true,
    sourcemap: true,
    allowOverwrite: true,
    outfile: 'public/app.js',
  });
  console.log(`App successfully compiled.`);
} catch(err) {
  console.log(`Errors found.`);
  process.exit(1);
}

fs.copyFile('src/index.html', 'public/index.html', err => {
  if (err) throw err;
  console.log('Assets sucessfully copied.');
});