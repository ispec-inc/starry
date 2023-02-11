import { Command } from "https://deno.land/x/cliffy@v0.25.7/command/mod.ts";
import {
  Checkbox,
  Confirm,
  Input,
  prompt
} from "https://deno.land/x/cliffy@v0.25.7/prompt/mod.ts";

export const Initialize = new Command()
  .description("Initialize Starry Repo")
  .action(async (options: any, source: string, destination?: string) => {
    await main()
  });

async function main() {
  const result = await prompt([
    {
      name: "type",
      message: "Monorepo?",
      type: Confirm,
    },
    {
      name: "Choose repos",
      message: "Select some animals",
      type: Checkbox,
      options: ["dog", "cat", "snake"],
    }
  ]);

  console.log({ result })
}
