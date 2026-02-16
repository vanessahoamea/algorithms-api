import * as esbuild from "esbuild";

await esbuild.build({
  entryPoints: ["script.ts"],
  bundle: true,
  outfile: "dist/script.js",
  external: ["k6"],
  format: "esm"
});