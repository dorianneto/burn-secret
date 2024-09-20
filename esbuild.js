import * as esbuild from "esbuild";
import stylePlugin from "esbuild-style-plugin";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

let ctx = await esbuild.context({
  logLevel: "info",
  entryPoints: ["web/src/app.tsx", "web/src/main.scss"],
  outdir: "public",
  bundle: true,
  minify: true,
  sourcemap: true,
  plugins: [
    stylePlugin({
      postcss: {
        plugins: [require("tailwindcss"), require("autoprefixer")],
      },
    }),
  ],
});

console.log("⚡ Build complete! ⚡");

if (process.argv.slice(2)[0] === "--watch") {
  await ctx.watch();
} else {
  await ctx.watch();
  setTimeout(async () => {
    await ctx.dispose();
  }, 1000);
}
