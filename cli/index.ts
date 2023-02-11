import { Command } from "https://deno.land/x/cliffy@v0.25.7/command/mod.ts";
import { Add } from "https://raw.githubusercontent.com/ispec-inc/starry/master/starry/cmds/add.ts"

const VERSION = "0.1.0"

let args = Deno.args
if (args.length === 0) {
  args = ["help"]
}
await new Command()
  .name("starry")
  .version(VERSION)
  .command("add", Add)
  .description("Starry cli 🚀")
  .default("help")
  .parse(args);
