import terser from "@rollup/plugin-terser";
import nodeResolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import postcss from 'rollup-plugin-postcss';
import del from 'rollup-plugin-delete'
import autoprefixer from "autoprefixer";


export default {
  input: 'web/src/main.js',
  output: {
    hashCharacters: 'base36',
    entryFileNames: 'bundle-[hash:8].js',
    dir: 'web/dist',
    format: 'cjs',
  },
  plugins: [
    del({
      targets: ['web/dist/*.js', 'web/dist/*.css'],
    }),
    commonjs(),
    nodeResolve(),
    terser(),
    postcss({
      extract: true,
      plugins: [autoprefixer()],
      minimize: true,
    }),
  ]
};