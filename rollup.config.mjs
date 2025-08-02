import terser from "@rollup/plugin-terser";
import nodeResolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";
import postcss from 'rollup-plugin-postcss';
import autoprefixer from "autoprefixer";


export default {
  input: 'web/src/main.js',
  output: {
    file: 'web/dist/bundle.min.js',
    format: 'cjs'
  },
  plugins: [
    commonjs(),
    nodeResolve(),
    terser(),
    postcss({
      extract: true,
      to: 'bundle.min.css',
      plugins: [autoprefixer()],
      minimize: true,
    }),
  ]
};