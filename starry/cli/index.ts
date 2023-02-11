import { Command } from "https://deno.land/x/cliffy@v0.25.7/command/mod.ts";
import { Initialize } from "./commands/initialize.ts"

const VERSION = "0.1.0"

await new Command()
  .name("starry")
  .version(VERSION)
  .command("init", Initialize)
  .description("Starry cli 🚀")
  .parse(Deno.args);
