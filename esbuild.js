import * as esbuild from "esbuild";
import stylePlugin from "esbuild-style-plugin";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

let ctx = await esbuild.context({
  logLevel: "debug",
  entryPoints: ["web/src/app.tsx", "web/src/main.scss"],
  outdir: "public",
  bundle: true,
  minify: true,
  plugins: [
    stylePlugin({
      postcss: {
        plugins: [require("tailwindcss"), require("autoprefixer")],
      },
      renderOptions: {
        sassOptions: {
          sourceMap: true,
        },
      },
    }),
  ],
});

console.log("⚡ Build complete! ⚡");

await ctx.watch();
