#!/usr/bin/env node
import { program } from "commander";
import pkg from "../package.json";
import { createTodo } from "./actions/create";

program.name("todos").description(pkg.description).version(pkg.version);

program
  .command("create")
  .option("-b, --body <body>", "the todo object to be created")
  .action(createTodo);

program.parse();
