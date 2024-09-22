import * as esbuild from "esbuild";
import stylePlugin from "esbuild-style-plugin";
import dotenv from "dotenv";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

dotenv.config({ path: `.env.${process.env.NODE_ENV}` });

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
  define: {
    "process.env.APP_HOST": JSON.stringify(process.env.APP_HOST),
  },
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
