const esbuild = require("esbuild");
const { sassPlugin } = require("esbuild-sass-plugin");

esbuild
  .build({
    entryPoints: ["frontend/App.tsx", "frontend/style.scss"],
    outdir: "public/assets",
    bundle: true,
    minify: true,
    plugins: [sassPlugin()],
  })
  .then(() => console.log("Build Complete"))
  .catch(() => process.exit(1));
