import { Command } from "https://deno.land/x/cliffy@v0.25.7/command/mod.ts";
import { colors } from "https://deno.land/x/cliffy@v0.25.7/ansi/mod.ts";
import {
  Checkbox,
  Confirm,
} from "https://deno.land/x/cliffy@v0.25.7/prompt/mod.ts";
import axiod from "https://deno.land/x/axiod/mod.ts";
import { download } from "https://deno.land/x/download/mod.ts";

import $ from "https://deno.land/x/dax/mod.ts";

export const Add = new Command()
  .description("Initialize Starry Repo")
  .action(async () => {
    await main()
  });

async function main() {
  const apiUrl = "https://api.github.com/repos/ispec-inc/starry/releases/latest"
  const response = await axiod.get(apiUrl)
  const data = response.data.assets.filter((f: any) => {
    return f.content_type === "application/zip"
  })

  const names = data.map((asset: any) => {
    return asset.name.replace(".zip", "")
  })

  let dirs = {}
  data.forEach((asset: any) => {
    const name = asset.name.replace(".zip", "")
    dirs[name] = asset.browser_download_url
  })

  const result = await Checkbox.prompt({
    name: "dirs",
    message: "Select Directory",
    type: Checkbox,
    options: names,
  });

  const confirmed = await Confirm.prompt({
    message: "Are you sure to initialize these packages?",
    hint: `${result.join(", ")}`,
  })

  if (!confirmed) {
    console.log("exit")
    Deno.exit()
  } else {
    console.log()
    console.log(colors.bold.blue("Affirmative"))
  }

  for (const dir of result) {
    const url = dirs[dir]
    const file = `${dir}.zip`
    await download(url, { file, dir: "." })
    const ok = await $`unzip -uq ${file} && rm ${file}`
    if (ok) {
      console.log(`${colors.bold.blue("Created")}: ${dir}`)
    }
  }
}
