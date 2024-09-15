import * as esbuild from "esbuild";
import { sassPlugin } from "esbuild-sass-plugin";

let ctx = await esbuild.context({
  logLevel: "debug",
  entryPoints: ["web/src/App.tsx", "web/src/main.scss"],
  outdir: "public",
  bundle: true,
  minify: true,
  plugins: [sassPlugin()],
});

console.log("⚡ Build complete! ⚡");

await ctx.watch();
