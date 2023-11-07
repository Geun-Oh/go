const esbuild = require("esbuild");

esbuild
  .build({
    entryPoints: ["frontend/*"],
    outdir: "public/assets",
    bundle: true,
    minify: true,
  })
  .then(() => console.log("Build Complete"))
  .catch(() => process.exit(1));
