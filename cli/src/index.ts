#!/usr/bin/env node
import { program } from "commander";
import pkg from "../package.json";
import { createTodo } from "./actions/create";
import { getTodo } from "./actions/get";

program.name("todos").description(pkg.description).version(pkg.version);

program
  .command("create")
  .option("-b, --body <body>", "the todo object to be created")
  .action(createTodo);

program
  .command("get")
  .option("-a, --all", "whether to get all todo items or not", true)
  .option("-i, --id", "the specific id of the todo item we want to get")
  .action(getTodo);

program.parse();
