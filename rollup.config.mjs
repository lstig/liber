import terser from "@rollup/plugin-terser";
import nodeResolve from "@rollup/plugin-node-resolve";

export default {
  input: 'web/src/main.js',
  output: {
    file: 'web/dist/bundle.min.js',
    format: 'cjs'
  },
  plugins: [nodeResolve(), terser()]
};